package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	dbcontroller "xyz-multifinance/db/db-controller"
	dbschema "xyz-multifinance/db/db-schema"
	"xyz-multifinance/schema"
	"xyz-multifinance/utils"
	logger "xyz-multifinance/utils"
	"xyz-multifinance/validation"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context, controller *dbcontroller.HandlersController) {
	apiCallID := logger.GenerateAPICallID()
	service_name := "CREATE_USER"
	payload, headers, params, err := validation.ExtractPayload(ctx, apiCallID)
	if err != nil {
		return
	}
	logger.PrintLogInfo(payload, headers, params, "REQUEST_START", apiCallID, ctx.ClientIP(), ctx.FullPath())

	// Validate Schema
	if err := validation.ValidatesSchema(payload, headers, params, apiCallID, service_name, ctx, schema.CreateUser{}); err != nil {
		fmt.Println(utils.IDLogger(apiCallID) + "Error: " + err.Error())
		return
	}

	// Mengonversi payload string menjadi JSON
	var payloadJSON map[string]interface{}
	if err := json.Unmarshal([]byte(payload), &payloadJSON); err != nil {
		schema.InvalidPayloadFormat("", "", "", ctx, apiCallID)
		return
	}

	user_id := utils.GenerateUserID()
	// Data yang ingin Anda masukkan
	dataUserBaru := dbschema.User{
		UserID:       user_id,
		UserStatus:   "ACTIVE",
		FullName:     payloadJSON["FullName"].(string),
		NIK:          payloadJSON["NIK"].(string),
		LegalName:    payloadJSON["LegalName"].(string),
		TempatLahir:  payloadJSON["TempatLahir"].(string),
		TanggalLahir: payloadJSON["TanggalLahir"].(string),
		Gaji:         payloadJSON["Gaji"].(float64),
		FotoKTP:      payloadJSON["FotoKTP"].(string),
		FotoSelfie:   payloadJSON["FotoSelfie"].(string),
	}

	dataLimitBalancer := dbschema.UserLimitBalance{
		UserID:          user_id,
		LimitOneMonth:   payloadJSON["LimitOneMonth"].(float64),
		LimitTwoMonth:   payloadJSON["LimitTwoMonth"].(float64),
		LimitThreeMonth: payloadJSON["LimitThreeMonth"].(float64),
		LimitSixth:      payloadJSON["LimitSixth"].(float64),
	}

	// Menambahkan data user baru ke dalam database
	err = dbcontroller.AddNewUser(controller.DB, &dataUserBaru)
	if err != nil {
		schema.DuplicateIDResponse(payload, headers, params, ctx, apiCallID)
		return
	}

	// Menambahkan data Limit BAlance User
	err = dbcontroller.AddLimitUser(controller.DB, &dataLimitBalancer)
	if err != nil {
		schema.DuplicateIDResponse(payload, headers, params, ctx, apiCallID)
	}

	message_data := map[string]interface{}{
		"UserID":   user_id,
		"FullName": payloadJSON["FullName"].(string),
		"Status":   "ACTIVE",
	}

	schema.SuccessWithMessage(message_data, payload, headers, params, ctx, apiCallID)
	ctx.Request.Body = ioutil.NopCloser(strings.NewReader(payload))
	return
}

