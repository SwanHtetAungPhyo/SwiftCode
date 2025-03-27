package swiftcode_test

import (
	"bytes"
	"context"
	"github.com/SwanHtetAungPhyo/swifcode/internal/handler"
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	repo2 "github.com/SwanHtetAungPhyo/swifcode/internal/repo"
	"github.com/SwanHtetAungPhyo/swifcode/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

var dummyData = model.CountryISO2Response{
	CountryISO2: "AW",
	CountryName: "ARUBA",
	SwiftCode: []model.SwiftCodeDto{
		{
			Address:       "VONDELLAAN 31  - ORANJESTAD ORANJESTAD-WEST AND ORANJESTAD-EAST ",
			BankName:      "BANCO DI CARIBE (ARUBA) N.V",
			CountryISO2:   "AW",
			IsHeadquarter: true,
			SwiftCode:     "BDCCAWAWXXX",
		},
		{
			Address:       "ITALIESTRAAT 36  - ORANJESTAD ORANJESTAD-WEST AND ORANJESTAD-EAST ",
			BankName:      "RBC ROYAL BANK (ARUBA) N.V. (FORMERLY RBTT BANK ARUBA N.V.)",
			CountryISO2:   "AW",
			IsHeadquarter: true,
			SwiftCode:     "RBTTAWAWXXX",
		},
		{
			Address:       "CAMACURI 12  - ORANJESTAD ORANJESTAD-WEST AND ORANJESTAD-EAST ",
			BankName:      "ARUBA BANK, LTD",
			CountryISO2:   "AW",
			IsHeadquarter: true,
			SwiftCode:     "ARUBAWAXXXX",
		},
		{
			Address:       "KAYA GILBERTO FRANCOIS CROES 53 ORANJESTAD, ORANJESTAD-WEST AND ORANJESTAD-EAST",
			BankName:      "CARIBBEAN MERCANTILE BANK N.V.",
			CountryISO2:   "AW",
			IsHeadquarter: true,
			SwiftCode:     "CMBAAWAXXXX",
		},
		{
			Address:       "WILHELMINASTRAAT 36  - ORANJESTAD ORANJESTAD-WEST AND ORANJESTAD-EAST ",
			BankName:      "AIB BANK NV",
			CountryISO2:   "AW",
			IsHeadquarter: true,
			SwiftCode:     "ANIBAWA1XXX",
		},
		{
			Address:       "CAYA G.F. CROES 38  - ORANJESTAD ORANJESTAD-WEST AND ORANJESTAD-EAST ",
			BankName:      "INTERBANK ARUBA NV",
			CountryISO2:   "AW",
			IsHeadquarter: true,
			SwiftCode:     "IARUAWA1XXX",
		},
		{
			Address:       "J.E. IRAUSQUIN 8  - ORANJESTAD ORANJESTAD-WEST AND ORANJESTAD-EAST ",
			BankName:      "CENTRALE BANK VAN ARUBA",
			CountryISO2:   "AW",
			IsHeadquarter: true,
			SwiftCode:     "CBARAWAWXXX",
		},
		{
			Address:       "TANKI LENDEERT 143  - ORANJESTAD ORANJESTAD-WEST AND ORANJESTAD-EAST ",
			BankName:      "IMTRADEX INTERNATIONAL N.V.",
			CountryISO2:   "AW",
			IsHeadquarter: true,
			SwiftCode:     "IMIEAWA1XXX",
		},
	},
}

func setupService(t *testing.T) (*services.SwiftCodeServices, func()) {
	ctx := context.Background()
	_, _, instance, clean, err := SetupPostgresContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}

	//defer clean()
	log := logrus.New()
	processor := services.NewBankProcessor(instance, log)
	processor.ProcessData("data/swift_codes.csv")

	repo := repo2.NewRepository(instance, log)
	service := services.NewService(repo, log)

	return service, clean
}

