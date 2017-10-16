package rules

import (
	"fmt"
	"testing"
)

func TestEffectiveCacheSize(t *testing.T) {

	cases := []struct {
		pgVersion float32
		env       EnvironmentName
		totalRAM  int
		expected  int
	}{
		{9.5, WebEnvironment, 8 * GIGABYTE, 6 * GIGABYTE},
		{9.2, DesktopEnvironment, 1536 * MEGABYTE, 384 * MEGABYTE},
		{8.0, WebEnvironment, 512 * GIGABYTE, -1},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%.1f", tc.pgVersion), func(t *testing.T) {

			result, _, err := EffectiveCacheSize(tc.pgVersion, tc.env, tc.totalRAM)

			if err != nil {
				t.Error(err)
			} else if result != tc.expected {
				t.Fatalf("Expected %v, but got %v", tc.expected, result)
			}

		})
	}
}
