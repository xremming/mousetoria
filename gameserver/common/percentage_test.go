package common_test

import (
	"gameserver/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreferenceAdd(t *testing.T) {
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
		assert.InEpsilon(t, tc.expected, float64(p1.Add(p2)), 10e-16)
	}
}
