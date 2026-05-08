package validate

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func PasswordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// อนุญาตเฉพาะอังกฤษ ตัวเลข และ symbol พื้นฐาน
	regex := regexp.MustCompile(`^[\x20-\x7E]+$`)

	return regex.MatchString(password)
}