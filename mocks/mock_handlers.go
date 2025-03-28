package mocks

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

// MockSwiftCodeHandlers is a mock implementation of the SwiftCodeHandlers interface.
// It is used for testing purposes to simulate handler behavior.
type MockSwiftCodeHandlers struct {
	mock.Mock
}

// Create simulates the Create handler method.
// It records the call and allows assertions on its behavior during tests.
func (m *MockSwiftCodeHandlers) Create(c *gin.Context) {
	m.Called(c)
}

// GetBySwiftCode simulates the GetBySwiftCode handler method.
// It records the call and allows assertions on its behavior during tests.
func (m *MockSwiftCodeHandlers) GetBySwiftCode(c *gin.Context) {
	m.Called(c)
}

// GetByCountryISO2Code simulates the GetByCountryISO2Code handler method.
// It records the call and allows assertions on its behavior during tests.
func (m *MockSwiftCodeHandlers) GetByCountryISO2Code(c *gin.Context) {
	m.Called(c)
}

// DeleteBySwiftCode simulates the DeleteBySwiftCode handler method.
// It records the call and allows assertions on its behavior during tests.
func (m *MockSwiftCodeHandlers) DeleteBySwiftCode(c *gin.Context) {
	m.Called(c)
}
