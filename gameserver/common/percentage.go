package common

import (
	"fmt"
	"math"
)

type Percentage float64

func (p Percentage) Norm() float64 {
	if p < 0 {
		return 0
	}
	if p > 1 {
		return 1
	}

	if math.IsNaN(float64(p)) {
		return 0
	}

	return float64(p)
}

func (p Percentage) Add(other Percentage) Percentage {
	return Percentage(1 - (1-p.Norm())*(1-other.Norm()))
}

func (p Percentage) Validate() error {
	if math.IsNaN(float64(p)) || math.IsInf(float64(p), 0) {
		return fmt.Errorf("preference must not be NaN or Inf")
	}

	if p < 0 || p > 1 {
		return fmt.Errorf("preference must be between 0 and 1")
	}

	return nil
}
