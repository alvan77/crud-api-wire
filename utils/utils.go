package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
)

func GetJsonTag(validate *validator.Validate) {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
}

func ErrorValidationMessage(err error) string {
	var data []string
	for _, err := range err.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required", "required_without_all", "required_without":
			data = append(data, fmt.Sprintf("%s is required", err.Field()))
		case "email":
			data = append(data, fmt.Sprintf("%s is not valid email", err.Field()))
		case "gte":
			data = append(data, fmt.Sprintf("%s value must be greater than %s", err.Field(), err.Param()))
		case "lte":
			data = append(data, fmt.Sprintf("%s value must be lower than %s", err.Field(), err.Param()))
		case "oneof":
			data = append(data, fmt.Sprintf("%s value unknown", err.Field()))
		default:
			data = append(data, "")
		}
	}

	return strings.Join(data, ", ")
}
