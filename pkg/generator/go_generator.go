package generator

import (
	"bytes"
	"fmt"

	"github.com/CzarSimon/mgen/pkg"
)

// Go is a code genreator for the go language.
type Go struct{}

// NewGo creates and returns a new go generator.
func NewGo() Go {
	return Go{}
}

// Generate creates a go code for the model described in the schema.
func (g *Go) Generate(schema pkg.Schema) (string, error) {
	switch schema.Type {
	case pkg.ObjectType:
		return generateStruct(schema)
	default:
		return "", makeUnrecognizedTypeError(schema.Type)
	}
}

func generateStruct(schema pkg.Schema) (string, error) {
	var block bytes.Buffer
	block.WriteString(makeTypeHeader(schema, "struct") + " {\n")
	for name, prop := range schema.Properties {
		attribute := fmt.Sprintf("%s%s %s\n", Indent, name, prop.Type)
		block.WriteString(attribute)
	}
	block.WriteString("}")
	return block.String(), nil
}

func makeTypeHeader(schema pkg.Schema, typeName string) string {
	var header string
	if schema.Description != "" {
		header += fmt.Sprintf("// %s %s\n", schema.Title, schema.Description)
	}
	header += fmt.Sprintf("type %s %s", schema.Title, typeName)
	return header
}
