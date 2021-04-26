package helpers

import (
	"math"
)

func GetLevelFromXp(xp int) int {
	return int(0.1666 * math.Sqrt(float64(xp)))
}
func GetXpFromLevel(level int) float64 {
	x := float64(level) / 0.1666
	return math.Pow(x, 2)
}

/*

level = const * sqrt(xp)

level / const * sqrt(xp)

pow(level /const) = xp




*/
