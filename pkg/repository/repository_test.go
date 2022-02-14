package repository

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		entrypoint  string
		destination string
	}
	tests := []struct {
		name    string
		args    args
		want    *Repository
		wantErr bool
	}{
		{
			name: "missing entrypoint",
			args: args{
				entrypoint:  "",
				destination: "github.com/user/repo",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "missing destination",
			args: args{
				entrypoint:  "github.com/user/repo",
				destination: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "missing protocol",
			args: args{
				entrypoint:  "github.com/user/repo",
				destination: "github.com/user/repo",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "git repository",
			args: args{
				entrypoint:  "example.org/repo",
				destination: "github.com/user/repo.git",
			},
			want: &Repository{
				EntryPoint:  "example.org/repo",
				Destination: "github.com/user/repo.git",
				Protocol:    ProtocolGit,
			},
			wantErr: false,
		},
		{
			name: "mercurial repository",
			args: args{
				entrypoint:  "example.org/project",
				destination: "bitbucket.org/user/project.hg",
			},
			want: &Repository{
				EntryPoint:  "example.org/project",
				Destination: "bitbucket.org/user/project.hg",
				Protocol:    ProtocolHg,
			},
			wantErr: false,
		},
		{
			name: "bazaar repository",
			args: args{
				entrypoint:  "example.org/project",
				destination: "launchpad.net/project.bzr",
			},
			want: &Repository{
				EntryPoint:  "example.org/project",
				Destination: "launchpad.net/project.bzr",
				Protocol:    ProtocolBzr,
			},
			wantErr: false,
		},
		{
			name: "fossil repository",
			args: args{
				entrypoint:  "example.org/project",
				destination: "example.org/user/project.fossil",
			},
			want: &Repository{
				EntryPoint:  "example.org/project",
				Destination: "example.org/user/project.fossil",
				Protocol:    ProtocolFossil,
			},
			wantErr: false,
		},
		{
			name: "subversion repository",
			args: args{
				entrypoint:  "example.org/project",
				destination: "example.org/user/project.svn",
			},
			want: &Repository{
				EntryPoint:  "example.org/project",
				Destination: "example.org/user/project.svn",
				Protocol:    ProtocolSvn,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.entrypoint, tt.args.destination)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
