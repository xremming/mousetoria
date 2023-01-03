package common

import "math"

type Modifiers struct {
	StaticPositive     []float64
	StaticNegative     []float64
	PercentagePositive []float64
	PercentageNegative []float64
}

type Gauge interface {
	ApplyModifiers(Modifiers) float64
}

// GaugeIncreasing represents a gauge where a larger values is "better".
type GaugeIncreasing struct {
	Default float64
	Min     float64
	Max     float64
}

func (g GaugeIncreasing) ApplyModifiers(m Modifiers) float64 {
	value := g.Default

	for _, modifier := range m.StaticPositive {
		value += modifier
	}

	for _, modifier := range m.StaticNegative {
		value -= modifier
	}

	var (
		percentagePositiveTotal float64
		percentageNegativeTotal float64
	)

	for _, percentagePositive := range m.PercentagePositive {
		percentagePositiveTotal = (1 - (1-percentagePositiveTotal)*(1-percentagePositive))
	}

	for _, percentageNegative := range m.PercentageNegative {
		percentageNegativeTotal = (1 - (1-percentageNegativeTotal)*(1-percentageNegative))
	}

	minDiff := math.Abs(value - g.Min)
	maxDiff := math.Abs(value - g.Max)

	percentageDifference := percentagePositiveTotal - percentageNegativeTotal
	if percentageDifference > 0 {
		value += maxDiff * percentageDifference
	}
	if percentageDifference < 0 {
		value -= minDiff * -percentageDifference
	}

	return value
}

// GaugeDecreasing represents a gauge where a smaller values is "better".
type GaugeDecreasing struct {
	Default float64
	Min     *float64
	Max     *float64
}
