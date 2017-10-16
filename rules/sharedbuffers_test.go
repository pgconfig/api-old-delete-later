package rules

import (
	"fmt"
	"testing"
)

func TestSharedBuffers(t *testing.T) {

	var bufferCases = []struct {
		osFamily  OSFamily
		pgVersion float32
		env       EnvironmentName
		totalRAM  int
		expected  int
	}{
		{WindowsOS, 9.5, WebEnvironment, 32 * GIGABYTE, 512 * MEGABYTE}, // until 10, 512mb max for windows
		{WindowsOS, 10, OLTPEnvironment, 32 * GIGABYTE, 8 * GIGABYTE},
		{LinuxOS, 10, DesktopEnvironment, 1536 * MEGABYTE, 128 * MEGABYTE}, // the value must be equals to the default value
		{LinuxOS, 9.2, DesktopEnvironment, 512 * MEGABYTE, 32 * MEGABYTE},
		{UnixOS, 8.4, WebEnvironment, 512 * GIGABYTE, -1}, // tests for versions older than 9.0
	}

	for _, tc := range bufferCases {
		t.Run(fmt.Sprintf("%s vs %.1f", tc.osFamily, tc.pgVersion), func(t *testing.T) {

			result, _, err := SharedBuffers(tc.pgVersion, tc.osFamily, tc.env, tc.totalRAM)

			if err != nil {
				t.Error(err)
			}
			if result != tc.expected {
				t.Fatalf("Expected %v, but got %v", tc.expected, result)
			}

		})
	}
}
