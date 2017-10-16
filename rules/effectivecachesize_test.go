package rules

import (
	"fmt"
	"testing"
)

var bufferCases = []struct {
	osFamily  string
	pgVersion float32
	env       EnvironmentName
	totalRAM  int
	expected  int
}{
	{"Windows", 9.5, WebEnvironment, 32 * GIGABYTE, 512 * MEGABYTE}, // until 10, 512mb max for windows
	{"Windows", 10, OLTPEnvironment, 32 * GIGABYTE, 8 * GIGABYTE},
	{"Linux", 10, DesktopEnvironment, 1536 * MEGABYTE, 128 * MEGABYTE}, // the value must be equals to the default value
	{"Linux", 9.2, DesktopEnvironment, 512 * MEGABYTE, 32 * MEGABYTE},
	{"Linux", 8.4, WebEnvironment, 512 * GIGABYTE, -1}, // tests for versions older than 9.0
}

func TestEffectiveCacheSize(t *testing.T) {

	cases := []struct {
		osFamily  string
		pgVersion float32
		env       EnvironmentName
		totalRAM  int
		expected  int
	}{
		{"Windows", 9.5, WebEnvironment, 8 * GIGABYTE, 6 * GIGABYTE},
		{"Linux", 9.2, DesktopEnvironment, 1536 * MEGABYTE, 384 * MEGABYTE},
		{"Linux", 8.0, WebEnvironment, 512 * GIGABYTE, -1},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%s vs %.1f", tc.osFamily, tc.pgVersion), func(t *testing.T) {

			result, _, err := EffectiveCacheSize(tc.pgVersion, tc.osFamily, tc.env, tc.totalRAM)

			if err != nil {
				t.Error(err)
			} else if result != tc.expected {
				t.Fatalf("Expected %v, but got %v", tc.expected, result)
			}

		})
	}
}
