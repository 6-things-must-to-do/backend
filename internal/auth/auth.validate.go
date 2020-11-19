package auth

import (
	"errors"
	sliceUtil "github.com/6-things-must-to-do/server/internal/shared/utils/slice"
	validateUtil "github.com/6-things-must-to-do/server/internal/shared/utils/validate"
)

func loginFormValidator(form *loginDto) error {
	// Check available
	available := []string{"google", "apple"}
	isIncluded := sliceUtil.Includes(available, form.Provider)
	if !isIncluded {
		return errors.New("unsupported provider")
	}

	// Check email validation
	if !validateUtil.IsEmail(form.Email) {
		return errors.New("invalid email form")
	}
	return nil
}
