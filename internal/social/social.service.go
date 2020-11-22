package social

import (
	"errors"
	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/6-things-must-to-do/server/internal/shared/database/schema"
	"github.com/6-things-must-to-do/server/internal/user"
	"github.com/guregu/dynamo"
)

type service struct {
	DB *database.DB
}

func schemaToProfile(schemaList *[]schema.ProfileSchema) *[]user.Profile {
	var list []user.Profile
	for _, sc := range *schemaList {
		profile := user.Profile{
			Email:        database.GetEmailFromSK(sc.SK),
			UUID:         database.GetUUIDFromPK(sc.PK),
			ProfileImage: sc.ProfileImage,
			Nickname:     sc.Nickname,
		}

		list = append(list, profile)
	}
	return &list
}

func (s *service) getFollowingList(userPK string) (*[]user.Profile, error) {
	uuid := database.GetUUIDFromPK(userPK)
	var followingSchema []schema.Follow
	err := s.DB.CoreTable.
		Get("PK", database.FollowFactory(uuid)).
		Range("SK", dynamo.BeginsWith, "PROFILE#").
		All(&followingSchema)
	if err != nil {
		return nil, err
	}

	list := make([]user.Profile, 0)
	if len(followingSchema) == 0 {
		return &list, nil
	}

	var items []dynamo.Keyed

	for _, sc := range followingSchema {
		item := dynamo.Keys{database.GetUserPK(sc.ProfileUUID), sc.SK}
		items = append(items, item)
	}

	var followingProfileList []schema.ProfileSchema

	err = s.DB.CoreTable.
		Batch("PK", "SK").
		Get(items...).
		All(&followingProfileList)

	if err != nil {
		return nil, err
	}

	list = *schemaToProfile(&followingProfileList)
	return &list, nil
}

func (s *service) getFollowerList (profileSK string) (*[]user.Profile, error) {
	var followerSchema []schema.Follow

	err := s.DB.CoreTable.
		Get("SK", profileSK).
		Range("PK", dynamo.BeginsWith, "FOLLOWER#").
		Index("Inverted").
		All(&followerSchema)
	if err != nil {
		return nil, err
	}

	list := make([]user.Profile, 0)
	if len(followerSchema) == 0 {
		return &list, nil
	}

	var items []dynamo.Keyed

	for _, sc := range followerSchema {
		item := dynamo.Keys{
			database.GetUserPKFromFollowerPK(sc.PK),
			database.GetProfileSK(sc.FollowerEmail),
		}
		items = append(items, item)
	}

	var followerProfileList []schema.ProfileSchema
	err = s.DB.CoreTable.
		Batch("PK", "SK").
		Get(items...).
		All(&followerProfileList)
	if err != nil {
		return nil, err
	}

	list = *schemaToProfile(&followerProfileList)

	return &list, nil
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

func (s *service) unfollow(userPK string, targetEmail string) error {
	return s.DB.CoreTable.
		Delete("PK", database.GetFollowerPKFromUserPK(userPK)).
		Range("SK", database.GetProfileSK(targetEmail)).Run()
}

func (s *service) follow(userPK, profileSK string, targetEmail string) error {
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
			ProfileUUID: database.GetUUIDFromPK(targetProfile.PK),
			FollowerEmail: database.GetEmailFromSK(profileSK),
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
