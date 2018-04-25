package pkg

import (
	"testing"
)

var testSchema = []byte(`
{
  "title": "Photo",
  "description": "describes photo metadata",
  "type": "object",
  "properties": {
    "contentType": {
      "type": "string"
    },
    "content": {
      "type": "string"
    },
    "size": {
      "type": "int"
    }
  },
  "required": ["contentType", "content"],
  "options": {
    "java": {
      "className": "PhotoDTO",
      "package": "com.foo.dto"
    },
    "go": {
      "package": "schema",
      "tags": ["json", "yaml"]
    }
  }
}
`)

func TestReadSchema(t *testing.T) {
	schema, err := ReadSchema("../resources/exampleSchema.json")
	if err != nil {
		t.Fatalf("ReadSchema returned unexpected error: %s", err)
	}
	if schema.Title != "Photo" {
		t.Errorf("ReadSchema returned wrong Title. Expected=Photo Got=%s", schema.Title)
	}
}

func TestParseGoOpts(t *testing.T) {
	schema := getTestSchema(t)
	opts := ParseGoOpts(schema.Options[Go])
	if opts.Package != "schema" {
		t.Errorf("ParseGoOpts returned wrong Package name. Expected=schema Got=%s",
			opts.Package)
	}
	if len(opts.Tags) != 2 {
		t.Errorf("ParseGoOpts returned wrong number of Tags name. Expected=2 Got=%d",
			len(opts.Tags))
	}
	expectedTags := []string{"json", "yaml"}
	for i, tag := range opts.Tags {
		if tag != expectedTags[i] {
			t.Errorf("%d - Wrong tag. Expected=%s Got=%s", i, expectedTags[i], tag)
		}
	}
}

func getTestSchema(t *testing.T) Schema {
	schema, err := NewSchema(testSchema)
	if err != nil {
		t.Fatalf("NewSchema returned unexpected error: %s", err)
	}
	return schema
}
