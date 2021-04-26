package services

import (
	"math"
)

func GetLevel(xp int) int64 {
	return int64(0.1666  * math.Sqrt(float64(xp)))

}