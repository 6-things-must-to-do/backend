package database

import (
	"errors"
	"fmt"
	"github.com/6-things-must-to-do/server/internal/shared/database/schema"
	sliceUtil "github.com/6-things-must-to-do/server/internal/shared/utils/slice"
	transformUtil "github.com/6-things-must-to-do/server/internal/shared/utils/transform"
	"strings"
)

func GetUserPKFromFollowerPK(followerPK string) string {
	uuid := strings.Split(followerPK, "#")[1]
	return GetUserPK(uuid)
}

func FollowFactory(uuid string) string {
	return fmt.Sprintf("FOLLOWER#%s", uuid)
}

func RequestFactory(uuid string) string {
	return fmt.Sprintf("REQ#FOLLOW#%s", uuid)
}

func GetCode(openSK string) int {
	code := strings.Split(openSK, "#")[2]
	return transformUtil.ToInt(code)
}

func GetOpennessCollection(opennessList *[]schema.Openness) *schema.OpennessCollection {
	opennessCollection := &schema.OpennessCollection{}
	for _, openness := range *opennessList {
		data := strings.Split(openness.SK, "#")
		opennessType := data[1]
		opennessCode := transformUtil.ToInt(data[2])
		switch opennessType {
		case "ACCOUNT":
			opennessCollection.Account = opennessCode
		case "TASK":
			opennessCollection.Task = opennessCode
		case "RECORD":
			opennessCollection.Record = opennessCode
		}
	}

	return opennessCollection
}

func OpenSKFactory(openType string, code int) string {
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

func GetUserPK(uuid interface{}) string {
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

func GetUUIDFromPK(pk string) string {
	return strings.Split(pk, "#")[1]
}

func GetEmailFromSK(sk string) string {
	return strings.Split(sk, "#")[1]
}
