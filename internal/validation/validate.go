package validation

import (
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	"github.com/SwanHtetAungPhyo/swifcode/internal/pkg/logging"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

var validate = validator.New()

func ValidateCreateRequest(request model.SwiftCodeAddRequest) error {
	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			logging.Logger.Info("Validation Error",
				zap.String("Field", err.Field()),
				zap.String("Tag", err.Tag()),
				zap.String("Param", err.Param()))
		}
		return err
	}
	return nil
}
