package transformUtil

import (
	"reflect"
	"strconv"
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

func ToUnixTimestamp(t int64) int64 {
	goUnixTimestamp := t / 1000;
	return goUnixTimestamp
}

func ToJSUnixTimestamp(t int64) int64 {
	jsUnixTimestamp := t * 1000;
	return jsUnixTimestamp
}