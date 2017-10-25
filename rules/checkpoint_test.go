package rules

import (
	"fmt"
	"testing"
)

type testeaaaa struct {
	pgVersion float32
	env       EnvironmentName
	expected  int
}

func TestCheckpointSegments(t *testing.T) {

	cases := []testeaaaa{
		{9.4, WebEnvironment, 32},
		{9.0, MixedEnvironment, 32},
		{9.2, OLTPEnvironment, 96},
		{9.1, DWEnvironment, 256},
		{8.4, DWEnvironment, -1},       // checking unsupported versions
		{10.0, DesktopEnvironment, -1}, // unsupported too
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%s:%.1f", tc.env, tc.pgVersion), func(t *testing.T) {

			result, _, err := CheckpointSegments(ParameterArgs{Env: tc.env, PGVersion: tc.pgVersion})

			if err != nil {
				t.Error(err)
			} else if result != tc.expected {
				t.Fatalf("Expected %v, but got %v", tc.expected, result)
			}

		})
	}
}
