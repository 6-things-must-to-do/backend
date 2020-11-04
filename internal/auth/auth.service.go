package auth

import (
	"github.com/6-things-must-to-do/server/internal/shared/configs"
	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"github.com/guregu/dynamo"
	"golang.org/x/crypto/bcrypt"
)

type ServiceInterface interface {
	getOrCreateUser(p *getOrCreateUserParam) (*database.Profile, error)
	getJwtToken(pk string) string
}

type service struct {
	DB *database.DB
}

type getOrCreateUserParam struct {
	appId string
	email string
	provider string
	nickname string
	profileImage string
}

func (s *service) getOrCreateUser(p *getOrCreateUserParam) (*database.Profile, error) {
	ret := &database.Profile{}
	err := s.DB.CoreTable.Get("AppID", hashAppId(p.appId)).Index("AppID").Range("SK", dynamo.Equal, database.GetProfileSK(p.email)).One(ret)
	if err == nil {
		return ret, nil
	}

	uid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	ret.Nickname = p.nickname
	ret.ProfileImage = p.profileImage
	ret.AppID = database.CreateAppID(p.appId, p.provider)
	ret.Provider = p.provider
	ret.PK = database.GetUserPK(uid)
	ret.SK = database.GetProfileSK(p.email)

	err = s.DB.CoreTable.Put(ret).Run()
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func hashAppId(appId string) string {
	bytes := []byte(appId)
	hash, err := bcrypt.GenerateFromPassword(bytes, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	return string(hash)
}

type JwtClaims struct {
	PK string `json:"pk"`
	jwt.StandardClaims
}

func (s *service) getJwtToken(pk string) string {
	secret := []byte(configs.GetConfig().SECRET)
	claims := JwtClaims{PK:pk}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		panic(err)
	}

	return token
}

func newService(DB *database.DB) *service {
	return &service{DB:DB}
}
