package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type PhotoParserConfig struct {
	Port               int `yaml:"port"`
	GracefulTimeoutSec int `yaml:"graceful.timeout.sec"`
}

func NewPhotoParserConfig(cfgPath string) *PhotoParserConfig {
	configData, err := os.ReadFile(cfgPath)
	if err != nil {
		panic(err)
	}
	var cfg PhotoParserConfig
	err = yaml.Unmarshal(configData, &cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}
