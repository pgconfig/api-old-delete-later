package rules

import (
	"fmt"
	"testing"
)

func TestSharedBuffers(t *testing.T) {

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