func Transaction(ctx *gin.Context, controller *dbcontroller.HandlersController) {
	apiCallID := logger.GenerateAPICallID()
	service_name := "TRANSACTION"
	payload, headers, params, err := validation.ExtractPayload(ctx, apiCallID)
	if err != nil {
		return
	}
	logger.PrintLogInfo(payload, headers, params, "REQUEST_START", apiCallID, ctx.ClientIP(), ctx.FullPath())

	// Validate Schema
	if err := validation.ValidatesSchema(payload, headers, params, apiCallID, service_name, ctx, schema.Transaction{}); err != nil {
		fmt.Println(utils.IDLogger(apiCallID)+"Error:", err)
		return
	}

	// Mengonversi payload string menjadi JSON
	var payloadJSON map[string]interface{}
	if err := json.Unmarshal([]byte(payload), &payloadJSON); err != nil {
		schema.InvalidPayloadFormat("", "", "", ctx, apiCallID)
		return
	}

	contractID := utils.GenerateContractID()
	user_id := payloadJSON["UserID"].(string)
	tenor := payloadJSON["Tenor"].(float64)
	allowedTenors := []float64{1, 2, 3, 6}
	var found bool
	for _, t := range allowedTenors {
		if tenor == t {
			found = true
			break
		}
	}
	// Jika tenor tidak ditemukan dalam rentang yang diizinkan
	if !found {
		schema.TenorNotAllowed(payload, headers, params, ctx, apiCallID)
		return
	}

	// check Limit Tenor
	var monthly_limit float64
	monthlyLimit, err := dbcontroller.GetMonthlyBillingsByUserID(controller.DB, user_id)
	if err != nil {
		schema.GeneralErrorRequest(payload, headers, params, ctx, apiCallID)
		return
	}

	for _, limitBalance := range monthlyLimit {
		switch tenor {
		case 1:
			monthly_limit = limitBalance.LimitOneMonth
		case 2:
			monthly_limit = limitBalance.LimitTwoMonth
		case 3:
			monthly_limit = limitBalance.LimitThreeMonth
		case 6:
			monthly_limit = limitBalance.LimitSixth
		default:
			// Handle invalid tenor value
		}
	}

	bayangan_transaksi := payloadJSON["OTR"].(float64) + payloadJSON["HargaAset"].(float64) + payloadJSON["AdminFee"].(float64)
	bunga_transaksi := bayangan_transaksi * payloadJSON["JumlahBunga"].(float64) / 100
	cicilan_bulanan := bunga_transaksi + (bayangan_transaksi / tenor)
	jumlah_cicilan := cicilan_bulanan * tenor
	tangal_kontrak := time.Now()

	if jumlah_cicilan > monthly_limit {
		schema.InsulficientMonthlyLimit(payload, headers, params, ctx, apiCallID)
		return
	}
	limit_new := monthly_limit - jumlah_cicilan

	// Data yang ingin Anda masukkan
	TransactionRecord := dbschema.Transaction{
		NomorKontrak:   contractID,
		UserID:         user_id,
		TanggalKontrak: tangal_kontrak,
		TangglUpdate:   time.Now(),
		OTR:            payloadJSON["OTR"].(float64),
		AdminFee:       payloadJSON["AdminFee"].(float64),
		HargaAset:      payloadJSON["HargaAset"].(float64),
		JumlahCicilan:  jumlah_cicilan,
		JumlahBunga:    payloadJSON["JumlahBunga"].(float64),
		Tenor:          tenor,
		CicilanBulanan: cicilan_bulanan,
		NamaAset:       payloadJSON["NamaAset"].(string),
	}

	// Menambahkan data Limit BAlance User
	err = dbcontroller.AddTransaction(controller.DB, &TransactionRecord)
	if err != nil {
		schema.GeneralErrorRequest(payload, headers, params, ctx, apiCallID)
		return
	}

	// Update Limit Sementara User
	err = dbcontroller.UpdateLimitBalance(controller.DB, user_id, float64(limit_new), float64(tenor))
	if err != nil {
		schema.GeneralErrorRequest(payload, headers, params, ctx, apiCallID)
		return
	}

	currentTime := time.Now()
	var array_bill []dbschema.MonthlyBilling
	for i := 1; i <= int(tenor); i++ {
		nextMonth := currentTime.AddDate(0, i, 0)

		// Buatkan Biling ID untuk tiap Cicilan Bulanan
		MonthlyBilling := dbschema.MonthlyBilling{
			BillingId:      "BILL-" + utils.GenerateContractID(),
			UserID:         user_id,
			NamaAset:       payloadJSON["NamaAset"].(string),
			Tenor:          float64(i),
			HargaCicilan:   cicilan_bulanan,
			TanggalTagihan: nextMonth,
		}
		array_bill = append(array_bill, MonthlyBilling)
		// Menambahkan data Limit BAlance User
		err = dbcontroller.AddMonthlyBill(controller.DB, &MonthlyBilling)
		if err != nil {
			schema.GeneralErrorRequest(payload, headers, params, ctx, apiCallID)
			return
		}
	}

	// check Limit Tenor
	var LimitOneMonth float64
	var LimitTwoMonth float64
	var LimitThreeMonth float64
	var LimitSixth float64

	NewmonthlyLimit, err := dbcontroller.GetMonthlyBillingsByUserID(controller.DB, user_id)
	if err != nil {
		schema.GeneralErrorRequest(payload, headers, params, ctx, apiCallID)
		return
	}

	for _, limitBalance := range NewmonthlyLimit {
		LimitOneMonth = limitBalance.LimitOneMonth
		LimitTwoMonth = limitBalance.LimitTwoMonth
		LimitThreeMonth = limitBalance.LimitThreeMonth
		LimitSixth = limitBalance.LimitSixth
	}

	message_data := map[string]interface{}{
		"UserID":         user_id,
		"TanggalKontrak": tangal_kontrak,
		"OTR":            payloadJSON["OTR"].(float64),
		"JumlahCicilan":  jumlah_cicilan,
		"JumlahBunga":    payloadJSON["JumlahBunga"].(float64),
		"Tenor":          payloadJSON["Tenor"].(float64),
		"NamaAset":       payloadJSON["NamaAset"].(string),
		"LimitSementar": dbschema.LimitSementara{
			LimitOneMonth:   LimitOneMonth,
			LimitTwoMonth:   LimitTwoMonth,
			LimitThreeMonth: LimitThreeMonth,
			LimitSixth:      LimitSixth,
		},
		"TagihanBulanan": array_bill,
	}

	schema.SuccessWithMessage(message_data, payload, headers, params, ctx, apiCallID)

	// schema.Success(payload, headers, params, ctx, apiCallID)
	ctx.Request.Body = ioutil.NopCloser(strings.NewReader(payload))
	return
}
