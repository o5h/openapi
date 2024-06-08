package generator

import (
	"testing"

	"github.com/o5h/openapi/internal/assert"
)

func TestPets(t *testing.T) {
	cfg := &Config{
		OpenAPIFile: "../testdata/examples/v3.0/petstore.yaml",
		TypeMap:     DefaultTypeMap,
	}
	g := generator{cfg: cfg}
	err := g.generate()
	assert.Nil(t, err)
	assert.Eq(t, g.api.Package, "pets")
}
