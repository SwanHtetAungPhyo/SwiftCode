package handler

import (
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	"reflect"
	"regexp"
)

// SwiftCode Validation
func (s *SwiftCodeHandlers) isValidSwiftCode(swiftCode string) bool {
	pattern := `^[A-Z]{4}[A-Z]{2}[A-Z0-9]{2}([A-Z0-9]{3})?$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(swiftCode)
}

// iso2Code Validation
func (s *SwiftCodeHandlers) isValidISO2Code(countryIso2Code string) bool {
	pattern := `[A-Z]{2}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(countryIso2Code)
}

// Nil input prevention
func (s *SwiftCodeHandlers) isAnyFieldNil(swiftCode *model.SwiftCodeDto) bool {
	if swiftCode == nil {
		return true
	}

	v := reflect.ValueOf(*swiftCode)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		switch field.Kind() {
		case reflect.String:
			if field.String() == "" {
				return true
			}
		case reflect.Ptr, reflect.Interface:
			if field.IsNil() {
				return true
			}
		case reflect.Slice, reflect.Array:
			if field.Len() == 0 {
				return true
			}
		default:
			return false
		}
	}
	return false
}
