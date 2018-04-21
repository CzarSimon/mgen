package pkg

import (
	"testing"
)

func TestReadSchema(t *testing.T) {
	schema, err := ReadSchema("../resources/exampleSchema.json")
	if err != nil {
		t.Fatalf("ReadSchema returned unexpected error: %s", err)
	}
	if schema.Title != "Photo" {
		t.Errorf("ReadSchema returned wrong Title. Expected=Photo Got=%s", schema.Title)
	}
}
