package database

import (
	"fmt"
	"github.com/gofrs/uuid"
	"strings"
)

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
