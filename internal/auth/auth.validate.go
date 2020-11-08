package auth

import (
	"errors"
	"github.com/6-things-must-to-do/server/internal/shared/utils/slice"
	"github.com/6-things-must-to-do/server/internal/shared/utils/validate"
)

func loginFormValidator(form *loginDto) error {
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
