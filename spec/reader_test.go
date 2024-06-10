package spec_test

import (
	"testing"

	"github.com/o5h/openapi/internal/assert"
	"github.com/o5h/openapi/spec"
)

func TestPetStore(t *testing.T) {
	openapi, err := spec.Load("../testdata/examples/v3.0/petstore.yaml")

	assert.Nil(t, err)
	assert.NotNil(t, openapi)
	assert.Eq(t, openapi.OpenAPI, "3.0.0")
	assert.Eq(t, openapi.Info.Title, "Swagger Petstore")
	assert.Eq(t, openapi.Info.Version, "1.0.0")
	assert.Eq(t, openapi.Info.License.Name, "MIT")

	path0 := openapi.Paths[0]
	assert.Eq(t, path0.Name, "/pets")
	assert.Eq(t, path0.Value[0].Name, spec.MethodGet)
	assert.Eq(t, path0.Value[0].Value.OperationId, "listPets")
	assert.Eq(t, openapi.Components.Schemas[0].Name, "Pet")
	assert.Eq(t, openapi.Components.Schemas[1].Name, "Pets")
	t.Log(openapi)

}

func TestUspto(t *testing.T) {
	openapi, err := spec.Load("../testdata/examples/v3.0/uspto.yaml")

	assert.Nil(t, err)
	assert.NotNil(t, openapi)
	assert.Eq(t, openapi.OpenAPI, "3.0.1")
	assert.Eq(t, openapi.Info.Title, "USPTO Data Set API")
	assert.Eq(t, openapi.Info.Version, "1.0.0")
	assert.Eq(t, openapi.Info.License.Name, "")

	path0 := openapi.Paths[0]
	assert.Eq(t, path0.Name, "/")

	assert.Eq(t, path0.Value[0].Value.OperationId, "list-data-sets")
	t.Log(openapi)
}
