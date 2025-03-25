package test

import (
	"github.com/SwanHtetAungPhyo/swifcode/app/internal/model"
	"github.com/SwanHtetAungPhyo/swifcode/app/internal/validation"
	"github.com/SwanHtetAungPhyo/swifcode/internal/pkg/log"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"testing"
)

func TestValidateCreateRequest(t *testing.T) {
	log.Logger = zaptest.NewLogger(t)

	type test struct {
		name     string
		input    model.SwiftCodeAddRequest
		expected bool
	}
	tests := make([]test, 2)
	tests[0] = test{
		name: "valid request",
		input: model.SwiftCodeAddRequest{
			Address:       "US , America",
			SwiftCode:     "US123456",
			BankName:      "Bank of America",
			CountryISO2:   "US",
			CountryName:   "United States",
			IsHeadquarter: true,
		},
		expected: true,
	}
	tests[1] = test{
		name: "invalid request - incorrect Swift Code length",
		input: model.SwiftCodeAddRequest{
			Address:       "US , America",
			SwiftCode:     "US126", // ❌ Invalid: Less than 8 characters
			BankName:      "Bank of America",
			CountryISO2:   "US",
			CountryName:   "United States",
			IsHeadquarter: true,
		},
		expected: false,
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validation.ValidateCreateRequest(tc.input)
			isValid := err == nil
			assert.Equal(t, tc.expected, isValid)
		})
	}

}
