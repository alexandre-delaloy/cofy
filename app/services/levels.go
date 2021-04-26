package services

import (
	"math"
)

func GetLevelFromXp(xp int) int {
	return int(0.1666 * math.Sqrt(float64(xp)))

}
