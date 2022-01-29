package utils

import (
	"hera/app/models"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/liip/sheriff"
	"github.com/thoas/go-funk"
)

// NewValidator func for create a new validator for model fields.
func NewValidator() *validator.Validate {
	validate := validator.New()

	// Custom validation for uuid.UUID fields.
	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})

	// Custom validation category
	_ = validate.RegisterValidation("category", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		categories := []string{
			models.ArticleCategorySport,
			models.ArticleCategoryTech,
			models.ArticleCategoryBusiness,
		}

		return !funk.ContainsString(categories, field)

	})

	return validate

}

// ValidatorErrors func for show validation errors for each invalid fields.
func ValidatorErrors(err error) map[string]string {
	// Define fields map.
	fields := map[string]string{}

	// Make error message for each invalid field.
	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = err.Error()
	}

	return fields
}

func MarshalUsers(data interface{}, groups ...string) (interface{}, error) {
	o := &sheriff.Options{
		Groups: groups,
	}

	data, err := sheriff.Marshal(o, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
