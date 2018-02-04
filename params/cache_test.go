package params

import (
	"reflect"
	"testing"
)

func Test_computeEffectiveCacheSize(t *testing.T) {

	tests := []struct {
		name    string
		p       *Parameter
		args    *Input
		wantOut interface{}
		wantErr bool
	}{
		{name: "Valid Desktop", p: &EffectiveCacheSize, args: &Input{Env: DesktopEnvironment, TotalRAM: 4 * GigaByte}, wantOut: 1 * GigaByte},
		{name: "Valid Web", p: &EffectiveCacheSize, args: &Input{Env: WebEnvironment, TotalRAM: 12 * GigaByte}, wantOut: 9 * GigaByte},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := computeEffectiveCacheSize(tt.p, tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("computeEffectiveCacheSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("computeEffectiveCacheSize() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
