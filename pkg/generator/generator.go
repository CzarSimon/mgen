package generator

import (
	"fmt"

	"github.com/CzarSimon/mgen/pkg"
)

// Formating constants.
const (
	Indent = "    "
)

// Generator is the main interface for code generators based on schema.
type Generator interface {
	Generate(pkg.Schema) (string, error)
}

// GenerateModel function signature of code generators.
type GenerateModel func(pkg.Schema) (string, error)

// typeMap is an inteface for mapping mgen supported types
// to their target language equivalents.
type typeMap interface {
	Map(pkg.TypeName) string
}

func makeUnrecognizedTypeError(typeName pkg.TypeName) error {
	return fmt.Errorf("Unsupported type: '%s'", typeName)
}
