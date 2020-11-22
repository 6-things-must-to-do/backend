package auth

import (
	"errors"
	"github.com/guregu/dynamo"

	"github.com/6-things-must-to-do/server/internal/shared/configs"
	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/6-things-must-to-do/server/internal/shared/database/schema"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	DB *database.DB
}

func (s *service) getOrCreateUser(p *loginDto) (*schema.ProfileSchema, error) {
	user := &schema.ProfileSchema{}
	dtoAppID := database.CreateAppID(p.ID, p.Provider)

	err := s.DB.CoreTable.
		Get("SK", database.GetProfileSK(p.Email)).
		Range("PK", dynamo.BeginsWith, "USER#").
		Index("Inverted").
		One(user)

	if err == nil {
		ok, err := Compare(user.AppID, dtoAppID)
		if err != nil {
			return nil, err
		}

		if !ok {
			return nil, errors.New("this email was already used with other provider")
		}

		return user, nil
	}

	uid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	hashed := hashAppID(dtoAppID)
	userPK := database.GetUserPK(uid)

	user.Nickname = p.Nickname
	user.ProfileImage = p.ProfileImage
	user.AppID = hashed
	user.Provider = p.Provider
	user.PK = userPK
	user.SK = database.GetProfileSK(p.Email)

	accountOpenness := &schema.Openness{
		PK: userPK,
		SK: database.OpenSKFactory("ACCOUNT", 2),
	}

	recordOpenness := &schema.Openness{
		PK: userPK,
		SK: database.OpenSKFactory("RECORD", 2),
	}

	taskOpenness := &schema.Openness{
		PK: userPK,
		SK: database.OpenSKFactory("TASK", 2),
	}

	_, err = s.DB.CoreTable.Batch().
		Write().
		Put(user, accountOpenness, recordOpenness, taskOpenness).
		Run()

	if err != nil {
		return nil, err
	}

	return user, nil
}

func hashAppID(appID string) string {
	bytes := []byte(appID)
	hash, err := bcrypt.GenerateFromPassword(bytes, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	return string(hash)
}

// Compare ...
func Compare(hash, appID string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(appID))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			// MEMO: err를 wrap 하여 상세를 전달하면 좋다
			return false, err
		}
		return false, err
	}
	return true, nil
}

// JwtClaims ...
type JwtClaims struct {
	PK string `json:"pk"`
	jwt.StandardClaims
}

func (s *service) getJwtToken(pk string) string {
	secret := []byte(configs.GetConfig().SECRET)
	claims := JwtClaims{PK: pk}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		panic(err)
	}

	return token
}

func newService(DB *database.DB) *service {
	return &service{DB: DB}
}
