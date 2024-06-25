package config

import (
	"os"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

var File string

func New() (*Config, error) {
	cfg := &Config{}

	c, err := os.ReadFile(File)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(c, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(c)

	type plain Config
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}

	return nil
}
