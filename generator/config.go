package generator

type TypeFormat struct {
	Type   string
	Format string
}

type Config struct {
	OpenAPIFile  string
	Package      string
	TemplateFile string
	TypeMap      map[TypeFormat]string
}

var DefaultTypeMap = map[TypeFormat]string{
	{Type: "integer", Format: "int32"}: "int32",
	{Type: "integer", Format: "int64"}: "int64",
	{Type: "string", Format: ""}:       "string",
}
