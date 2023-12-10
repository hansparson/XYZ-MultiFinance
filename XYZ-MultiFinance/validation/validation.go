package validation

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
	"xyz-multifinance/db/redis_db"
	"xyz-multifinance/schema"
	"xyz-multifinance/utils"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func generateHMACSHA256(message string) string {
	secretKey := os.Getenv("SERVER_KEY")
	key := []byte(secretKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	signature := hex.EncodeToString(h.Sum(nil))
	return signature
}

func ExtractPayload(ctx *gin.Context, apiCallID string) (string, string, string, error) {
	// Mengekstrak payload dari body request
	var payload map[string]interface{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		// Mengembalikan response ctx dengan pesan kesalahan jika ekstraksi payload gagal
		schema.InvalidPayloadFormat("", "", "", ctx, apiCallID)
		return "", "", "", fmt.Errorf("invalid format Payload: %w", err)
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", "", "", err
	}
	payloadStr := string(payloadJSON)

	// Mengekstrak headers dari request
	headersJSON, err := json.Marshal(ctx.Request.Header)
	if err != nil {
		return "", "", "", err
	}
	headersStr := string(headersJSON)

	// Mengekstrak  (query parameters) dari URL
	paramsJSON, err := json.Marshal(ctx.Request.URL.Query())
	if err != nil {
		return "", "", "", err
	}
	paramsStr := string(paramsJSON)

	return payloadStr, headersStr, paramsStr, nil
}

func ValidatePayload(payload string, apiCallID string, ctx *gin.Context, SchemaPayload interface{}) error {
	// Mengonversi payload string menjadi JSON
	var payloadJSON map[string]interface{}
	if err := json.Unmarshal([]byte(payload), &payloadJSON); err != nil {
		schema.InvalidPayloadFormat("", "", "", ctx, apiCallID)
		return err
	}

	// Mengonversi SchemaPayload menjadi bentuk map[string]interface{}
	schemaPayloadMap := make(map[string]interface{})
	b, _ := json.Marshal(SchemaPayload)
	json.Unmarshal(b, &schemaPayloadMap)

	// Membandingkan setiap key dalam payload dengan struktur SchemaPayload
	for key := range payloadJSON {
		if _, ok := schemaPayloadMap[key]; !ok {
			message := fmt.Sprintf("Invalid payload key: %s", key)
			schema.InvalidPayloadKey(message, "", "", "", ctx, apiCallID)
			return fmt.Errorf("Invalid payload key")
		}
	}

	// Mengecek ketersediaan field mandatory dalam SchemaPayload
	schemaPayloadValue := reflect.ValueOf(SchemaPayload)
	for i := 0; i < schemaPayloadValue.NumField(); i++ {
		field := schemaPayloadValue.Type().Field(i)
		fieldName := field.Tag.Get("json")

		if field.Tag.Get("validate") == "mandatory" {
			if _, ok := payloadJSON[fieldName]; !ok {
				message := fmt.Sprintf("Payload key '%s' is needed!", fieldName)
				schema.InvalidPayloadMandatory(message, "", "", "", ctx, apiCallID)
				return fmt.Errorf("Invalid payload mandatory")
			}
		}
	}
	return nil
}

func ValidateHeader(headers string, apiCallID string, ctx *gin.Context) error {
	// Mengonversi payload string menjadi JSON
	headers = strings.ReplaceAll(headers, "[", "")
	headers = strings.ReplaceAll(headers, "]", "")
	var headersJSON map[string]interface{}
	if err := json.Unmarshal([]byte(headers), &headersJSON); err != nil {
		schema.InvalidPayloadFormat("", "", "", ctx, apiCallID)
		return err
	}

	// Membuat daftar kunci yang diharapkan ada dalam headersJSON
	expectedKeys := []string{"X-External-Id", "X-Signature"}

	// Memeriksa keberadaan kunci tertentu dalam headersJSON
	for _, key := range expectedKeys {
		if _, ok := headersJSON[key]; !ok {
			message := fmt.Sprintf("Invalid header %s is missing", key)
			schema.InvalidHeaderMandatory(message, "", "", "", ctx, apiCallID)
			return fmt.Errorf(message)
		}
	}

	return nil
}

func ValidateExternalKey(headers string, apiCallID string, service_name string, ctx *gin.Context) error {
	expiration := 24 * time.Hour
	redisClient := redis_db.SetupRedisConnection()

	// Mengonversi payload string menjadi JSON
	headers = strings.ReplaceAll(headers, "[", "")
	headers = strings.ReplaceAll(headers, "]", "")
	var headersJSON map[string]interface{}
	if err := json.Unmarshal([]byte(headers), &headersJSON); err != nil {
		schema.InvalidPayloadFormat("", "", "", ctx, apiCallID)
		return err
	}

	external_key := headersJSON["X-External-Id"].(string) + "__" + service_name
	fmt.Println(utils.IDLogger(apiCallID) + external_key)

	// Check External Id on Redis server
	_, err := redisClient.Get(context.Background(), external_key).Result()
	if err == nil {
		schema.ConflicExternalId("", "", "", ctx, apiCallID)
		return fmt.Errorf("External Id Terdaftar")
	} else {
		err := redisClient.Set(context.Background(), external_key, "ok", expiration).Err()
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func ValidateSignature(headers string, apiCallID string, service_name string, ctx *gin.Context) error {
	// Mengonversi payload string menjadi JSON
	headers = strings.ReplaceAll(headers, "[", "")
	headers = strings.ReplaceAll(headers, "]", "")
	var headersJSON map[string]interface{}
	if err := json.Unmarshal([]byte(headers), &headersJSON); err != nil {
		schema.InvalidPayloadFormat("", "", "", ctx, apiCallID)
		return err
	}

	external_key := headersJSON["X-External-Id"].(string)
	request_signature := headersJSON["X-Signature"].(string)

	generate_signature := generateHMACSHA256(external_key)
	fmt.Println(utils.IDLogger(apiCallID)+"Generating True signature: ", generate_signature)

	if request_signature != generate_signature {
		schema.InvalidSignature("", "", "", ctx, apiCallID)
		return fmt.Errorf("External Id Terdaftar")
	}
	fmt.Println(utils.IDLogger(apiCallID) + external_key)

	return nil
}

func ValidatesSchema(payload string, headers string, params string, apiCallID string, service_name string, ctx *gin.Context, SchemaPayload interface{}) error {
	// Validate the payload
	if err := ValidatePayload(payload, apiCallID, ctx, SchemaPayload); err != nil {
		fmt.Println(utils.IDLogger(apiCallID)+"Error:", err)
		return err
	}

	// Validate the Header
	if err := ValidateHeader(headers, apiCallID, ctx); err != nil {
		fmt.Println(utils.IDLogger(apiCallID)+"Error:", err)
		return err
	}

	// Validate Signature
	if err := ValidateSignature(headers, apiCallID, service_name, ctx); err != nil {
		fmt.Println(utils.IDLogger(apiCallID)+"Error:", err)
		return err
	}

	// Validate External Key
	if err := ValidateExternalKey(headers, apiCallID, service_name, ctx); err != nil {
		fmt.Println(utils.IDLogger(apiCallID)+"Error:", err)
		return err
	}

	return nil
}
