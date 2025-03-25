package handler

import (
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	"github.com/SwanHtetAungPhyo/swifcode/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Methods interface {
	Create(c *gin.Context)
	GetBySwiftCode(c *gin.Context)
	GetByCountryISO2Code(c *gin.Context)
	DeleteBySwiftCode(c *gin.Context)
}

type SwiftCodeHandlers struct {
	srvInst    *services.SwiftCodeServices
	handlerLog *logrus.Logger
}

func NewSwiftCodeHandlers(ser *services.SwiftCodeServices, handlerLog *logrus.Logger) *SwiftCodeHandlers {
	return &SwiftCodeHandlers{
		srvInst:    ser,
		handlerLog: handlerLog,
	}
}

func (s *SwiftCodeHandlers) Create(c *gin.Context) {
	s.handlerLog.Info("Received request to create Swift Code")
	var swiftCode model.SwiftCodeDto
	if err := c.ShouldBindJSON(&swiftCode); err != nil {
		s.handlerLog.WithError(err).Error("Failed to bind Swift Code request")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request format"})
		return
	}

	if err := s.srvInst.Create(&swiftCode); err != nil {
		s.handlerLog.WithError(err).Error("Failed to create Swift Code")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create Swift Code"})
		return
	}

	s.handlerLog.Info("Swift Code created successfully")
	c.JSON(http.StatusCreated, gin.H{"message": "Swift Code created successfully"})
}

func (s *SwiftCodeHandlers) GetBySwiftCode(c *gin.Context) {
	s.handlerLog.Info("Received request to fetch Swift Code")
	swiftCode := c.Param("swift-code")
	if swiftCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Swift Code is required"})
		return
	}

	resp, err := s.srvInst.GetBySwiftCode(swiftCode)
	if err != nil {
		s.handlerLog.WithError(err).Error("Failed to fetch Swift Code")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch Swift Code"})
		return
	}

	s.handlerLog.Info("Swift Code fetched successfully")
	c.JSON(http.StatusOK, resp)
}

func (s *SwiftCodeHandlers) GetByCountryISO2Code(c *gin.Context) {
	s.handlerLog.Info("Received request to fetch Swift Codes by country ISO2")
	iso2Code := c.Param("countryISO2code")
	if iso2Code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Country ISO2 Code is required"})
		return
	}

	resp, err := s.srvInst.GetByCountryISO(iso2Code)
	if err != nil {
		s.handlerLog.WithError(err).Error("Failed to fetch Swift Codes for country")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch Swift Codes for country"})
		return
	}

	s.handlerLog.Info("Swift Codes for country fetched successfully")
	c.JSON(http.StatusOK, resp)
}

func (s *SwiftCodeHandlers) DeleteBySwiftCode(c *gin.Context) {
	s.handlerLog.Info("Received request to delete Swift Code")
	swiftCode := c.Param("swift-code")
	if swiftCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Swift Code is required"})
		return
	}

	if err := s.srvInst.Delete(swiftCode); err != nil {
		s.handlerLog.WithError(err).Error("Failed to delete Swift Code")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete Swift Code"})
		return
	}

	s.handlerLog.Infof("Swift Code %s deleted successfully", swiftCode)
	c.JSON(http.StatusOK, gin.H{
		"message":      "Swift Code deleted successfully",
		"deleted_code": swiftCode,
	})
}
