package mouses

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCountAliveMouse(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                     string
		originalAmount           int
		lifeSpan                 int
		months                   int
		expectedAliveMouseAmount int
	}{
		{"zero_mouse_from_original", 0, 1, 1, 0},
		{"mouse_with_life_span_zero", 1, 0, 1, 0},
		{"months_less_than_life_span", 1, 3, 2, 4},
		{"months_equal_to_life_span", 1, 3, 3, 7},
		{"months_greater_than_life_span", 1, 3, 4, 13},
	}
	for _, tesCase := range testCases {
		tc := tesCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			aliveMouseAmount := CountAliveMouse(tc.originalAmount, tc.lifeSpan, tc.months)
			require.Equal(t, tc.expectedAliveMouseAmount, aliveMouseAmount)
		})
	}
}
