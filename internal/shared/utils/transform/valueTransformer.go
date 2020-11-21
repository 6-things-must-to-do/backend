package transformUtil

import (
	"math"
)

func GetRecordPercent(inComplete int, complete int) float64 {
	val := 1 / float64(3) * 100
	return math.Round(val*100) / 100
}