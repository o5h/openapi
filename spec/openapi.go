package spec

import "gopkg.in/yaml.v3"

type Method string
type In string
type Format string
type Type string

const (
	MethodGet     = Method("GET")
	MethodHead    = Method("HEAD")
	MethodPost    = Method("POST")
	MethodPut     = Method("PUT")
	MethodPatch   = Method("PATCH")
	MethodDelete  = Method("DELETE")
	MethodConnect = Method("CONNECT")
	MethodOptions = Method("OPTIONS")
	MethodTrace   = Method("TRACE")

	TypeInteger = Type("integer")
	TypeArray   = Type("array")
	TypeString  = Type("string")

	InQuery  = In("query")
	InHeader = In("header")
	InPath   = In("path")
	InCookie = In("cookie")

	FormatInt32 = Format("int32")
	FormatInt64 = Format("int64")
)

type OpenAPI struct {
	OpenAPI    string     `yaml:"openapi"`
	Info       *Info      `yaml:"info"`
	Servers    []Server   `yaml:"servers"`
	Paths      Paths      `yaml:"paths"`
	Components Components `yaml:"components"`
}

type Server struct {
	URL         string                    `yaml:"url"`
	Description string                    `yaml:"description"`
	Variables   map[string]ServerVariable `yaml:"variables"`
}

type ServerVariable struct {
	Enum        []string `yaml:"enum"`
	Default     string   `yaml:"default"`
	Description string   `yaml:"description"`
}

type Components struct {
	Schemas         map[string]Schema `yaml:"schemas"`
	SecuritySchemes map[string]any    `yaml:"securitySchemes"`
}

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
	Summary     string           `yaml:"summary,omitempty"`
	OperationId string           `yaml:"operationId,omitempty"`
	Tags        []string         `yaml:"tags,omitempty"`
	Parameters  []Parameter      `yaml:"parameters,omitempty"`
	RequestBody Ref[RequestBody] `yaml:"requestBody,omitempty"`
	Security    Security         `yaml:"security,omitempty"`
	Responses   Responses        `yaml:"responses,omitempty"`
}

type Ref[T any] struct {
	Reference *Reference `yaml:",inline"`
	Value     *T         `yaml:",inline"`
}

type RequestBodyOrReference struct {
	Reference   *Reference   `yaml:",inline"`
	RequestBody *RequestBody `yaml:",inline"`
}

type RequestBody struct {
	Description string               `yaml:"description,omitempty"`
	Content     map[string]MediaType `yaml:"content"`
}

type MediaType struct{}

type Parameter struct {
	Name            string      `yaml:"name"`
	In              In          `yaml:"in"`
	Description     string      `yaml:"description,omitempty"`
	Required        bool        `yaml:"required"`
	Deprecated      bool        `yaml:"deprecated"`
	AllowEmptyValue bool        `yaml:"allowEmptyValue"`
	Schema          Ref[Schema] `yaml:"schema,omitempty"`
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
	Type       string         `yaml:"type"`
	Minimum    *int64         `yaml:"minimum,omitempty"`
	Maximum    *int64         `yaml:"maximum,omitempty"`
	MaxItems   *int64         `yaml:"maxItems,omitempty"`
	Items      Ref[Schema]    `yaml:"items,omitempty"`
	Format     Format         `yaml:"format,omitempty"`
	Required   []string       `yaml:"required,omitempty"`
	Properties map[string]any `yaml:"properties"`
	Example    any            `yaml:"example"`
}

type Reference struct {
	Ref         string `yaml:"$ref"`
	Summary     string `yaml:"summary,omitempty"`
	Description string `yaml:"description,omitempty"`
}

func (ref *Ref[T]) UnmarshalYAML(node *yaml.Node) error {
	v := make(map[string]any)
	node.Decode(v) // TODO: can be done without decoding
	if _, ok := v["$ref"]; ok {
		ref.Reference = &Reference{}
		return node.Decode(ref.Reference)
	}
	var val T
	ref.Value = &val
	return node.Decode(ref.Value)
}
