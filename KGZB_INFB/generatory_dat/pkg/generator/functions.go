package generator

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/exp/constraints"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Number is a defined constraint for a generic type
type Number interface {
	constraints.Integer | constraints.Float
}

// RandomNormal generates a random number in the given range [min, max] using normal distribution
// with the given mean and standard deviation.
func RandomNormal[T Number](mean T, stdev float64, min, max T) T {
	x := rand.NormFloat64()*stdev + float64(mean)
	if x < float64(min) || x > float64(max) {
		return RandomNormal(mean, stdev, min, max)
	}
	return T(x)
}

// RandomCentered generates a random number in the given range [min, max] by averaging [centeringFactor] times random numbers
// and scaling the result to the range [min, max]. The average is biased towards the middle of the range, so the
// values close to the middle are more likely to occur with higher [centeringFactor]. [centeringFactor] < 1 returns unbiased random number.
func RandomCentered[T Number](min, max T, centeringFactor int) T {
	if centeringFactor <= 0 {
		return RandomRange(min, max)
	}

	rndSum := rand.Float64()
	for i := 0; i < centeringFactor; i++ {
		rndSum += rand.Float64()
	}

	rnd := float64(min) + (rndSum/float64(centeringFactor+1))*(float64(max)-float64(min))
	return T(rnd)
}

// RandomRange generates a random number in the given range [min, max].
func RandomRange[T Number](min, max T) T {
	x := float64(min) + rand.Float64()*(float64(max)-float64(min))
	return T(x)
}

// RandomChoice returns a random variable from the given slice of variables. If the slice is empty, an error is returned.
func RandomChoice[T any](a []T) (T, error) {
	if len(a) == 0 {
		var ret T
		return ret, fmt.Errorf("RandomChoice: The soucre slice is empty")
	}
	return a[rand.Intn(len(a))], nil
}

// MustRandomChoice returns a random variable from the given slice of variables. If the slice is empty, it returns an uninitialized value of the given type.
func MustRandomChoice[T any](a []T) T {
	if len(a) == 0 {
		var ret T
		return ret
	}
	return a[rand.Intn(len(a))]
}
