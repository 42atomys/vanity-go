package server

import (
	"fmt"

	"atomys.codes/go-proxy/pkg/repository"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	ApiVersion int            `mapstructure:"apiVersion,omitempty"`
	Proxies    []*ConfigProxy `mapstructure:"proxies,omitempty"`
}

type ConfigProxy struct {
	Namespace    string
	Entries      map[string]string
	Repositories []*repository.Repository
}

var config *Config

func loadConfig() error {
	var err error
	if err = viper.Unmarshal(&config); err != nil {
		return fmt.Errorf("error unmarshalling config")
	}

	switch viper.GetInt("apiVersion") {
	case 1:
		err = loadV1Config(config)
	}

	if err != nil {
		return err
	}
	return nil
}

// GetConfig returns the configuration of the proxy
func GetConfig() *Config {
	return config
}

// RepositoriesForNamespace returns the repositories for the given namespace
// If the namespace is empty or invalid, it returns nil
func RepositoriesForNamespace(namespace string) []*repository.Repository {
	for _, proxy := range GetConfig().Proxies {
		if proxy.Namespace == namespace {
			return proxy.Repositories
		}
	}
	return nil
}

// loadV1Config loads the config from the v1 format
func loadV1Config(cfg *Config) error {
	log.Info().Msgf("Load config from version v%d", viper.GetInt("apiVersion"))
	defer log.Info().Msgf("Loaded %d proxies", len(cfg.Proxies))

	for _, proxy := range cfg.Proxies {
		for entrypoint, destination := range proxy.Entries {
			repo, err := repository.New(entrypoint, destination)
			if err != nil {
				log.Error().Err(err).Str("entrypoint", entrypoint).Str("destination", destination).Msg("error creating repository")
				return err
			}

			proxy.Repositories = append(proxy.Repositories, repo)
			log.Debug().Str("proxy", proxy.Namespace).Msgf("Loaded repository %s", entrypoint)
		}

		log.Debug().Str("proxy", proxy.Namespace).Msgf("Loaded %d repositories", len(proxy.Repositories))
	}
	return nil
}
