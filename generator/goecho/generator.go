package goecho

import (
	_ "embed"

	"github.com/o5h/openapi/generator"
)

var DefaultTypeMap = map[generator.TypeFormat]string{
	{Type: "integer", Format: "int32"}: "int32",
	{Type: "integer", Format: "int64"}: "int64",
	{Type: "string", Format: ""}:       "string",
}

//go:embed api.tmpl
var DefaultTemplate string
