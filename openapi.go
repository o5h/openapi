package openapi

import "gopkg.in/yaml.v3"

type HTTPMethod string

const (
	MethodGet     = HTTPMethod("GET")
	MethodHead    = HTTPMethod("HEAD")
	MethodPost    = HTTPMethod("POST")
	MethodPut     = HTTPMethod("PUT")
	MethodPatch   = HTTPMethod("PATCH")
	MethodDelete  = HTTPMethod("DELETE")
	MethodConnect = HTTPMethod("CONNECT")
	MethodOptions = HTTPMethod("OPTIONS")
	MethodTrace   = HTTPMethod("TRACE")
)

type In string

const (
	InQuery  = In("query")
	InHeader = In("header")
	InPath   = In("path")
	InCookie = In("cookie")
)

type OpenAPI struct {
	OpenAPI    string     `yaml:"openapi"`
	Info       *Info      `yaml:"info"`
	Paths      Paths      `yaml:"paths"`
	Components Components `yaml:"components"`
}

type Components struct {
	Schemas         Schemas `yaml:"schemas"`
	SecuritySchemes Schemas `yaml:"securitySchemes"`
}

type Schemas map[string]any
type SecuritySchemes map[string]any

type Info struct {
	Title   string  `yaml:"title"`
	Version string  `yaml:"version"`
	License License `yaml:"license"`
}
type License struct {
	Name       string `yaml:"name"`
	Identifier string `yaml:"identifier,omitempty"`
	URL        string `yaml:"url,omitempty"`
}
type Paths map[string]*PathItem
type PathItem struct {
	Get  *Operation `yaml:"get,omitempty"`
	Post *Operation `yaml:"post,omitempty"`
}

type Operation struct {
	Summary     string                 `yaml:"summary,omitempty"`
	OperationId string                 `yaml:"operationId,omitempty"`
	Tags        []string               `yaml:"tags,omitempty"`
	Parameters  []Parameter            `yaml:"parameters,omitempty"`
	RequestBody RequestBodyOrReference `yaml:"requestBody,omitempty"`
	Security    Security               `yaml:"security,omitempty"`
	Responses   Responses              `yaml:"responses,omitempty"`
}

type RequestBodyOrReference struct {
	Reference   *Reference   `yaml:",inline"`
	RequestBody *RequestBody `yaml:",inline"`
}

func (x *RequestBodyOrReference) UnmarshalYAML(value *yaml.Node) error {
	var v map[string]any = make(map[string]any)
	value.Decode(v)
	if _, ok := v["$ref"]; ok {
		x.Reference = &Reference{}
		return value.Decode(x.Reference)
	}
	x.RequestBody = &RequestBody{}
	return value.Decode(x.RequestBody)
}

type RequestBody struct {
	Description string               `yaml:"description,omitempty"`
	Content     map[string]MediaType `yaml:"content"`
}

type MediaType struct{}

type Parameter struct {
	Name            string `yaml:"name"`
	In              In     `yaml:"in"`
	Description     string `yaml:"description,omitempty"`
	Required        bool   `yaml:"required"`
	Deprecated      bool   `yaml:"deprecated"`
	AllowEmptyValue bool   `yaml:"allowEmptyValue"`
	Schema          Schema `yaml:"schema,omitempty"`
}
type Security []map[string][]string
type Responses map[string]Response
type Response struct {
	Content map[string]Content `yaml:"content"`
}
type Content struct {
	Description string  `yaml:"description"`
	Schema      *Schema `yaml:"schema"`
}

type Schema struct {
	Ref *Ref `yaml:"$ref"`
}

type Reference struct {
	Ref         *Ref   `yaml:"$ref"`
	Summary     string `yaml:"summary"`
	Description string `yaml:"description"`
}

type Ref string
