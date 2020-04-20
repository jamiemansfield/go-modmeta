package modmeta

import (
	"strings"
	"testing"
)

var (
	testMcModInfoV1 = `[
  {
    "modid": "example",
    "name": "Example Mod",
    "description": "Example Mod.",
    "version": "1.0.0",
    "url": "https://examplemod.com",
    "authorList": ["Bob", "Vance"]
  }
]`
	testMcModInfoV2 = `{
  "modList": [
    {
      "modid": "example",
      "name": "Example Mod",
      "description": "Example Mod.",
      "version": "1.0.0",
      "url": "https://examplemod.com",
      "authorList": ["Bob", "Vance"]
    }
  ]
}`
	testBadButValidMcModInfo = `{
  "modList": [
    {
      "description": "Example
 Mod."
    }
  ]
}`
	testModsToml = `modLoader="javaFml"
[[mods]]
	modId = "example"
	version = "1.0.0"
	displayName = "Example Mod"
	authors = "Bob, Vance"
	description = "Example Mod."
	displayURL = "https://examplemod.com"
`
)

func TestReadMcModInfo_V1(t *testing.T) {
	mods, err := ReadMcModInfo(strings.NewReader(testMcModInfoV1))
	if err != nil {
		t.Error(err)
		return
	}

	if len(mods) != 1 {
		t.Errorf("There should be 1 mod not %d", len(mods))
		return
	}
	testModMetadata(t, mods[0])
}

func TestReadMcModInfo_V2(t *testing.T) {
	mods, err := ReadMcModInfo(strings.NewReader(testMcModInfoV2))
	if err != nil {
		t.Error(err)
		return
	}

	if len(mods) != 1 {
		t.Errorf("There should be 1 mod not %d", len(mods))
		return
	}
	testModMetadata(t, mods[0])
}

func TestReadMcModInfo_BadButValid(t *testing.T) {
	mods, err := ReadMcModInfo(strings.NewReader(testBadButValidMcModInfo))
	if err != nil {
		t.Error(err)
		return
	}

	if len(mods) != 1 {
		t.Errorf("There should be 1 mod not %d", len(mods))
		return
	}

	if mods[0].Description != "Example\n Mod." {
		t.Errorf("Failed to read multi-line string!")
	}
}

func TestReadForgeModsToml(t *testing.T) {
	mods, err := ReadForgeModsToml(strings.NewReader(testModsToml))
	if err != nil {
		t.Error(err)
		return
	}

	if len(mods) != 1 {
		t.Errorf("There should be 1 mod not %d", len(mods))
		return
	}
	testModMetadata(t, mods[0])
}
