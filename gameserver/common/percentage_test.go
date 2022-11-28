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
		p1 := common.Percentage(tc.a)
		p2 := common.Percentage(tc.b)
		res := float64(p1.Add(p2))
		assert.InEpsilonf(t, tc.expected, res, 10e-16, "expected %f + %f = %f, got %f", tc.a, tc.b, tc.expected, res)
	}
}

func TestPercentageSub(t *testing.T) {
	var testCases = []struct {
		a        float64
		b        float64
		expected float64
	}{
		{0.19, 0.1, 0.1},
		{0.75, 0.5, 0.5},
		{0.99, 0.9, 0.9},
	}

	for _, tc := range testCases {
		p1 := common.Percentage(tc.a)
		p2 := common.Percentage(tc.b)
		res := float64(p1.Sub(p2))
		assert.InEpsilonf(t, tc.expected, res, 10e-16, "expected %f - %f = %f, got %f", tc.a, tc.b, tc.expected, res)
	}
}
