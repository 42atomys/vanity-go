package server

import (
	"os"
	"reflect"
	"testing"

	"atomys.codes/go-proxy/pkg/repository"
	"github.com/spf13/viper"
)

func init() {
	if err := os.Chdir("../../.."); err != nil {
		panic(err)
	}
	viper.SetConfigFile("tests/config.test.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func Test_loadV1Config(t *testing.T) {
	type args struct {
		cfg *Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "load config from version v1",
			args: args{
				cfg: &Config{
					ApiVersion: 1,
					Proxies: []*ConfigProxy{
						{
							Namespace: "test",
							Entries: map[string]string{
								"test": "test.git",
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "load config with error on protocol",
			args: args{
				cfg: &Config{
					ApiVersion: 1,
					Proxies: []*ConfigProxy{
						{
							Namespace: "test",
							Entries: map[string]string{
								"test": "test.invalid",
							},
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := loadV1Config(tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("loadV1Config() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetConfig(t *testing.T) {
	loadConfig() //nolint

	tests := []struct {
		name string
		want *Config
	}{
		{"get config", config},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetConfig(); !reflect.DeepEqual(*got, *tt.want) {
				t.Errorf("GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoriesForNamespace(t *testing.T) {

	type args struct {
		namespace string
	}
	tests := []struct {
		name string
		args args
		want []*repository.Repository
	}{
		{
			name: "get repositories from invalid namespace",
			args: args{
				namespace: "invalid",
			},
			want: nil,
		},
		{
			name: "get repositories from valid namespace",
			args: args{
				namespace: "atomys.lab",
			},
			want: GetConfig().Proxies[0].Repositories,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RepositoriesForNamespace(tt.args.namespace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoriesForNamespace() = %v, want %v", got, tt.want)
			}
		})
	}
}
