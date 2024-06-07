package spec

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func Read(r io.Reader) (*OpenAPI, error) {
	decoder := yaml.NewDecoder(r)
	var spec OpenAPI
	err := decoder.Decode(&spec)
	return &spec, err
}

func Load(file string) (*OpenAPI, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Read(f)
}
