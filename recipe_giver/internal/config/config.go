package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Http struct {
	Port int `yaml:"port"`
}

type Config struct {
	Http Http `yaml:"http"`
}

func NewConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func ValidateConfigPath(filePath string) error {
	s, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", filePath)
	}
	return nil
}
