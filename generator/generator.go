package generator

import (
	_ "embed"
	"html/template"
	"os"

	"github.com/o5h/openapi/spec"
)

//go:embed default.tmpl
var defaultTemplate string

type generator struct {
	openapi *spec.OpenAPI
	cfg     *Config
	api     *API
	tmpl    *template.Template
}

func (g *generator) generate() (err error) {

	if err = g.loadOpenAPI(); err != nil {
		return
	}

	if err = g.generateAPI(); err != nil {
		return
	}

	if err = g.applyTemplate(); err != nil {
		return
	}

	return nil
}

func (g *generator) loadOpenAPI() error {
	var err error
	g.openapi, err = spec.Load(g.cfg.OpenAPIFile)
	return err
}

func (g *generator) generateAPI() (err error) {
	g.api = &API{}
	g.definePackageName()
	g.defineAPIName()
	g.generateModel()
	return
}

func (g *generator) applyTemplate() (err error) {
	if err = g.prepareTemplate(); err != nil {
		return
	}
	if err = g.executeTemplate(); err != nil {
		return
	}
	return
}

func (g *generator) prepareTemplate() (err error) {
	tmpl := defaultTemplate
	if g.cfg.TemplateFile != "" {
		tmpl = ""
	}
	g.tmpl, err = template.New("api").Parse(tmpl)
	return
}

func (g *generator) executeTemplate() (err error) {
	err = g.tmpl.Execute(os.Stdout, g.api)
	return
}

func (g *generator) definePackageName() {
	if g.cfg.Package != "" {
		g.api.Package = g.cfg.Package
		return
	}
	packageName := normalizePackageName(g.findFirstTag())
	if packageName != "" {
		g.api.Package = packageName
		return
	}
	g.api.Package = "api"
}

func (g *generator) defineAPIName() {
	g.api.APIName = toPascalCase(g.api.Package)
}

func (g *generator) generateModel() {
	g.convertComponentsToTypeDefs()
}

func (g *generator) convertComponentsToTypeDefs() {
	for _, schema := range g.openapi.Components.Schemas {
		switch schema.Type {
		case "object":
			g.defineObject(schema.Name, &schema.Schema)
		case "array":
			g.defineArray(schema.Name, &schema.Schema)
		default:
			panic("Unsupported schema type " + schema.Type)
		}
	}
}

func isRequired(schema *spec.Schema, name string) bool {
	for _, n := range schema.Required {
		if n == name {
			return true
		}
	}
	return false
}

func (g *generator) defineObject(name string, schema *spec.Schema) {
	typedef := TypeDef{Name: toPascalCase(name)}
	for _, prop := range schema.Properties {
		field := Field{
			Name:     toPascalCase(prop.Name),
			Required: isRequired(schema, name),
		}
		field.Type = g.resolvePropertyType(&prop.Schema)
		typedef.Fields = append(typedef.Fields, field)
	}
	g.api.TypesDefs = append(g.api.TypesDefs, typedef)
}

func (g *generator) resolvePropertyType(schema *spec.Schema) string {
	//TODO: $ref support
	typeFormat := TypeFormat{Type: schema.Type, Format: schema.Format}
	if targetType, ok := g.cfg.TypeMap[typeFormat]; ok {
		return targetType
	}
	return schema.Type
}

func (g *generator) defineArray(name string, schema *spec.Schema) {
	def := TypeDef{
		Type: ArrayType,
		Name: toPascalCase(name),
	}
	if schema.Items.Reference != nil {
		//TODO: resolve item type ref
	}
	g.api.TypesDefs = append(g.api.TypesDefs, def)
}

func (g *generator) findFirstTag() string {
	for _, path := range g.openapi.Paths {
		for _, method := range path {
			for _, tag := range method.Tags {
				if tag != "" {
					return tag
				}
			}
		}
	}
	return ""
}
