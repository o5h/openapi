package openapi_test

import (
	"testing"

	"github.com/o5h/openapi"
	"github.com/o5h/openapi/internal/assert"
)

func TestRead(t *testing.T) {
	spec, err := openapi.Load("testdata/examples/v3.0/petstore.yaml")

	assert.Nil(t, err)
	assert.NotNil(t, spec)
	assert.Eq(t, spec.OpenAPI, "3.0.0")
	assert.Eq(t, spec.Info.Title, "Swagger Petstore")
	assert.Eq(t, spec.Info.Version, "1.0.0")
	assert.Eq(t, spec.Info.License.Name, "MIT")

	pets := spec.Paths["/pets"]
	assert.Eq(t, pets.Get.OperationId, "listPets")
	t.Log(spec)

}
