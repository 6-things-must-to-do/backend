package database

import (
	"errors"
	"fmt"
	sliceUtil "github.com/6-things-must-to-do/server/internal/shared/utils/slice"
	"github.com/gofrs/uuid"
	"strings"
)

func GetOpenSK(openType string, code int) string {
	availableCode := []interface{}{1, 2, 3}
	availableType := []interface{}{"ACCOUNT", "RECORD", "TASK"}
	if !sliceUtil.Includes(availableCode, code) {
		panic(errors.New("invalid open code"))
	}

	if !sliceUtil.Includes(availableType, openType) {
		panic(errors.New("invalid open type"))
	}

	ret := fmt.Sprintf("OPEN#%s#%d",openType, code)
	return ret
}

func GetUserPK(uuid uuid.UUID) string {
	ret := fmt.Sprintf("USER#%s", uuid)
	return ret
}

func GetRecordSK(lockTime int64) string {
	ret := fmt.Sprintf("RECORD#%d", lockTime)
	return ret
}

func GetTaskSK(index int) string {
	ret := fmt.Sprintf("TASK#%d", index)
	return ret
}

func GetProfileSK(email string) string {
	ret := fmt.Sprintf("PROFILE#%s", email)
	return ret
}

func CreateAppID(provider string, id string) string {
	ret := fmt.Sprintf("%s|%s", provider, id)
	return ret
}

func CreateUserPK(uuid string) string {
	ret := fmt.Sprintf("USER#%s", uuid)
	return ret
}

func GetUUIDFromPK(pk string) string {
	return strings.Split(pk, "#")[1]
}

func GetEmailFromSK(sk string) string {
	return strings.Split(sk, "#")[1]
}
