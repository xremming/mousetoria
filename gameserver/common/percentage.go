package common

import (
	"math"
)

func MultAdd(a, b float64) float64 {
	return 1 - (1-a)*(1-b)
}

func AsPercentage(v float64) float64 {
	if math.IsInf(v, -1) {
		return 0
	}
	if math.IsInf(v, 1) {
		return 1
	}

	if math.IsNaN(v) {
		return 0
	}

	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}

	return v
}

type RangePair struct {
	Zero float64
	Min  float64
	Max  float64
}

func NewRangePair(zero, min, max float64) RangePair {
	if min > zero || zero > max {
		panic("min <= zero <= max must be true for RangePair")
	}
	return RangePair{zero, min, max}
}

func (r RangePair) WithModifiers(staticModifiers []float64, positiveModifiers []float64, negativeModifiers []float64) float64 {
	newZero := r.Zero

	for _, modifier := range staticModifiers {
		if math.IsNaN(modifier) {
			continue
		}

		if math.IsInf(modifier, -1) {
			return r.Min
		}
		if math.IsInf(modifier, 1) {
			return r.Max
		}

		newZero += modifier
	}

	if newZero < r.Min {
		return r.Min
	}
	if newZero > r.Max {
		return r.Max
	}

	var (
		buff float64
		nerf float64
	)

	for _, modifier := range positiveModifiers {
		normalizedModifier := AsPercentage(modifier)
		buff = (1 - (1-buff)*(1-normalizedModifier))
	}

	for _, modifier := range negativeModifiers {
		normalizedModifier := AsPercentage(modifier)
		nerf = (1 - (1-nerf)*(1-normalizedModifier))
	}

	totalBuff := buff - nerf
	if totalBuff >= 0 {
		newZero = newZero + (r.Max-newZero)*totalBuff
	} else {
		newZero = newZero - (newZero-r.Min)*totalBuff
	}

	if newZero < r.Min {
		return r.Min
	}
	if newZero > r.Max {
		return r.Max
	}

	return newZero
}
