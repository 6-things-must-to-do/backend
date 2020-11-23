package transformUtil

import (
	"math"
)

func GetRecordPercent(inComplete int, complete int) float64 {
	val := float64(complete) / float64(inComplete + complete) * 100
	return math.Round(val*100) / 100
}