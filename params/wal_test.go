package params

import (
	"reflect"
	"testing"
)

func Test_computeMinWalSize(t *testing.T) {

	tests := []struct {
		name    string
		p       *Parameter
		args    *Input
		wantOut interface{}
		wantErr bool
	}{
		{p: &MinWalSize, name: "invalid version", args: &Input{PGVersion: 9.2}, wantErr: true},
		{p: &MinWalSize, name: "valid web", args: &Input{Env: WebEnvironment}, wantOut: 512 * MegaByte},
		{p: &MinWalSize, name: "valid mix", args: &Input{Env: MixedEnvironment}, wantOut: 512 * MegaByte},
		{p: &MinWalSize, name: "valid desktop", args: &Input{Env: DesktopEnvironment}, wantOut: 256 * MegaByte},
		{p: &MinWalSize, name: "valid oltp", args: &Input{Env: OLTPEnvironment}, wantOut: 1 * GigaByte},
		{p: &MinWalSize, name: "valid dw", args: &Input{Env: OLTPEnvironment}, wantOut: 2 * GigaByte},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := computeMinWalSize(tt.p, tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("computeMinWalSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("computeMinWalSize() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