func setupRouters(service *services.SwiftCodeServices) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	log := logrus.New()

	swiftHandler := handler.NewSwiftCodeHandlers(service, log)
	router.GET("/v1/swift-codes/:swift-code", swiftHandler.GetBySwiftCode)
	router.GET("/v1/swift-codes/country/:countryISO2code", swiftHandler.GetByCountryISO2Code)
	router.POST("/v1/swift-codes", swiftHandler.Create)
	router.DELETE("/v1/swift-codes/:swift-code", swiftHandler.DeleteBySwiftCode)

	return router
}
func sendRequest(t *testing.T, router *gin.Engine, method, url string, body interface{}) *httptest.ResponseRecorder {
	var req *http.Request
	if body == nil {
		req = httptest.NewRequest(method, url, nil)
	} else {
		b, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		req = httptest.NewRequest(method, url, bytes.NewBuffer(b))
		req.Header.Set("Content-Type", "application/json")
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
func TestCreateSwiftCode(t *testing.T) {
	service, clean := setupService(t)
	defer clean()
	router := setupRouters(service)
	t.Run("WithoutBody", func(t *testing.T) {
		w := sendRequest(t, router, http.MethodPost, "/v1/swift-codes", nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("WithBody", func(t *testing.T) {
		w := sendRequest(t, router, http.MethodPost, "/v1/swift-codes", &model.SwiftCodeDto{
			Address:       "Krakow",
			BankName:      "Santander Banks",
			SwiftCode:     "KRAKOWDDXXX",
			CountryISO2:   "PL",
			CountryName:   "POLAND",
			IsHeadquarter: true,
		})
		assert.Equal(t, http.StatusCreated, w.Code)
	})
	t.Run("EmptyBody", func(t *testing.T) {
		w := sendRequest(t, router, http.MethodPost, "/v1/swift-codes", &model.SwiftCodeDto{})
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "You need to provide Swift Code and Data")
	})
}
func TestGetBySwiftCode(t *testing.T) {
	service, _ := setupService(t)
	router := setupRouters(service)

	mustResponse := model.HeadquarterResponse{
		Address:       "AGUSTINAS 975, FLOOR 2 CASA MATRIZ BANCO DE CHILE - SANTIAGO PROVINCIA DE SANTIAGO, 8320000",
		BankName:      "BANCHILE CORREDORES DE BOLSA S.A.",
		CountryISO2:   "CL",
		CountryName:   "CHILE",
		SwiftCode:     "BCOSCLR1XXX",
		IsHeadquarter: true,
	}
	testSwiftCode := "BCOSCLR1XXX"
	nonExistedSwiftCode := "MMMMDJJKXXX"
	malformedSwiftCode := "ABCD"
	t.Run("Swift Code fetched successfully", func(t *testing.T) {
		w := sendRequest(t, router, "GET", "/v1/swift-codes/"+testSwiftCode, nil)
		assert.Equal(t, http.StatusOK, w.Code)
		var expectedResponse model.HeadquarterResponse
		err := json.Unmarshal(w.Body.Bytes(), &expectedResponse)
		assert.NoError(t, err, "Response should be valid JSON")
		assert.Equal(t, mustResponse, expectedResponse)
	})
	t.Run("Swift Code does not exist", func(t *testing.T) {
		w := sendRequest(t, router, "GET", "/v1/swift-codes/"+nonExistedSwiftCode, nil)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
	t.Run("Empty  Params", func(t *testing.T) {
		w := sendRequest(t, router, "GET", "/v1/swift-codes/"+malformedSwiftCode, nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Swift Code is Malformed")
	})
	t.Run("Concurrent Access", func(t *testing.T) {
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				w := sendRequest(t, router, "GET", "/v1/swift-codes/"+testSwiftCode, nil)
				assert.Equal(t, http.StatusOK, w.Code)
			}()
		}
		wg.Wait()
	})
}
func TestGetByCountryISO2(t *testing.T) {
	service, _ := setupService(t)
	router := setupRouters(service)

	testCountryISO2 := "AW"
	badRequestISO2 := "A"
	t.Run("Failed ISO2 code input", func(t *testing.T) {
		w := sendRequest(t, router, "GET", "/v1/swift-codes/country/"+badRequestISO2, nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)

	})
	t.Run("Success ISO2 code input", func(t *testing.T) {
		w := sendRequest(t, router, "GET", "/v1/swift-codes/country/"+testCountryISO2, nil)
		assert.Equal(t, 200, w.Code, "Expected status code to be 200")
		var expectedResponse model.CountryISO2Response
		if err := json.Unmarshal(w.Body.Bytes(), &expectedResponse); err != nil {
			assert.Nil(t, err, "response should be valid json")
			t.Fatal(err)
		}
		assert.Equal(t, dummyData.CountryISO2, expectedResponse.CountryISO2)
		assert.Equal(t, dummyData.CountryName, expectedResponse.CountryName)
		assert.NotNil(t, expectedResponse.SwiftCode)

		for _, swiftCode := range expectedResponse.SwiftCode {
			assert.Contains(t, dummyData.SwiftCode, swiftCode)
		}
	})
}

func TestDeleteBySwiftCode(t *testing.T) {
	service, _ := setupService(t)
	router := setupRouters(service)
	nonExistedSwiftCode := "ABCDEFGHIJX"
	malformedSwiftCode := "ABCD"
	existingSwiftCode := "IARUAWA1XXX"
	t.Run("Swift Code deleted successfully", func(t *testing.T) {
		w := sendRequest(t, router, http.MethodDelete, "/v1/swift-codes/"+existingSwiftCode, nil)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Swift Code deleted successfully")
	})
	t.Run("Swift Code does not exist", func(t *testing.T) {
		w := sendRequest(t, router, http.MethodDelete, "/v1/swift-codes/"+nonExistedSwiftCode, nil)
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "Swift Code not found")
	})
	t.Run("Empty  Params", func(t *testing.T) {
		w := sendRequest(t, router, http.MethodDelete, "/v1/swift-codes/"+malformedSwiftCode, nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Swift Code is Malformed")
	})
}
