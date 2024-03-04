package yaml_modifier

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/dorimon-1/vmaster/service"
)

func ParseYAML(path string) (service.Microservices, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening config file %s\n", err.Error())
	}
	defer f.Close()

	microservices := make(service.Microservices)

	decoder := yaml.NewDecoder(f)

	if err := decoder.Decode(&microservices); err != nil {
		return nil, fmt.Errorf("decoding config file %s\n", err.Error())
	}

	return microservices, nil
}

func UpdateYAML(path string, objects service.Microservices) error {
	// Open the file for writing, create if not exists, and truncate if it does
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("opening config file: %s", err)
	}
	defer f.Close()

	encoder := yaml.NewEncoder(f)

	if err := encoder.Encode(objects); err != nil {
		return fmt.Errorf("encoding config file: %s", err)
	}
	return nil
}
