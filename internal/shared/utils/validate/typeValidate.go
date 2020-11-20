package validateUtil

import (
	"reflect"
)

func IsInt(n interface {}) bool {
	t := reflect.TypeOf(n).Name()
	return t == "int"
}