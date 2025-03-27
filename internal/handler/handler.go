package handler

import (
	"errors"
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	"github.com/SwanHtetAungPhyo/swifcode/internal/services"
	"github.com/SwanHtetAungPhyo/swifcode/pkg/custom_errors"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"reflect"
	"regexp"
	"time"
)

var (
	cachestore = persistence.NewInMemoryStore(time.Minute)
)

type Methods interface {
	Create(c *gin.Context)
	GetBySwiftCode(c *gin.Context)
	GetByCountryISO2Code(c *gin.Context)
	DeleteBySwiftCode(c *gin.Context)
}

type SwiftCodeHandlers struct {
	srvInst    services.ServiceMethods
	HandlerLog *logrus.Logger
}

func NewSwiftCodeHandlers(serviceMethod services.ServiceMethods, handlerLog *logrus.Logger) *SwiftCodeHandlers {
	return &SwiftCodeHandlers{
		srvInst:    serviceMethod,
		HandlerLog: handlerLog,
	}
}

// Create CreateNewSwiftCodeHandler godoc
// @Summary      Create a new Swift Code
// @Description  Create a new Swift Code by providing a JSON payload
// @Tags         SwiftCode
// @Accept       json
// @Produce      json
// @Param        swiftCode body model.SwiftCodeDto true "Swift Code data"
// @Success      201 {object} model.ApiResponse
// @Failure      400 {object} model.ApiResponse
// @Failure      500 {object} model.ApiResponse
// @Router       /v1/swift-codes [post]
func (s *SwiftCodeHandlers) Create(c *gin.Context) {
	s.HandlerLog.Info("Received request to create Swift Code")
	var swiftCode *model.SwiftCodeDto
	if err := c.ShouldBindJSON(&swiftCode); err != nil {
		s.HandlerLog.WithError(err).Error("Failed to bind Swift Code request")
		c.JSON(http.StatusBadRequest, model.ApiResponse{
			Message: "Invalid request format",
		})
		return
	}
	if s.isAnyFieldNil(swiftCode) {
		c.JSON(http.StatusBadRequest, model.ApiResponse{
			Message: "You need to provide Swift Code and Data And all fields must be filled",
		})
	}
	if err := s.srvInst.Create(swiftCode); err != nil {
		s.HandlerLog.WithError(err).Error("Failed to create Swift Code")
		c.JSON(http.StatusInternalServerError, model.ApiResponse{Message: "Failed to create Swift Code"})
		return
	}

	s.HandlerLog.Info("Swift Code created successfully")
	c.JSON(http.StatusCreated, model.ApiResponse{Message: "Swift Code created successfully"})
}

// GetBySwiftCode godoc
// @Summary      Fetch a Swift Code by its code
// @Description  Retrieve the Swift Code by providing the Swift Code identifier
// @Tags         SwiftCode
// @Accept       json
// @Produce      json
// @Param        swift-code path string true "Swift Code"
// @Success      200 {object} model.ApiResponse
// @Failure      400 {object} model.ApiResponse
// @Failure      404 {object} model.ApiResponse
// @Router       /v1/swift-codes/{swift-code} [get]
func (s *SwiftCodeHandlers) GetBySwiftCode(c *gin.Context) {
	s.HandlerLog.Info("Received request to fetch Swift Code")
	swiftCode := c.Param("swift-code")
	if !s.isValidSwiftCode(swiftCode) {
		s.HandlerLog.Info("Swift Code is malformed")
		c.JSON(http.StatusBadRequest, model.ApiResponse{Message: "Swift Code is Malformed"})
		return
	}

	resp, err := s.srvInst.GetBySwiftCode(swiftCode)
	if err != nil {
		s.HandlerLog.WithError(err).Error("Failed to fetch Swift Code")
		c.JSON(http.StatusNotFound, model.ApiResponse{Message: "Swift Code not found"})
		return
	}
	s.HandlerLog.Info("Swift Code fetched successfully")
	c.JSON(http.StatusOK, resp)
}

// GetByCountryISO2Code godoc
// @Summary      Fetch Swift Codes by Country ISO2 Code
// @Description  Retrieve all Swift Codes for a specific country using the country ISO2 code
// @Tags         SwiftCode
// @Accept       json
// @Produce      json
// @Param        countryISO2code path string true "Country ISO2 Code"
// @Success      200 {object}  []model.SwiftCodeDto
// @Failure      400 {object} model.ApiResponse
// @Failure      500 {object} model.ApiResponse
// @Router       /v1/swift-codes/country/{countryISO2code} [get]
func (s *SwiftCodeHandlers) GetByCountryISO2Code(c *gin.Context) {
	s.HandlerLog.Info("Received request to fetch Swift Codes by country ISO2")
	iso2Code := c.Param("countryISO2code")
	if !s.isValidISO2Code(iso2Code) {
		c.JSON(http.StatusBadRequest, model.ApiResponse{Message: "Invalid Country ISO2 Code"})
		return
	}

	resp, err := s.srvInst.GetByCountryISO(iso2Code)
	if err != nil {
		s.HandlerLog.WithError(err).Error("Failed to fetch Swift Codes for country")
		c.JSON(http.StatusInternalServerError, model.ApiResponse{Message: "Failed to fetch Swift Codes for country"})
		return
	}

	s.HandlerLog.Info("Swift Codes for country fetched successfully")
	c.JSON(http.StatusOK, resp)
}

// DeleteBySwiftCode godoc
// @Summary      Delete a Swift Code by its code
// @Description  Delete a Swift Code by providing the Swift Code identifier
// @Tags         SwiftCode
// @Accept       json
// @Produce      json
// @Param        swift-code path string true "Swift Code"
// @Success      200 {object} model.ApiResponse
// @Failure      400 {object} model.ApiResponse
// @Failure      500 {object} model.ApiResponse
// @Router       /v1/swift-codes/{swift-code} [delete]
func (s *SwiftCodeHandlers) DeleteBySwiftCode(c *gin.Context) {
	s.HandlerLog.Info("Received request to delete Swift Code")
	swiftCode := c.Param("swift-code")
	if !s.isValidSwiftCode(swiftCode) {
		s.HandlerLog.Info("Swift Code is malformed")
		c.JSON(http.StatusBadRequest, model.ApiResponse{Message: "Swift Code is Malformed"})
		return
	}

	err := s.srvInst.Delete(swiftCode)
	if errors.Is(err, custom_errors.ErrSwiftCodeNotFound) {
		s.HandlerLog.Info("Swift Code not found")
		c.JSON(http.StatusNotFound, model.ApiResponse{Message: "Swift Code not found"})
		return
	}

	if err != nil {
		s.HandlerLog.WithError(err).Error("Failed to delete Swift Code")
	}
	s.HandlerLog.Infof("Swift Code %s deleted successfully", swiftCode)
	c.JSON(http.StatusOK, model.ApiResponse{Message: "Swift Code deleted successfully"})
}

func (s *SwiftCodeHandlers) isAnyFieldNil(swiftCode *model.SwiftCodeDto) bool {
	if swiftCode == nil {
		return true
	}

	v := reflect.ValueOf(*swiftCode)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).String() == "" {
			return true
		}
	}
	return false
}

func (s *SwiftCodeHandlers) isValidSwiftCode(swiftCode string) bool {
	pattern := `^[A-Z]{4}[A-Z]{2}[A-Z0-9]{2}([A-Z0-9]{3})?$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(swiftCode)
}

func (s *SwiftCodeHandlers) isValidISO2Code(countryIso2Code string) bool {
	pattern := `[A-Z]{2}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(countryIso2Code)
}
