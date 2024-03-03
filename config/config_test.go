package config

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	path := "./config.yaml"
	cfg, err := LoadConfig(path)
	if err != nil {
		t.Errorf("Error loading config: %v", err)
	}

	fmt.Println(*cfg)
	if cfg.Enviroments["Development"].Name != "Development" {
		t.Errorf("Expected Development got %s", cfg.Enviroments["Development"].Name)
	}
}

func TestSaveConfig(t *testing.T) {
	path := "./config_test.yaml"
	cfg := GenerateConfig()

	err := SaveConfig(path, cfg)
	if err != nil {
		t.Errorf("Error saving config: %v", err)
	}
}
