package generator

import "text/template"

type TypeFormat struct {
	Type   string
	Format string
}

type Config struct {
	OpenAPIFile string
	Package     string
	TypeMap     map[TypeFormat]string
	APITemplate *template.Template
}
