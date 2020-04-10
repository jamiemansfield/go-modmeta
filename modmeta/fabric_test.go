package modmeta

import (
	"strings"
	"testing"
)

var (
	testFabricModJson = `{
	"id": "example",
	"name": "Example Mod",
	"version": "1.0.0",
	"description": "Example Mod.",
	"contact": {
		"homepage": "https://examplemod.com"
	},
    "authors": ["Bob", "Vance"]
}`
)

func TestReadFabricModJson(t *testing.T) {
	mod, err := ReadFabricModJson(strings.NewReader(testFabricModJson))
	if err != nil {
		t.Error(err)
		return
	}

	if mod.ID != "example" {
		t.Errorf("Mod ID should be example, not %s", mod.ID)
	}
}
