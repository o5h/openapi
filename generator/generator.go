package generator

import "github.com/o5h/openapi/spec"

type API struct {
	Endpoint []*Endpoint
}

type Endpoint struct {
	Method     spec.Method
	Operation  string
	Path       string
	Parameters []Parameter
}

type Parameter struct {
	Name   string
	Type   string
	Format string
}

func Generate(openapi *spec.OpenAPI) *API {
	api := API{}
	return &api
}
