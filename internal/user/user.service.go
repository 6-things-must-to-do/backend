package user

import (
	"errors"
	transformUtil "github.com/6-things-must-to-do/server/internal/shared/utils/transform"
	"strings"

	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/6-things-must-to-do/server/internal/shared/database/schema"
	"github.com/guregu/dynamo"
)

type service struct {
	DB *database.DB
}

func newService(DB *database.DB) *service {
	return &service{DB: DB}
}

func transformUserProfileFromProfileSchema(p *schema.ProfileSchema) *Profile {
	profile := &Profile{
		Email:            database.GetEmailFromSK(p.SK),
		UUID:             database.GetUUIDFromPK(p.PK),
		ProfileImage:     p.ProfileImage,
		Nickname:         p.Nickname,
	}

	return profile
}

func getPermissionStatus(opennessList *[]schema.Openness) *map[string]int {
	ret := map[string]int{}
	for _, open := range *opennessList {
		sk := strings.Split(open.SK, "#")
		var openType = strings.ToLower(sk[1])
		var openCode = transformUtil.ToInt(sk[2])
		ret[openType] = openCode
	}
	return &ret
}

func (s *service) removeUser (userPK string) error {
	// TODO 에러
	return s.DB.CoreTable.Delete("PK", userPK).Run()
}

func (s *service) getUserOpenness(userPK string) (*map[string]int, error) {
	var opennessList []schema.Openness
	err := s.DB.CoreTable.Get("PK", userPK).Range("SK", dynamo.BeginsWith, "OPEN#").All(&opennessList)
	if err != nil {
		return nil, err
	}

	return getPermissionStatus(&opennessList), nil
}

func (s *service) setTaskAlert(user *ProfileWithSetting) {
	m, err := dynamo.MarshalItem(user)
	if err != nil {
		panic(err)
	}

	err = s.DB.CoreTable.Update("PK", m["PK"]).Range("SK", m["SK"]).Set("TaskAlertSetting", m["TaskAlertSetting"]).Run()
	if err != nil {
		panic(err)
	}
}

func (s *service) getUserProfile(pk string) (*Profile, error) {
	profile := &schema.ProfileSchema{}

	err := s.DB.CoreTable.Get("PK", pk).Range("SK", dynamo.BeginsWith, "PROFILE#").One(profile)
	if err != nil {
		return nil, errors.New("user not found")
	}
	up := transformUserProfileFromProfileSchema(profile)

	return up, nil
}
