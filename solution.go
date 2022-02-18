package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type sideType int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const SidesTriangle = 3
const SidesSquare = 4
const SidesCircle = 0

func CalcSquare(sideLen float64, sidesNum sideType) float64 {
	if sidesNum == SidesTriangle {
		return math.Sqrt(3) * sideLen / 4
	} else if sidesNum == SidesSquare {
		return sideLen * sideLen
	} else if sidesNum == SidesCircle {
		return math.Pi * sideLen * sideLen
	} else {
		return 0
	}
}
