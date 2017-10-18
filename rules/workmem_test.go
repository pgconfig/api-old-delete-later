package rules

import (
	"fmt"
	"testing"
)

func TestWorkMem(t *testing.T) {

	cases := []struct {
		pgVersion float32
		env       EnvironmentName
		totalRAM  int
		maxConn   int
		expected  int
	}{
		{9.5, WebEnvironment, 32 * GIGABYTE, 100, 32 * GIGABYTE / 100},
		{10.0, DWEnvironment, 128 * GIGABYTE, 100, 128 * GIGABYTE / 100 / 2},
		{9.2, DesktopEnvironment, 512 * MEGABYTE, 100, 1 * MEGABYTE}, // checks the default
		{9.4, DesktopEnvironment, 256 * MEGABYTE, 100, 4 * MEGABYTE}, // checks the default
		{9.2, DesktopEnvironment, 2 * GIGABYTE, 100, 2 * GIGABYTE / 100 / 6},
		{8.0, WebEnvironment, 512 * GIGABYTE, 100, -1},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%s: %.1f with %d connections", tc.env, tc.pgVersion, tc.maxConn), func(t *testing.T) {

			result, _, err := WorkMem(ParameterArgs{PGVersion: tc.pgVersion, Env: tc.env, TotalRAM: tc.totalRAM, MaxConn: tc.maxConn})

			if err != nil {
				t.Error(err)
			} else if result != tc.expected {
				t.Fatalf("Expected %v, but got %v", tc.expected, result)
			}

		})
	}

}

func TestMaintenanceWorkMem(t *testing.T) {

	cases := []struct {
		pgVersion float32
		env       EnvironmentName
		totalRAM  int
		expected  int
	}{
		{9.5, WebEnvironment, 32 * GIGABYTE, 32 * GIGABYTE / 16},
		{10.0, DWEnvironment, 2 * GIGABYTE, 2 * GIGABYTE / 8},
		{10.0, OLTPEnvironment, 128 * GIGABYTE, 2 * GIGABYTE},    //checks max value
		{9.2, DesktopEnvironment, 256 * MEGABYTE, 16 * MEGABYTE}, // checks the default
		{9.4, MixedEnvironment, 768 * MEGABYTE, 64 * MEGABYTE},   // checks the default
		{8.0, DWEnvironment, 512 * GIGABYTE, -1},                 // checking for unsupported versions
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%s:%.1f", tc.env, tc.pgVersion), func(t *testing.T) {

			result, _, err := MaintenanceWorkMem(ParameterArgs{PGVersion: tc.pgVersion, Env: tc.env, TotalRAM: tc.totalRAM})

			if err != nil {
				t.Error(err)
			} else if result != tc.expected {
				t.Fatalf("Expected %v, but got %v", tc.expected, result)
			}

		})
	}

}
