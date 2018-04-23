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
	ObjectType = "object"
	StringType = "string"
)

// A Schema describes a data entity for which source code should be genereated.
type Schema struct {
	Title       string                   `json:"title" yaml:"title"`
	Description string                   `json:"description" yaml:"description"`
	Type        string                   `json:"type" yaml:"type"`
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

// Language is the name of a supported target language.
type Language string

// Property describes the attributes of an individial schema property.
type Property struct {
	Type        string `json:"type" yaml:"type"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
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
