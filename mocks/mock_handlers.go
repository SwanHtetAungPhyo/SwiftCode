package mocks

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockSwiftCodeHandlers struct {
	mock.Mock
}

func (m *MockSwiftCodeHandlers) Create(c *gin.Context) {
	m.Called(c)
}

func (m *MockSwiftCodeHandlers) GetBySwiftCode(c *gin.Context) {
	m.Called(c)
}

func (m *MockSwiftCodeHandlers) GetByCountryISO2Code(c *gin.Context) {
	m.Called(c)
}

func (m *MockSwiftCodeHandlers) DeleteBySwiftCode(c *gin.Context) {
	m.Called(c)
}
