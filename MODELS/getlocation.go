package MODELS

import (
	"math"
)

func GetLocation(a, b, c, d, e, f, g, h, i float64) [2]float64 {

	P := ((math.Pow(c, 2) - math.Pow(a, 2) - math.Pow(b, 2)) - (math.Pow(f, 2) - math.Pow(d, 2) - math.Pow(e, 2)))
	Q := ((math.Pow(c, 2) - math.Pow(a, 2) - math.Pow(b, 2)) - (math.Pow(i, 2) - math.Pow(g, 2) - math.Pow(h, 2)))

	X := ((P * (-b + h)) - (Q * (-b + e))) / ((2 * (-a + d) * (-b + h)) - (2 * (-a + g) * (-b + e)))

	Y := (P - ((2 * X) * (-a + d))) / (2 * (-b + e))

	var res [2]float64

	res[0] = X
	res[1] = Y
	return res
}
