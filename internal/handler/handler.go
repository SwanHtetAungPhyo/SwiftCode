package handler

import (
	"github.com/gin-gonic/gin"
)

// Methods defines HTTP handler methods for managing resources.
type Methods interface {
	// Create handles resource creation (POST request).
	Create(c *gin.Context)

	// GetBySwiftCode retrieves a resource by SWIFT code (GET request).
	GetBySwiftCode(c *gin.Context)

	// GetByCountryISO2Code retrieves a resource by country ISO2 code (GET request).
	GetByCountryISO2Code(c *gin.Context)

	// DeleteBySwiftCode deletes a resource by SWIFT code (DELETE request).
	DeleteBySwiftCode(c *gin.Context)
}
