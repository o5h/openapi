package generator

import "github.com/o5h/openapi/spec"

type API struct {
	Package  string
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
