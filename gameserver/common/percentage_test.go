package common_test

import (
	"gameserver/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPercentageAdd(t *testing.T) {
	var testCases = []struct {
		a        float64
		b        float64
		expected float64
	}{
		{0.1, 0.1, 0.19},
		{0.5, 0.5, 0.75},
		{0.9, 0.9, 0.99},
	}

	for _, tc := range testCases {
		p1 := common.AsPercentage(tc.a)
		p2 := common.AsPercentage(tc.b)
		res := common.MultAdd(p1, p2)
		assert.InEpsilonf(t, tc.expected, res, 10e-16, "expected %f + %f = %f, got %f", tc.a, tc.b, tc.expected, res)
	}
}

func TestRangeLogic(t *testing.T) {
	r := common.NewRangePair(100, 0, 200)

	assert.LessOrEqual(t, 121., r.WithModifiers([]float64{10}, []float64{0.1}, nil))
}
