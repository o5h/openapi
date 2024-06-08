package generator

import "github.com/o5h/openapi/spec"

type Type string

const (
	ObjectType = Type("object")
	ArrayType  = Type("array")
)

type Field struct {
	Name     string
	Type     string
	Required bool
}

type TypeDef struct {
	Type      Type
	Name      string
	ItemsType *TypeDef
	Fields    []Field
}
type API struct {
	Package   string
	APIName   string
	Endpoint  []*Endpoint
	TypesDefs []TypeDef
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
