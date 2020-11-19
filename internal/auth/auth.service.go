package auth

import (
	"errors"
	"fmt"

	"github.com/6-things-must-to-do/server/internal/shared/configs"
	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/6-things-must-to-do/server/internal/shared/database/schema"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type serviceInterface interface {
	getOrCreateUser(p *loginDto) (*schema.ProfileWithSetting, error)
	getJwtToken(pk string) string
}

type service struct {
	DB *database.DB
}

func (s *service) getOrCreateUser(p *loginDto) (*schema.ProfileSchema, error) {
	ret := &schema.ProfileSchema{}
	dtoAppID := database.CreateAppID(p.ID, p.Provider)

	err := s.DB.CoreTable.Get("SK", database.GetProfileSK(p.Email)).Index("Inverted").One(ret)
	if err == nil {
		ok, err := Compare(ret.AppID, dtoAppID)
		if err != nil {
			return nil, err
		}

		if !ok {
			return nil, errors.New("this email was already used with other provider")
		}

		return ret, nil
	}

	uid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	hashed := hashAppID(dtoAppID)

	ret.Nickname = p.Nickname
	ret.ProfileImage = p.ProfileImage
	ret.AppID = hashed
	ret.Provider = p.Provider
	ret.PK = database.GetUserPK(uid)
	ret.SK = database.GetProfileSK(p.Email)

	err = s.DB.CoreTable.Put(ret).Run()
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func hashAppID(appID string) string {
	fmt.Println(appID)
	bytes := []byte(appID)
	hash, err := bcrypt.GenerateFromPassword(bytes, bcrypt.MinCost)
	fmt.Println(string(hash))
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
