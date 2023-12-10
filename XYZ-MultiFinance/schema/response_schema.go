package schema

import (
	"encoding/json"
	"net/http"
	logger "xyz-multifinance/utils"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message_Action string `json:"message_action" example:"DUPLICATE_ID"`
	Message_Data   string `json:"message_data"`
	API_Call_ID    string `json:"api_call_id"`
}

type ResponseWithMesage struct {
	Message_Action string                 `json:"message_action" example:"DUPLICATE_ID"`
	Message_Data   map[string]interface{} `json:"message_data"`
	API_Call_ID    string                 `json:"api_call_id"`
}

func MarshalResponseToJSON(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func LogResponseAndJSONMarshalError(payload string, header string, params string, ctx *gin.Context, apiCallID string, responseData interface{}) string {
	responseJSON, err := MarshalResponseToJSON(responseData)
	if err != nil {
		// Handle error saat melakukan marshaling ke JSON
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal response"})
		return ""
	}

	logger.PrintLogInfo(payload, header, params, "RESPONSE_END", apiCallID, ctx.ClientIP(), ctx.FullPath(), responseJSON)
	return responseJSON
}

// All Response
func Success(payload string, header string, params string, ctx *gin.Context, apiCallID string) {
	SuccesREsponse := &Response{
		Message_Action: "SUCCESS",
		Message_Data:   "Data successfully stored",
		API_Call_ID:    apiCallID,
	}
	LogResponseAndJSONMarshalError(payload, header, params, ctx, apiCallID, SuccesREsponse)
	ctx.JSON(200, SuccesREsponse)
}

func SuccessWithMessage(message_data map[string]interface{}, payload string, header string, params string, ctx *gin.Context, apiCallID string) {
	SuccesREsponse := &ResponseWithMesage{
		Message_Action: "SUCCESS",
		Message_Data:   message_data,
		API_Call_ID:    apiCallID,
	}
	LogResponseAndJSONMarshalError(payload, header, params, ctx, apiCallID, SuccesREsponse)
	ctx.JSON(200, SuccesREsponse)
}

func DuplicateIDResponse(payload string, header string, params string, ctx *gin.Context, apiCallID string) {
	duplicateIDResponse := &Response{
		Message_Action: "DUPLICATE_ID",
		Message_Data:   "Your data is duplicate.",
		API_Call_ID:    apiCallID,
	}
	LogResponseAndJSONMarshalError(payload, header, params, ctx, apiCallID, duplicateIDResponse)
	ctx.JSON(409, duplicateIDResponse)
}

func InvalidPayloadFormat(payload string, header string, params string, ctx *gin.Context, apiCallID string) {
	duplicateIDResponse := &Response{
		Message_Action: "INVALID_PAYLOAD_FORMAT",
		Message_Data:   "your format payload is invalid",
		API_Call_ID:    apiCallID,
	}
	LogResponseAndJSONMarshalError(payload, header, params, ctx, apiCallID, duplicateIDResponse)
	ctx.JSON(404, duplicateIDResponse)
}

func InvalidHeaderFormat(payload string, header string, params string, ctx *gin.Context, apiCallID string) {
	duplicateIDResponse := &Response{
		Message_Action: "INVALID_HEADER_FORMAT",
		Message_Data:   "your format header is invalid",
		API_Call_ID:    apiCallID,
	}
	LogResponseAndJSONMarshalError(payload, header, params, ctx, apiCallID, duplicateIDResponse)
	ctx.JSON(404, duplicateIDResponse)
}

func InvalidHeaderMandatory(message, payload string, header string, params string, ctx *gin.Context, apiCallID string) {
	duplicateIDResponse := &Response{
		Message_Action: "INVALID_HEADER_Mandatory",
		Message_Data:   message,
		API_Call_ID:    apiCallID,
	}
	LogResponseAndJSONMarshalError(payload, header, params, ctx, apiCallID, duplicateIDResponse)
	ctx.JSON(404, duplicateIDResponse)
}

func InvalidPayloadKey(message, payload string, header string, params string, ctx *gin.Context, apiCallID string) {
	duplicateIDResponse := &Response{
		Message_Action: "INVALID_PAYLOAD_KEY",
		Message_Data:   message,
		API_Call_ID:    apiCallID,
	}
	LogResponseAndJSONMarshalError(payload, header, params, ctx, apiCallID, duplicateIDResponse)
	ctx.JSON(404, duplicateIDResponse)
}

func InvalidPayloadMandatory(message, payload string, header string, params string, ctx *gin.Context, apiCallID string) {
	duplicateIDResponse := &Response{
		Message_Action: "INVALID_MANDATORY",
		Message_Data:   message,
		API_Call_ID:    apiCallID,
	}
	LogResponseAndJSONMarshalError(payload, header, params, ctx, apiCallID, duplicateIDResponse)
	ctx.JSON(404, duplicateIDResponse)
}

func ConflicExternalId(payload string, header string, params string, ctx *gin.Context, apiCallID string) {
	duplicateIDResponse := &Response{
		Message_Action: "CONFLICT_EXTERNAL_ID",
		Message_Data:   "External Id already in use",
		API_Call_ID:    apiCallID,
	}
	LogResponseAndJSONMarshalError(payload, header, params, ctx, apiCallID, duplicateIDResponse)
	ctx.JSON(404, duplicateIDResponse)
}

func InvalidSignature(payload string, header string, params string, ctx *gin.Context, apiCallID string) {
	duplicateIDResponse := &Response{
		Message_Action: "INVALID_SIGNATURE",
		Message_Data:   "signature not valid",
		API_Call_ID:    apiCallID,
	}
	LogResponseAndJSONMarshalError(payload, header, params, ctx, apiCallID, duplicateIDResponse)
	ctx.JSON(409, duplicateIDResponse)
}

func GeneralErrorRequest(payload string, header string, params string, ctx *gin.Context, apiCallID string) {
	duplicateIDResponse := &Response{
		Message_Action: "GENERAL_ERROR_REQUEST",
		Message_Data:   "internal error on system, please try again",
		API_Call_ID:    apiCallID,
	}
	LogResponseAndJSONMarshalError(payload, header, params, ctx, apiCallID, duplicateIDResponse)
	ctx.JSON(404, duplicateIDResponse)
}

func TenorNotAllowed(payload string, header string, params string, ctx *gin.Context, apiCallID string) {
	duplicateIDResponse := &Response{
		Message_Action: "TEENOR_NOT_ALLOWED",
		Message_Data:   "tenor not allowed, only allow 1, 2, 3, 6 month tenor",
		API_Call_ID:    apiCallID,
	}
	LogResponseAndJSONMarshalError(payload, header, params, ctx, apiCallID, duplicateIDResponse)
	ctx.JSON(404, duplicateIDResponse)
}

func InsulficientMonthlyLimit(payload string, header string, params string, ctx *gin.Context, apiCallID string) {
	duplicateIDResponse := &Response{
		Message_Action: "INSULFICIENT_MONLY_LIMIT",
		Message_Data:   "monly limit not enoufh for transaction, please try again",
		API_Call_ID:    apiCallID,
	}
	LogResponseAndJSONMarshalError(payload, header, params, ctx, apiCallID, duplicateIDResponse)
	ctx.JSON(404, duplicateIDResponse)
}
