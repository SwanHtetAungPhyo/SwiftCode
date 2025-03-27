package swiftcode_test

import (
	"github.com/SwanHtetAungPhyo/swifcode/internal/routes"
	"github.com/SwanHtetAungPhyo/swifcode/mocks"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSetUpRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	log := logrus.New()

	mockHandlers := new(mocks.MockSwiftCodeHandlers)
	mockHandlers.On("Create", mock.Anything).Return()
	mockHandlers.On("GetBySwiftCode", mock.Anything).Return()
	mockHandlers.On("GetByCountryISO2Code", mock.Anything).Return()
	mockHandlers.On("DeleteBySwiftCode", mock.Anything).Return()

	routes.SetUpRoute(router, mockHandlers, log)

	require.NotNil(t, router)
}
