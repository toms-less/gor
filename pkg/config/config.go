package config

import (
	"fmt"
	"os"

	"github.com/go-yaml/yaml"
)

type ServerConfig struct {
	Port uint32
}

type Config struct {
	Server ServerConfig
}

func Init() (*Config, error) {
	cwd, wdErr := os.Getwd()
	if wdErr != nil {
		return nil, wdErr
	}
	path := fmt.Sprintf("%s%s", cwd, "/config/config.yaml")
	file, fileErr := os.Open(path)
	if fileErr != nil {
		return nil, fileErr
	}

	config := &Config{}
	yamlErr := yaml.NewDecoder(file).Decode(config)
	if yamlErr != nil {
		return nil, yamlErr
	}

	// check user configuration.
	configErr := check(config)
	if configErr != nil {
		return nil, configErr
	}

	return config, nil
}

func check(config *Config) error {
	serverConfig := config.Server
	if serverConfig.Port < 1024 || serverConfig.Port > 65536 {
		return fmt.Errorf("Server port is not invalid, should be between 1024 and 65536.")
	}
	return nil
}
