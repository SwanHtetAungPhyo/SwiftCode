package handler

import (
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	"github.com/SwanHtetAungPhyo/swifcode/internal/pkg/converter"
	"github.com/SwanHtetAungPhyo/swifcode/internal/services"
	"github.com/SwanHtetAungPhyo/swifcode/internal/validation"
	"github.com/gofiber/fiber/v2"
)

// SwiftCodeHandlerInterface defines the methods to handle Swift code requests.
type SwiftCodeHandlerInterface interface {
	Get(c *fiber.Ctx) error
	GetWithISO2(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

// SwiftCodeHandlers holds the instance of service layer and implements the handler interface.
type SwiftCodeHandlers struct {
	serviceLayerInstance services.SwiftCodeServiceImpl
}

// NewSwiftCodeHandlers initializes the handlers with a service instance.
func NewSwiftCodeHandlers(serviceLayerInstance services.SwiftCodeServiceImpl) *SwiftCodeHandlers {
	return &SwiftCodeHandlers{serviceLayerInstance: serviceLayerInstance}
}

// Get retrieves the SwiftCode based on a query parameter.
func (impl *SwiftCodeHandlers) Get(c *fiber.Ctx) error {
	swiftCode := c.Query("swift-code")
	if swiftCode == "" || len(swiftCode) <= 11 || len(swiftCode) >= 8 {
		return handleErrorResponse(c, fiber.StatusBadRequest, "swift-code is required")
	}
	return nil
}

// GetWithISO2 retrieves SwiftCodes based on a country ISO2 code.
func (impl *SwiftCodeHandlers) GetWithISO2(c *fiber.Ctx) error {
	countryISO2 := c.Params("countryISO2code")
	if countryISO2 == "" {
		return handleErrorResponse(c, fiber.StatusBadRequest, "countryISO2code is required")
	}
	data := impl.serviceLayerInstance.GetWithISO2(countryISO2)
	return c.JSON(data)
}

// Create creates a new SwiftCode entry.
func (impl *SwiftCodeHandlers) Create(c *fiber.Ctx) error {
	var request model.SwiftCodeAddRequest
	if err := c.BodyParser(&request); err != nil {
		return handleErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}
	if err := validation.ValidateCreateRequest(request); err != nil {
		return handleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	newSwiftCode := converter.ConvertToSwiftCodeModel(request)
	impl.serviceLayerInstance.Create(newSwiftCode)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "SwiftCode created successfully",
	})
}

// Delete deletes a SwiftCode entry by SwiftCode value (if implemented).
func (impl *SwiftCodeHandlers) Delete(c *fiber.Ctx) error {
	return nil // This method would be implemented later.
}

// Helper function to handle errors and return consistent responses.
func handleErrorResponse(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"error": message,
	})
}
