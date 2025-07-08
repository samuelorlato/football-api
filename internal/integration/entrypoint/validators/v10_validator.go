package validators

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/samuelorlato/football-api/internal/integration/ports"
)

type v10Validator struct {
	instance *validator.Validate
}

func NewV10Validator() ports.Validator {
	return &v10Validator{
		instance: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (v *v10Validator) Struct(obj interface{}) error {
	return v.instance.Struct(obj)
}

type ValidationErrorResponse struct {
	Field   string `json:"campo"`
	Message string `json:"mensagem"`
}

func (v *v10Validator) GetErrors(err error, input interface{}) interface{} {
	var errors []ValidationErrorResponse

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		t := reflect.TypeOf(input)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}

		for _, e := range validationErrors {
			field := e.StructField()
			jsonField := field

			if f, ok := t.FieldByName(field); ok {
				tag := f.Tag.Get("json")
				if tag != "" && tag != "-" {
					jsonField = strings.Split(tag, ",")[0]
				}
			}

			errors = append(errors, ValidationErrorResponse{
				Field:   jsonField,
				Message: validationErrorMessage(jsonField, e),
			})
		}
	}
	return errors
}

func validationErrorMessage(field string, e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("o campo %s é obrigatório", field)
	case "email":
		return fmt.Sprintf("o campo %s deve ser um email válido", field)
	default:
		return fmt.Sprintf("o campo %s é inválido", field)
	}
}
