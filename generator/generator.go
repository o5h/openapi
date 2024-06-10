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
	openapi        *spec.OpenAPI
	cfg            *Config
	api            *API
	tmpl           *template.Template
	types          map[string]*TypeDef
	typeRefLinkage []func()
}

func (g *generator) generate() (err error) {

	g.types = make(map[string]*TypeDef)

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
	g.linkTypeRefs()
	return
}
func (g *generator) linkTypeRefs() {
	for _, link := range g.typeRefLinkage {
		link()
	}
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
			g.defineComponentObject(schema.Name, &schema.Schema)
		case "array":
			g.defineComponentArray(schema.Name, &schema.Schema)
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

func (g *generator) defineObject(name string, schema *spec.Schema) *TypeDef {
	def := &TypeDef{Name: toPascalCase(name)}
	for _, prop := range schema.Properties {
		field := Field{
			Name:     toPascalCase(prop.Name),
			Required: isRequired(schema, prop.Name),
		}
		field.Type = g.resolvePropertyType(&prop.Schema)
		def.Fields = append(def.Fields, field)
	}
	return def
}

func (g *generator) defineComponentObject(name string, schema *spec.Schema) {
	def := g.defineObject(name, schema)
	g.types["#/components/schemas/"+name] = def
	g.api.TypesDefs = append(g.api.TypesDefs, def)
}

func (g *generator) resolvePropertyType(schema *spec.Schema) string {
	//TODO: $ref support
	typeFormat := TypeFormat{Type: schema.Type, Format: schema.Format}
	if targetType, ok := g.cfg.TypeMap[typeFormat]; ok {
		return targetType
	}
	return schema.Type
}

func (g *generator) defineComponentArray(name string, schema *spec.Schema) {
	def := &TypeDef{
		Type: ArrayType,
		Name: toPascalCase(name),
	}
	if schema.Items.Reference != nil {
		ref := schema.Items.Reference.Ref
		g.typeRefLinkage = append(g.typeRefLinkage, func() {
			def.ItemsType = g.types[ref] // resolve type later
		})
	} else {
		//TODO
		panic("not implemented yet")
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
