package pkg

import (
	"encoding/json"
	"io/ioutil"
)

// Names of supported languages
const (
	Java       Language = "java"
	Go                  = "go"
	Python              = "python"
	Javascript          = "javascript"
)

// Supported types
const (
	ObjectType   TypeName = "object"
	StringType            = "string"
	IntegerType           = "int"
	FloatType             = "float"
	DatetimeType          = "dataTime"
)

// A Schema describes a data entity for which source code should be genereated.
type Schema struct {
	Title       string                   `json:"title" yaml:"title"`
	Description string                   `json:"description" yaml:"description"`
	Type        TypeName                 `json:"type" yaml:"type"`
	Properties  map[string]Property      `json:"properties" yaml:"properties"`
	Required    []string                 `json:"required" yaml:"required"`
	Options     map[Language]interface{} `json:"options" yaml:"options"`
}

// NewSchema creates a new schema based on a raw byte array.
func NewSchema(data []byte) (Schema, error) {
	var schema Schema
	err := json.Unmarshal(data, &schema)
	return schema, err
}

// ReadSchema reads a schema from a source file.
func ReadSchema(filename string) (Schema, error) {
	var schema Schema
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return schema, err
	}
	return NewSchema(data)
}

// TypeName is the name of a supported type.
type TypeName string

// Language is the name of a supported target language.
type Language string

// Property describes the attributes of an individial schema property.
type Property struct {
	Type        TypeName            `json:"type" yaml:"type"`
	Description string              `json:"description,omitempty" yaml:"description,omitempty"`
	Properties  map[string]Property `json:"properties,omitempty" yaml:"properties,omitempty"`
}

// JavaOptions are addiontonal instructctions for generating java files.
type JavaOptions struct {
	Package     string   `json:"package" yaml:"package"`
	ClassName   string   `json:"className" yaml:"className"`
	Annotations []string `json:"annotations" yaml:"annotations"`
}

// GoOptions are addiontonal instructions for generation go structs.
type GoOptions struct {
	Package string   `json:"package" yaml:"package"`
	Tags    []string `json:"tags" yaml:"tags"`
}

// DefaultGoOptions is the default go options.
var DefaultGoOptions = GoOptions{
	Package: "main",
	Tags:    []string{"json"},
}

// ParseGoOpts attempts to parse go options and returns default if unsuccessfull.
func ParseGoOpts(opts interface{}) GoOptions {
	goOpts := GoOptions{}
	optsMap, ok := opts.(map[string]interface{})
	if !ok {
		return DefaultGoOptions
	}
	goOpts.Package, ok = optsMap["package"].(string)
	if !ok {
		goOpts.Package = DefaultGoOptions.Package
	}
	goOpts.Tags, ok = castToStringSlice(optsMap["tags"])
	if !ok {
		goOpts.Tags = DefaultGoOptions.Tags
	}
	return goOpts
}

func castToStringSlice(v interface{}) ([]string, bool) {
	elements, ok := v.([]interface{})
	if !ok {
		return nil, false
	}
	strSlice := make([]string, len(elements))
	for i, elem := range elements {
		strSlice[i], ok = elem.(string)
		if !ok {
			return nil, false
		}
	}
	return strSlice, true
}
