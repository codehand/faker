package faker

import (
	"errors"
	"math/rand"
)

// Latitude ...
func Latitude() float64 { return (rand.Float64() * 180) - 90 }

// LatitudeInRange ...
func LatitudeInRange(min, max float64) (float64, error) {
	if min > max || min < -90 || min > 90 || max < -90 || max > 90 {
		return 0, errors.New("input range is invalid")
	}
	return randFloat64Range(min, max), nil
}

// Longitude ...
func Longitude() float64 { return (rand.Float64() * 360) - 180 }

// LongitudeInRange ...
func LongitudeInRange(min, max float64) (float64, error) {
	if min > max || min < -180 || min > 180 || max < -180 || max > 180 {
		return 0, errors.New("input range is invalid")
	}
	return randFloat64Range(min, max), nil
}
