package social

import (
	"errors"
	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/6-things-must-to-do/server/internal/shared/database/schema"
	"github.com/guregu/dynamo"
)

type service struct {
	DB *database.DB
}

func getTargetProfileSchema(table *dynamo.Table, email string) (*schema.ProfileSchema, error) {
	var targetProfile schema.ProfileSchema
	err := table.
		Get("SK", database.GetProfileSK(email)).
		Range("PK", dynamo.BeginsWith, "USER#").
		Index("Inverted").
		One(&targetProfile)

	if err != nil {
		return nil, err
	}

	return &targetProfile, nil
}

func getAccountOpennessCode(table *dynamo.Table, userPK string) (int, error) {
	var accountOpenness schema.Openness

	err := table.
		Get("PK", userPK).
		Range("SK", dynamo.BeginsWith, "OPEN#ACCOUNT#").
		One(&accountOpenness)
	if err != nil {
		return 0, err
	}

	code := database.GetCode(accountOpenness.SK)

	return code, nil
}

func (s *service) follow(userPK, targetEmail string) error {
	targetProfile, err := getTargetProfileSchema(&s.DB.CoreTable, targetEmail)
	if err != nil {
		return err
	}

	code, err := getAccountOpennessCode(&s.DB.CoreTable, targetProfile.PK)
	if err != nil {
		return err
	}

	uuid := database.GetUUIDFromPK(userPK)

	switch code {
	case 0: // 친구 추가 불가능
		return errors.New("account is closed")
	case 1: // 친구 요청
		request := &schema.Request{
			PK: database.RequestFactory(uuid),
			SK: database.GetProfileSK(targetEmail),
		}
		err = s.DB.CoreTable.Put(request).Run()
	case 2: // 바로 팔로우
		follow := &schema.Follow{
			PK: database.FollowFactory(uuid),
			SK: database.GetProfileSK(targetEmail),
		}
		err = s.DB.CoreTable.Put(follow).Run()
	}

	if err != nil {
		return err
	}

	return nil
}

func (s *service) getUser(targetEmail string) (*Profile, error) {
	targetProfile, err := getTargetProfileSchema(&s.DB.CoreTable, targetEmail)
	if err != nil {
		return nil, err
	}

	code, err := getAccountOpennessCode(&s.DB.CoreTable, targetProfile.PK)
	if err != nil {
		return nil, err
	}

	if code < 1{
		return nil, errors.New("user not found")
	}

	socialProfile := &Profile{
		Profile: schema.Profile{
			Provider:     targetProfile.Provider,
			ProfileImage: targetProfile.ProfileImage,
			Nickname:     targetProfile.Nickname,
		},
		Email:   targetEmail,
	}

	return socialProfile, nil
}

func newService(DB *database.DB) *service {
	return &service{DB: DB}
}
