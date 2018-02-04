package params

import (
	"reflect"
	"testing"
)

func Test_computeCheckPointSegments(t *testing.T) {

	tests := []struct {
		name    string
		p       *Parameter
		args    *Input
		wantOut interface{}
		wantErr bool
	}{
		{p: &CheckPointSegments, name: "Valid Web", args: &Input{PGVersion: 9.2, Env: WebEnvironment}, wantOut: 32},
		{p: &CheckPointSegments, name: "Valid Mix", args: &Input{PGVersion: 9.0, Env: MixedEnvironment}, wantOut: 32},
		{p: &CheckPointSegments, name: "Valid oltp", args: &Input{PGVersion: 9.0, Env: OLTPEnvironment}, wantOut: 96},
		{p: &CheckPointSegments, name: "Valid dw", args: &Input{PGVersion: 9.0, Env: DWEnvironment}, wantOut: 256},
		{p: &CheckPointSegments, name: "Valid dev machine", args: &Input{PGVersion: 9.0, Env: DesktopEnvironment}, wantOut: 16},
		{p: &CheckPointSegments, name: "no env", args: &Input{PGVersion: 9.2}, wantOut: 16},
		{p: &CheckPointSegments, name: "Invalid version greater", args: &Input{PGVersion: 10.0}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := computeCheckPointSegments(tt.p, tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("computeCheckPointSegments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("computeCheckPointSegments() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
