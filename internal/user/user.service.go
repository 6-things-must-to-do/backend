package user

import (
	"errors"

	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/6-things-must-to-do/server/internal/shared/database/schema"
	"github.com/guregu/dynamo"
)

// ServiceInterface ...
type ServiceInterface interface {
	getUserProfile(pk string, sk string)
}

func transformUserProfileFromProfileSchema(p *schema.ProfileSchema) *userProfile {
	up := &userProfile{
		Email:            database.GetEmailFromSK(p.SK),
		UUID:             database.GetUUIDFromPK(p.PK),
		ProfileImage:     p.ProfileImage,
		Nickname:         p.Nickname,
		TaskAlertSetting: p.TaskAlertSetting,
	}

	return up
}

func (s *service) getUserProfile(pk string) (*userProfile, error) {
	profile := &schema.ProfileSchema{}

	err := s.DB.CoreTable.Get("PK", pk).Range("SK", dynamo.BeginsWith, "PROFILE#").One(profile)
	if err != nil {
		return nil, errors.New("user not found")
	}
	up := transformUserProfileFromProfileSchema(profile)

	return up, nil
}

type service struct {
	DB *database.DB
}

func newService(DB *database.DB) *service {
	return &service{DB: DB}
}
