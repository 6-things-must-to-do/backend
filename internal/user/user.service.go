package user

import (
	"errors"
	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/guregu/dynamo"
)

type ServiceInterface interface {
	getUserProfile(pk string, sk string)
}

func transformUserProfileFromProfileSchema(p *database.Profile) *userProfile {
	up := &userProfile{
		Email:        database.GetEmailFromSK(p.SK),
		UUID:         database.GetUUIDFromPK(p.PK),
		ProfileImage: p.ProfileImage,
		Nickname:     p.Nickname,
	}

	return up
}

func (s *service) getUserProfile (pk string) (*userProfile, error) {
	profile := &database.Profile{}

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
