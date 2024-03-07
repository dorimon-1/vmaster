package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Environments map[string]Environment `yaml:"environments"`
}

type Environment struct {
	Name          string `yaml:"name"`
	VersionPrefix string `yaml:"version_prefix"`
	FilePath      string `yaml:"file_path"`
}

func NewConfig() *Config {
	return &Config{
		Environments: make(map[string]Environment),
	}
}

func LoadConfig(path string) (*Config, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("getting absolute path: %v", err)
	}

	f, err := os.Open(absPath)
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

	cfg.Environments["Development"] = Environment{
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
