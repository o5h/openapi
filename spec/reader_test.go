package spec_test

import (
	"testing"

	"github.com/o5h/openapi/internal/assert"
	"github.com/o5h/openapi/spec"
)

func TestPetStore(t *testing.T) {
	spec, err := spec.Load("../testdata/examples/v3.0/petstore.yaml")

	assert.Nil(t, err)
	assert.NotNil(t, spec)
	assert.Eq(t, spec.OpenAPI, "3.0.0")
	assert.Eq(t, spec.Info.Title, "Swagger Petstore")
	assert.Eq(t, spec.Info.Version, "1.0.0")
	assert.Eq(t, spec.Info.License.Name, "MIT")

	pets := spec.Paths["/pets"]
	assert.Eq(t, pets["get"].OperationId, "listPets")
	assert.Eq(t, spec.Components.Schemas[0].Name, "Pet")
	assert.Eq(t, spec.Components.Schemas[1].Name, "Pets")
	t.Log(spec)

}

func TestUspto(t *testing.T) {
	spec, err := spec.Load("../testdata/examples/v3.0/uspto.yaml")

	assert.Nil(t, err)
	assert.NotNil(t, spec)
	assert.Eq(t, spec.OpenAPI, "3.0.1")
	assert.Eq(t, spec.Info.Title, "USPTO Data Set API")
	assert.Eq(t, spec.Info.Version, "1.0.0")
	assert.Eq(t, spec.Info.License.Name, "")

	root, ok := spec.Paths["/"]
	assert.True(t, ok)
	assert.NotNil(t, root)

	assert.Eq(t, root["get"].OperationId, "list-data-sets")
	t.Log(spec)
}
