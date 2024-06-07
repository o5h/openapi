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
