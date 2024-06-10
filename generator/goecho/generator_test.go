package goecho_test

import (
	"testing"
	"text/template"

	"github.com/o5h/openapi/generator"
	"github.com/o5h/openapi/generator/goecho"
	"github.com/o5h/openapi/internal/assert"
)

func TestPets(t *testing.T) {
	apiTemplate := template.Must(template.New("api").Parse(goecho.DefaultTemplate))
	cfg := &generator.Config{
		OpenAPIFile: "../../testdata/examples/v3.0/petstore.yaml",
		TypeMap:     goecho.DefaultTypeMap,
		APITemplate: apiTemplate,
	}
	err := generator.Generate(cfg)
	assert.Nil(t, err)
}
