package social

import "github.com/6-things-must-to-do/server/internal/shared/database/schema"

type Profile struct {
	schema.Profile
	Email string `json:"email"`
}
