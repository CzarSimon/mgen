package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"strings"

	"github.com/CzarSimon/mgen/pkg"
)

// Go is a code genreator for the go language.
type Go struct {
	typeMap typeMap
	opts    pkg.GoOptions
}

// NewGo creates and returns a new go generator.
func NewGo(opts interface{}) Go {
	return Go{
		typeMap: newGoMap(),
		opts:    pkg.ParseGoOpts(opts),
	}
}

// Generate creates a go code for the model described in the schema.
func (g Go) Generate(schema pkg.Schema) (string, error) {
	header := g.makeTypeHeader(schema)
	body, err := g.makeTypeBody(schema.Type, schema.Properties)
	if err != nil {
		return "", err
	}
	model := fmt.Sprintf("%s %s", header, body)
	formatedModel, err := format.Source([]byte(model))
	return string(formatedModel), err
}

func (g Go) makeTypeBody(typeName pkg.TypeName, props map[string]pkg.Property) (string, error) {
	if typeName == pkg.ObjectType {
		return g.generateStruct(props)
	}
	goType := g.typeMap.Map(typeName)
	if goType == "" {
		return "", makeUnrecognizedTypeError(typeName)
	}
	return goType, nil
}

func (g Go) generateStruct(properties map[string]pkg.Property) (string, error) {
	var block bytes.Buffer
	block.WriteString("struct {")
	if len(properties) != 0 {
		block.WriteString("\n")
	}
	for name, prop := range properties {
		typeBody, err := g.makeTypeBody(prop.Type, prop.Properties)
		if err != nil {
			return "", err
		}
		attribute := fmt.Sprintf("%s %s %s\n",
			strings.Title(name), typeBody, g.makeTags(name))
		block.WriteString(attribute)
	}
	block.WriteString("}")
	return block.String(), nil
}

func (g Go) makeTypeHeader(schema pkg.Schema) string {
	header := fmt.Sprintf("package %s\n\n", g.opts.Package)
	if schema.Description != "" {
		header += fmt.Sprintf("// %s %s\n", schema.Title, schema.Description)
	}
	header += fmt.Sprintf("type %s", schema.Title)
	return header
}

func (g Go) makeTags(name string) string {
	tagName := fmt.Sprintf(":\"%s\"", name)
	tagsLen := len(g.opts.Tags)
	if tagsLen == 0 {
		return ""
	}
	tags := make([]string, tagsLen)
	for i, tag := range g.opts.Tags {
		tags[i] = tag + tagName
	}
	return fmt.Sprintf("`%s`", strings.Join(tags, " "))
}

// goMap type map for golang.
type goMap struct {
	types map[pkg.TypeName]string
}

func (m goMap) Map(typeName pkg.TypeName) string {
	langType, ok := m.types[typeName]
	if !ok {
		return ""
	}
	return langType
}

func newGoMap() goMap {
	typeMap := map[pkg.TypeName]string{
		pkg.ObjectType:   "struct",
		pkg.DatetimeType: "time.Time",
		pkg.IntegerType:  "int",
		pkg.FloatType:    "float64",
		pkg.StringType:   "string",
	}
	return goMap{types: typeMap}
}
