package auth

import (
	"errors"
	"fmt"
	"github.com/6-things-must-to-do/server/internal/shared/utils/slice"
	"github.com/6-things-must-to-do/server/internal/shared/utils/validate"
)

func loginFormValidator(form *loginDto) error {
	required := map[string]string{"nickname": form.Nickname, "provider": form.Provider, "email": form.Email, "id": form.ID}

	// Check Required Field/
	for field, value := range required {
		if value == "" {
			errString := fmt.Sprintf("%s is required", field)
			return errors.New(errString)
		}
	}

	// Check available
	available := []string{"google", "apple"}
	isIncluded := slice.Includes(available, form.Provider)
	if !isIncluded {
		return errors.New("unsupported provider")
	}

	// Check email validation
	if !validate.IsEmail(form.Email) {
		return errors.New("invalid email form")
	}

	return nil
}
