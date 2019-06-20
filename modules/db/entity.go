package db

import (
	"fmt"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"

	dbext "go.zenithar.org/protoc-gen-blueprint/protobuf/blueprint/db"
)

// EntityModule is used to generate database models
type EntityModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	tpl *template.Template
}

// Entity returns an initialized "entity" module
func Entity() *EntityModule { return &EntityModule{ModuleBase: &pgs.ModuleBase{}} }

// -----------------------------------------------------------------------------

// InitContext is used to iniatile template and function
func (p *EntityModule) InitContext(c pgs.BuildContext) {
	p.ModuleBase.InitContext(c)
	p.ctx = pgsgo.InitContext(c.Parameters())

	tpl := template.New("entity").Funcs(map[string]interface{}{
		"package":      p.ctx.PackageName,
		"name":         p.ctx.Name,
		"description":  p.description,
		"has_entity":   p.hasEntity,
		"entity":       p.entity,
		"ftype":        p.ctx.Type,
		"fdescription": p.fdescription,
	})

	p.tpl = template.Must(tpl.Parse(entityTpl))
}

// Name satisfies the generator.Plugin interface.
func (p *EntityModule) Name() string { return "entity" }

// Execute is used to merge template with model
func (p *EntityModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {

	for _, t := range targets {
		p.generate(t)
	}

	return p.Artifacts()
}

// -----------------------------------------------------------------------------

func (p *EntityModule) hasEntity(m pgs.Message) bool {
	_, ok := ExtEntity(m)
	return ok
}

func (p *EntityModule) entity(m pgs.Message) *dbext.Entity {
	res, _ := ExtEntity(m)
	return res
}

func (p *EntityModule) description(m pgs.Message) pgs.Name {
	ext, ok := ExtEntity(m)
	if !ok {
		return "..."
	}
	return pgs.Name(ext.Description)
}

func (p *EntityModule) fdescription(f pgs.Field) pgs.Name {
	ext, ok := ExtAttribute(f)
	if !ok {
		return ""
	}
	return pgs.Name(ext.Description)
}

// -----------------------------------------------------------------------------

// ExtEntity check for entity extension
func ExtEntity(m pgs.Message) (*dbext.Entity, bool) {
	opts := &dbext.Entity{}
	ok, _ := m.Extension(dbext.E_Entity, opts)
	if !ok {
		return nil, false
	}
	return opts, true
}

// ExtAttribute check for attribute extension
func ExtAttribute(f pgs.Field) (*dbext.Attribute, bool) {
	opts := &dbext.Attribute{}
	ok, _ := f.Extension(dbext.E_Attribute, opts)
	if !ok {
		return nil, false
	}
	return opts, true
}

// -----------------------------------------------------------------------------

func (p *EntityModule) generate(f pgs.File) {
	if len(f.Messages()) == 0 {
		return
	}

	// For all messages
	for _, m := range f.Messages() {
		if p.hasEntity(m) && p.entity(m).Enabled {
			// If entity is enabled or globally enabled
			name := fmt.Sprintf("%s.go", m.Name().LowerSnakeCase())
			p.AddGeneratorTemplateFile(name, p.tpl, m)
		}
	}
}

// -----------------------------------------------------------------------------

const entityTpl = `package {{ package . }}

// {{ name . }} {{ description . }}
type {{ name . }} struct {
{{ range .Fields }}	// {{ .Name.UpperCamelCase }} {{ fdescription . }}
	{{ .Name.UpperCamelCase }} {{ ftype . }}
{{ end }}
}

// New returns an initialized entity with default values
func New{{ name .}}() *{{ name .}} {
	return &{{ name . }}{}
}

// -----------------------------------------------------------------------------

// Validate model according model constraints
func (m *{{ name . }}) Validate() error {
	return nil
}
`
