package swiftcode_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/SwanHtetAungPhyo/swifcode/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SwanHtetAungPhyo/swifcode/internal/handler"
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func setupRouter() (*gin.Engine, *mocks.MockServiceMethods) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockService := new(mocks.MockServiceMethods)
	log := logrus.New()

	swiftCodeHandlers := handler.NewSwiftCodeHandlers(mockService, log)

	router.POST("/v1/swift-codes", swiftCodeHandlers.Create)
	router.GET("/v1/swift-codes/:swift-code", swiftCodeHandlers.GetBySwiftCode)
	router.GET("/v1/swift-codes/country/:countryISO2code", swiftCodeHandlers.GetByCountryISO2Code)
	router.DELETE("/v1/swift-codes/:swift-code", swiftCodeHandlers.DeleteBySwiftCode)
	return router, mockService
}

func TestCreateSwiftCode_Success(t *testing.T) {
	router, mockService := setupRouter()

	swiftCodeDto := &model.SwiftCodeDto{
		Address:       "123 Example St",
		BankName:      "Test Bank",
		CountryISO2:   "US",
		CountryName:   "United States",
		IsHeadquarter: true,
		SwiftCode:     "TESTUS33",
	}

	payload, _ := json.Marshal(swiftCodeDto)
	req, _ := http.NewRequest("POST", "/v1/swift-codes", bytes.NewReader(payload))
	w := httptest.NewRecorder()

	mockService.On("Create", swiftCodeDto).Return(nil).Once()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, `{"message":"Swift Code created successfully"}`, w.Body.String())
	mockService.AssertExpectations(t)
}

func TestCreateSwiftCode_InvalidJSON(t *testing.T) {
	router, mockService := setupRouter()

	invalidJSON := `{"swiftCode": "TEST123"`
	req, _ := http.NewRequest("POST", "/v1/swift-codes", bytes.NewReader([]byte(invalidJSON)))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"message":"Invalid request format"}`, w.Body.String())
	mockService.AssertNotCalled(t, "Create")
}

func TestCreateSwiftCode_FailedToCreate(t *testing.T) {
	router, mockService := setupRouter()

	swiftCodeDto := &model.SwiftCodeDto{
		Address:       "456 Sample Ave",
		BankName:      "Failed Bank",
		CountryISO2:   "",
		CountryName:   "United Kingdom",
		IsHeadquarter: false,
		SwiftCode:     "FAILUK22",
	}

	payload, _ := json.Marshal(swiftCodeDto)
	req, _ := http.NewRequest("POST", "/v1/swift-codes", bytes.NewReader(payload))
	w := httptest.NewRecorder()

	mockService.On("Create", swiftCodeDto).Return(errors.New("database error")).Once()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	expectedResponse := `{"message": "You need to provide Swift Code and Data. All fields must be filled"}`
	assert.JSONEq(t, expectedResponse, string(w.Body.Bytes()))

}

func TestGetBySwiftCode_failure(t *testing.T) {
	router, mockService := setupRouter()

	mockService.On("GetBySwiftCode", "VALID123").Return(nil, errors.New("Invalid Swift Code")).Once()

	req, _ := http.NewRequest("GET", "/v1/swift-codes/VALID123", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetByCountryISO2Code_MissingParameter(t *testing.T) {
	mockService := new(mocks.MockServiceMethods)
	log := logrus.New()
	handlers := handler.NewSwiftCodeHandlers(mockService, log)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{
		{Key: "countryISO2code", Value: ""},
	}

	handlers.GetByCountryISO2Code(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"message":"Invalid Country ISO2 Code"}`, w.Body.String())
	mockService.AssertNotCalled(t, "GetByCountryISO")
}

func TestGetByCountryISO2Code_ServiceError(t *testing.T) {
	router, mockService := setupRouter()

	testISO := "US"
	mockService.On("GetByCountryISO", testISO).Return(nil, errors.New("database failure")).Once()

	req, _ := http.NewRequest("GET", "/v1/swift-codes/country/"+testISO, nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.JSONEq(t, `{"message":"Failed to fetch Swift Codes for country"}`, w.Body.String())
	mockService.AssertExpectations(t)
}

func TestDeleteBySwiftCode_MissingParameter(t *testing.T) {
	mockService := new(mocks.MockServiceMethods)
	log := logrus.New()
	handler := handler.NewSwiftCodeHandlers(mockService, log)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{}

	handler.DeleteBySwiftCode(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"message":"Swift Code is Malformed"}`, w.Body.String())
	mockService.AssertNotCalled(t, "Delete")
}
