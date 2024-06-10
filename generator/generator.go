package generator

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/o5h/openapi/spec"
)

func Generate(cfg *Config) error {
	gen := generator{cfg: cfg}
	return gen.generate()
}

type generator struct {
	openapi        *spec.OpenAPI
	cfg            *Config
	api            *API
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

	if err = g.applyTemplates(); err != nil {
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

func (g *generator) applyTemplates() (err error) {
	if err = g.applyAPITemplate(); err != nil {
		return
	}
	return
}

func (g *generator) applyAPITemplate() (err error) {
	j, _ := json.Marshal(g.api)
	fmt.Println(string(j))
	err = g.cfg.APITemplate.Execute(os.Stdout, g.api)
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
		switch schema.Value.Type {
		case "object":
			g.defineComponentObject(schema.Name, &schema.Value)
		case "array":
			g.defineComponentArray(schema.Name, &schema.Value)
		default:
			panic("Unsupported schema type " + schema.Value.Type)
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
	def := &TypeDef{Type: ObjectType, Name: toPascalCase(name)}
	for _, prop := range schema.Properties {
		field := Field{
			Name:     toPascalCase(prop.Name),
			Required: isRequired(schema, prop.Name),
		}
		field.Type = g.resolvePropertyType(&prop.Value)
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
		for _, method := range path.Value {
			for _, tag := range method.Value.Tags {
				if tag != "" {
					return tag
				}
			}
		}
	}
	return ""
}
