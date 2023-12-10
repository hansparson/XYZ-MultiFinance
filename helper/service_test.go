package helper_test

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
	dbcontroller "xyz-multifinance/db/db-controller"
	dbschema "xyz-multifinance/db/db-schema"
	service "xyz-multifinance/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectGorm() *gorm.DB {
	// Load environment variables from .env file
	dbURI := "xyz:Password123@tcp(localhost:3306)/xyz_multifinance?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	// Migrate Product Database
	errors := db.AutoMigrate(dbschema.User{}, dbschema.Transaction{}, dbschema.Bill{}, dbschema.MonthlyBilling{}, dbschema.UserLimitBalance{})
	if errors != nil {
		log.Println(errors.Error())
	}

	return db
}

func TestCreateUser(t *testing.T) {
	db := ConnectGorm()
	control := dbcontroller.Controller(db)
	// Create a Gin router instance and define the route for testing
	router := gin.New()
	router.POST("/create-user", func(c *gin.Context) {
		service.CreateUser(c, control) // pass control (HandlersController) instead of db
	})

	currentTime := time.Now()
	randomNIKString := currentTime.Format("2006-01-02 15:04:05")
	// Payload JSON for the test request
	payload := []byte(`{
		"FullName": "Hans Parson",
		"NIK": "` + randomNIKString + `",
		"LegalName": "Hans Parson",
		"TempatLahir": "Polewali",
		"TanggalLahir": "1997-03-02",
		"Gaji": 12000000,
		"FotoKTP": "https://assets.pikiran-rakyat.com/crop/0x0:0x0/x/photo/2021/04/20/750175463.jpg",
		"FotoSelfie": "https://assets.pikiran-rakyat.com/crop/0x0:0x0/x/photo/2021/04/20/750175463.jpg",
		"LimitOneMonth": 1000000,
		"LimitTwoMonth": 3000000,
		"LimitThreeMonth": 5000000,
		"LimitSixth": 100000000
	}`)

	req, err := http.NewRequest("POST", "/create-user", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	secretKey := os.Getenv("SERVER_KEY")
	key := []byte(secretKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(randomNIKString))
	signature := hex.EncodeToString(h.Sum(nil))

	req.Header.Set("X-EXTERNAL-ID", randomNIKString)
	req.Header.Set("X-SIGNATURE", signature)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	expectedMessageAction := "SUCCESS"

	assert.Equal(t, expectedMessageAction, response["message_action"])
	fmt.Println("+++++++++++++++++++++++++++++++++++++ CREATE USER TESTED +++++++++++++++++++++++++++++++++++++")
}

func TestTransaction(t *testing.T) {
	db := ConnectGorm()
	user_id := "43189094"
	currentTime := time.Now()
	randomNIKString := currentTime.Format("2006-01-02 15:04:05")

	// Set up a mock HTTP request
	requestBody := []byte(`{
		"UserID": "` + user_id + `",
		"OTR": 100,
		"AdminFee": 100,
		"HargaAset": 1000,
		"JumlahBunga": 5,
		"Tenor": 6,
		"NamaAset": "Baju Bola Messi"
	}`)

	req, err := http.NewRequest("POST", "/transaction", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	secretKey := os.Getenv("SERVER_KEY")
	key := []byte(secretKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(randomNIKString))
	signature := hex.EncodeToString(h.Sum(nil))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-EXTERNAL-ID", randomNIKString)
	req.Header.Set("X-SIGNATURE", signature)

	// Create a mock HTTP context
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	// Mock the controller and call the Transaction function
	mockController := dbcontroller.Controller(db) // Replace with your mock controller

	service.Transaction(ctx, mockController)

	// Validate the response or side effects here using assertions
	// For example, check if the response status code is 200 OK
	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
	// Add more assertions based on your specific logic and expected behavior
	fmt.Println("+++++++++++++++++++++++++++++++++++++ TRANSACTION TESTED +++++++++++++++++++++++++++++++++++++")
}
