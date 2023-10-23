package shared

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	validate := validator.New()

	// Đăng ký hàm kiểm tra tùy chỉnh "customPhoneCheck"
	_ = validate.RegisterValidation("customPhoneCheck", customPhoneCheck)

	return &Validator{
		validate: validate,
	}
}

func (v *Validator) Validate(data interface{}) error {
	return v.validate.Struct(data)
}

func customPhoneCheck(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()

	phoneDigits := strings.Count(strings.ReplaceAll(phoneNumber, " ", ""), "1234567890")
	return phoneDigits > 10
}
