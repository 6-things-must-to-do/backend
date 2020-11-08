package auth

import (
	"errors"
	"fmt"
	"github.com/6-things-must-to-do/server/internal/shared/configs"
	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type ServiceInterface interface {
	getOrCreateUser(p *loginDto) (*database.Profile, error)
	getJwtToken(pk string) string
}

type service struct {
	DB *database.DB
}

func (s *service) getOrCreateUser(p *loginDto) (*database.ProfileWithSetting, error) {
	ret := &database.ProfileWithSetting{}
	dtoAppId := database.CreateAppID(p.ID, p.Provider)

	err := s.DB.CoreTable.Get("SK", database.GetProfileSK(p.Email)).Index("Inverted").One(ret)
	if err == nil {
		ok, err := Compare(ret.AppID, dtoAppId)
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

	hashedAppId := hashAppId(dtoAppId)

	ret.Nickname = p.Nickname
	ret.ProfileImage = p.ProfileImage
	ret.AppID = hashedAppId
	ret.Provider = p.Provider
	ret.PK = database.GetUserPK(uid)
	ret.SK = database.GetProfileSK(p.Email)

	err = s.DB.CoreTable.Put(ret).Run()
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func hashAppId(appId string) string {
	fmt.Println(appId)
	bytes := []byte(appId)
	hash, err := bcrypt.GenerateFromPassword(bytes, bcrypt.MinCost)
	fmt.Println(string(hash))
	if err != nil {
		panic(err)
	}

	return string(hash)
}

func Compare(hash, appId string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(appId))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			// MEMO: err를 wrap 하여 상세를 전달하면 좋다
			return false, err
		}
		return false, err
	}
	return true, nil
}

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
