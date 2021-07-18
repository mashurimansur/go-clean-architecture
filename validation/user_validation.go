package validation

import (
	"clean-arch-go/model"

	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidateUser(request model.UserRequest) error {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Age, validation.Required, validation.Min(20)),
	)

	if err != nil {
		return err
	}

	return nil
}
