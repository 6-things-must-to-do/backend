package transformUtil

import (
	"reflect"
	"strconv"
	"time"
)

func strToInt (str string) int {
	ret, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return ret
}

func ToInt(n interface{}) int {
	t := reflect.TypeOf(n).Name()
	var ret int
	switch t {
	case "string":
		ret = strToInt(n.(string))
	}
	return ret
}

func GetTimeFromJSUnixTimestamp(jsUnixTimestamp int64) time.Time {
	t := time.Unix(ToUnixTimestamp(jsUnixTimestamp), 0)
	return t
}

func GetJSUnixTimestampFromTime(t time.Time) int64 {
	return ToJSUnixTimestamp(t.Unix())
}

func ToUnixTimestamp(t int64) int64 {
	return t / 1000
}

func ToJSUnixTimestamp(t int64) int64 {
	jsUnixTimestamp := t * 1000
	return jsUnixTimestamp
}