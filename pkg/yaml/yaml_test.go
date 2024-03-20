package yaml

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

	objects[0]["microserviceA"].(YamlObject)["image"].(YamlObject)["tag"] = "5.0.1"
	UpdateYAML(path, objects)

	_, err = ParseYAML(path)
	if err != nil {
		t.Errorf("Error parsing yaml: %s", err)
	}
}
