package yaml_modifier

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/dorimon-1/vmaster/pkg/yaml_reader"
)

type YamlObject map[any]any

func ParseYAML(path string) ([]YamlObject, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening config file %s\n", err.Error())
	}
	defer f.Close()

	var objects []YamlObject
	reader := yaml_reader.NewYAMLReader(bufio.NewReader(f))

	for {
		doc, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("reading config file: %s", err)
		}
		obj := make(YamlObject)
		if err := yaml.Unmarshal(doc, &obj); err != nil {
			return nil, fmt.Errorf("unmarshaling config file: %s", err)
		}

		objects = append(objects, obj)
	}
	return objects, nil
}

func UpdateYAML(path string, objects []YamlObject) error {
	// Open the file for writing, create if not exists, and truncate if it does
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("opening config file: %s", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	for i, obj := range objects {
		data, err := yaml.Marshal(obj)
		if err != nil {
			return fmt.Errorf("marshaling config: %s", err)
		}

		fmt.Println(string(data)) // Debug print, consider removing for production
		if _, err := w.Write(data); err != nil {
			return fmt.Errorf("writing config: %s", err)
		}

		if i < len(objects)-1 {
			if _, err := w.WriteString("---\n"); err != nil {
				return fmt.Errorf("writing separator: %s", err)
			}
		}
	}

	// Flush at the end, checking for errors
	if err := w.Flush(); err != nil {
		return fmt.Errorf("flushing buffer: %s", err)
	}

	return nil
}
