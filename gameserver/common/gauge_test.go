package common_test

import (
	"gameserver/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGaugeNegativeValues(t *testing.T) {
	g := common.GaugeIncreasing{
		Default: 0,
		Min:     -100,
		Max:     200,
	}

	assert.Equal(t, 18., g.ApplyModifiers(common.Modifiers{
		StaticPositive:     []float64{10},
		PercentagePositive: []float64{0.1},
	}))
}
