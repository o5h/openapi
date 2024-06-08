package spec

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Method string
type In string
type Type string

const (
	MethodGet     = Method("get")
	MethodHead    = Method("head")
	MethodPost    = Method("post")
	MethodPut     = Method("put")
	MethodPatch   = Method("patch")
	MethodDelete  = Method("delete")
	MethodConnect = Method("connect")
	MethodOptions = Method("options")
	MethodTrace   = Method("trace")

	TypeInteger = Type("integer")
	TypeArray   = Type("array")
	TypeString  = Type("string")

	InQuery  = In("query")
	InHeader = In("header")
	InPath   = In("path")
	InCookie = In("cookie")
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
	Schemas         NamedSchemas   `yaml:"schemas"`
	SecuritySchemes map[string]any `yaml:"securitySchemes"`
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
type Paths map[string]PathItem
type PathItem map[string]*Operation

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
	Type       string       `yaml:"type"`
	Minimum    *int64       `yaml:"minimum,omitempty"`
	Maximum    *int64       `yaml:"maximum,omitempty"`
	MaxItems   *int64       `yaml:"maxItems,omitempty"`
	Items      Ref[Schema]  `yaml:"items,omitempty"`
	Format     string       `yaml:"format,omitempty"`
	Required   []string     `yaml:"required,omitempty"`
	Properties NamedSchemas `yaml:"properties"`
	Example    any          `yaml:"example"`
}

type NamedSchema struct {
	Schema
	Name string
}
type NamedSchemas []NamedSchema

func (shema *NamedSchemas) UnmarshalYAML(node *yaml.Node) error {
	if node.Kind != yaml.MappingNode {
		return fmt.Errorf("`commands` must contain YAML mapping, has %v", node.Kind)
	}
	*shema = make([]NamedSchema, len(node.Content)/2)
	for i := 0; i < len(node.Content); i += 2 {
		var res = &(*shema)[i/2]
		if err := node.Content[i].Decode(&res.Name); err != nil {
			return err
		}
		if err := node.Content[i+1].Decode(&res.Schema); err != nil {
			return err
		}
	}
	return nil
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
