package middlewares

import (
	"errors"
	"github.com/6-things-must-to-do/server/internal/shared"
	"github.com/6-things-must-to-do/server/internal/shared/configs"
	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/guregu/dynamo"
	"net/http"
	"strings"
)

func extractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	token := strings.Split(bearerToken, " ")[1]
	return token
}

func tokenValid(tokenString string) (*jwt.Token, error) {
	errInvalid := errors.New("invalid token")

	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errInvalid
		}
		return []byte(configs.GetConfig().SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errInvalid
	}

	return token, nil
}

func extractPK(token *jwt.Token) string {
	claims, _ := token.Claims.(jwt.MapClaims)
	pk, ok := claims["pk"].(string)
	if !ok {
		return ""
	}
	return pk
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := extractToken(c.Request)
		if tokenString == "" {
			shared.UnAuthenticatedError(c, errors.New("cannot extract token").Error())
			return
		}

		token, err := tokenValid(tokenString)
		if err != nil {
			shared.UnAuthenticatedError(c, err.Error())
			return
		}

		pk := extractPK(token)
		if pk == "" {
			shared.UnAuthenticatedError(c, errors.New("pk claims not found").Error())
			return
		}

		profile := &database.Profile{}
		db := database.GetDB()
		err = db.CoreTable.Get("PK", pk).Range("SK", dynamo.BeginsWith, "PROFILE#").One(profile)

		if err != nil {
			shared.UnAuthenticatedError(c, errors.New("cannot found user").Error())
			return
		}

		c.Set("User", profile)
		c.Next()
	}
}

func GetUserProfile(c *gin.Context) *database.Profile {
	user, ok := c.Get("User")
	if !ok {
		panic(errors.New("not authenticated"))
	}

	var profile *database.Profile
	profile, ok = user.(*database.Profile)
	if !ok {
		panic(errors.New("invalid user data"))
	}

	return profile
}