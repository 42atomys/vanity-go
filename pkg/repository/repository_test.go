package repository

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		entrypoint  string
		destination string
		namespace   string
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
			name: "missing namespace",
			args: args{
				entrypoint:  "github.com/user/repo",
				destination: "github.com/user/repo",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "missing protocol",
			args: args{
				entrypoint:  "github.com/user/repo",
				destination: "github.com/user/repo",
				namespace:   "example.org",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "git repository",
			args: args{
				entrypoint:  "repo",
				destination: "github.com/user/repo.git",
				namespace:   "example.org",
			},
			want: &Repository{
				EntryPoint:  "repo",
				Destination: "github.com/user/repo.git",
				Namespace:   "example.org",
				Protocol:    ProtocolGit,
			},
			wantErr: false,
		},
		{
			name: "mercurial repository",
			args: args{
				entrypoint:  "project",
				destination: "bitbucket.org/user/project.hg",
				namespace:   "example.org",
			},
			want: &Repository{
				EntryPoint:  "project",
				Destination: "bitbucket.org/user/project.hg",
				Namespace:   "example.org",
				Protocol:    ProtocolHg,
			},
			wantErr: false,
		},
		{
			name: "bazaar repository",
			args: args{
				entrypoint:  "project",
				destination: "launchpad.net/project.bzr",
				namespace:   "example.org",
			},
			want: &Repository{
				EntryPoint:  "project",
				Destination: "launchpad.net/project.bzr",
				Namespace:   "example.org",
				Protocol:    ProtocolBzr,
			},
			wantErr: false,
		},
		{
			name: "fossil repository",
			args: args{
				entrypoint:  "project",
				destination: "example.org/user/project.fossil",
				namespace:   "example.org",
			},
			want: &Repository{
				EntryPoint:  "project",
				Destination: "example.org/user/project.fossil",
				Namespace:   "example.org",
				Protocol:    ProtocolFossil,
			},
			wantErr: false,
		},
		{
			name: "subversion repository",
			args: args{
				entrypoint:  "project",
				destination: "example.org/user/project.svn",
				namespace:   "example.org",
			},
			want: &Repository{
				EntryPoint:  "project",
				Destination: "example.org/user/project.svn",
				Namespace:   "example.org",
				Protocol:    ProtocolSvn,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.entrypoint, tt.args.destination, tt.args.namespace)
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

func TestRepositoryURL(t *testing.T) {
	tests := []struct {
		name  string
		input *Repository
		want  string
	}{
		{
			name: "sub folder",
			input: &Repository{
				Namespace:   "example.org",
				Destination: "github.com/user/repo.git",
				EntryPoint:  "sub/project",
			},
			want: "example.org/sub/project",
		},
		{
			name: "sub folder with slash",
			input: &Repository{
				Namespace:   "example.org",
				Destination: "github.com/user/repo.git",
				EntryPoint:  "/sub/project",
			},
			want: "example.org/sub/project",
		},
		{
			name: "with ending slash",
			input: &Repository{
				Namespace:   "example.org",
				Destination: "github.com/user/repo.git",
				EntryPoint:  "project/",
			},
			want: "example.org/project",
		},
		{
			name: "root case with -",
			input: &Repository{
				Namespace:   "example.org",
				Destination: "github.com/user/repo.git",
				EntryPoint:  "-",
			},
			want: "example.org",
		},
		{
			name: "root case with /",
			input: &Repository{
				Namespace:   "example.org",
				Destination: "github.com/user/repo.git",
				EntryPoint:  "/",
			},
			want: "example.org",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, _ := New(tt.input.EntryPoint, tt.input.Destination, tt.input.Namespace)
			if repo.URL() != tt.want {
				t.Errorf("URL() = %v, want %v", repo.URL(), tt.want)
			}
		})
	}
}
