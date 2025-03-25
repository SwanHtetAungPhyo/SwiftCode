package validation

import (
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	"github.com/go-playground/validator/v10"
	"log"
)

var validate = validator.New()

func ValidateCreateRequest(request model.SwiftCodeAddRequest) error {
	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Println(err.Namespace(), err.Tag())
		}
		return err
	}
	return nil
}
