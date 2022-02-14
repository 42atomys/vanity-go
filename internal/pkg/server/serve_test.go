//go:build !skip

package server

import (
	_ "embed"
	"testing"
)

func TestServe(t *testing.T) {
	type args struct {
		port int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"invalid port", args{port: -1}, true},
		{"invalid port", args{port: 1_000_000}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Serve(tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("Serve() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
