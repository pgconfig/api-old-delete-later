package rules

import (
	"fmt"
	"testing"
)

func TestMaxWalSize(t *testing.T) {

	cases := []struct {
		pgVersion float32
		env       EnvironmentName
		expected  int
	}{
		{8.2, WebEnvironment, -1},
		{10.0, MixedEnvironment, 2 * GIGABYTE},
		{9.6, WebEnvironment, 2 * GIGABYTE},
		{9.5, OLTPEnvironment, 3 * GIGABYTE},
		{9.5, DesktopEnvironment, 1 * GIGABYTE},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%s: %.1f", tc.env, tc.pgVersion), func(t *testing.T) {

			result, _, err := MaxWalSize(ParameterArgs{PGVersion: tc.pgVersion, Env: tc.env})

			if err != nil {
				t.Error(err)
			} else if result != tc.expected {
				t.Fatalf("Expected %v, but got %v", tc.expected, result)
			}

		})
	}

}
