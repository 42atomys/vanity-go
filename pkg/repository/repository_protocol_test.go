package repository

import (
	"testing"
)

func TestProtocol_Validate(t *testing.T) {
	tests := []struct {
		name    string
		rp      Protocol
		wantErr bool
	}{
		{"bzr", ProtocolBzr, false},
		{"fossil", ProtocolFossil, false},
		{"git", ProtocolGit, false},
		{"hg", ProtocolHg, false},
		{"svn", ProtocolSvn, false},
		{"invalid", Protocol("invalid"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.rp.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Protocol.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProtocol_String(t *testing.T) {
	tests := []struct {
		name string
		rp   Protocol
		want string
	}{
		{"bzr", ProtocolBzr, "bzr"},
		{"fossil", ProtocolFossil, "fossil"},
		{"git", ProtocolGit, "git"},
		{"hg", ProtocolHg, "hg"},
		{"svn", ProtocolSvn, "svn"},
		{"invalid", Protocol("invalid"), ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rp.String(); got != tt.want {
				t.Errorf("Protocol.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
