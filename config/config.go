package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Enviroments map[string]Enviroment `yaml:"enviroments"`
}

type Enviroment struct {
	Name          string `yaml:"name"`
	VersionPrefix string `yaml:"version_prefix"`
	FilePath      string `yaml:"file_path"`
}

func NewConfig() *Config {
	return &Config{
		Enviroments: make(map[string]Enviroment),
	}
}

func LoadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening file %s: %v", path, err)
	}
	defer f.Close()

	cfg := &Config{}
	decoder := yaml.NewDecoder(f)

	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("decoding %s: %v", path, err)
	}

	return cfg, nil
}

func GenerateConfig() *Config {
	cfg := NewConfig()

	cfg.Enviroments["Development"] = Enviroment{
		Name:          "Development",
		VersionPrefix: "v1",
		FilePath:      "./development.yaml",
	}

	return cfg
}

func SaveConfig(path string, cfg *Config) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return fmt.Errorf("opening file %s: %v", path, err)
	}
	defer f.Close()

	encoder := yaml.NewEncoder(f)

	encoder.Encode(cfg)
	return nil
}
