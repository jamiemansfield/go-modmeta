package modmeta

import (
	"strings"
	"testing"
)

var (
	testLiteModJson = `{
	"name": "example",
	"displayName": "Example Mod",
	"version": "1.0.0",
	"mcversion": "1.12.2",
	"author": "Bob, Vance",
	"description": "Example Mod.",
	"url": "https://examplemod.com"
}`
)

func TestReadLiteModJson(t *testing.T) {
	mod, err := ReadLiteModJson(strings.NewReader(testLiteModJson))
	if err != nil {
		t.Error(err)
		return
	}
	testModMetadata(t, mod.ToModMetadata())
}
