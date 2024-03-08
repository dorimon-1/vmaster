package yaml_modifier

import (
	"testing"
)

func TestParseWrite(t *testing.T) {
	path := "../../example/5.0.1/values.yaml"
	objects, err := ParseYAML(path)
	if err != nil {
		t.Errorf("Error parsing yaml: %s", err)
		return
	}

	objects["microserviceA"].Image.Tag = "5.0.2"
	UpdateYAML(path, objects)

	_, err = ParseYAML(path)
	if err != nil {
		t.Errorf("Error parsing yaml: %s", err)
	}
}
