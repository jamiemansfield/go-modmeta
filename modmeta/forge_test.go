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

func TestReadMcModInfoV1(t *testing.T) {
	mods, err := ReadMcModInfoV1(strings.NewReader(testMcModInfoV1))
	if err != nil {
		t.Error(err)
		return
	}

	if len(mods) != 1 {
		t.Errorf("There should be 1 mod not %d", len(mods))
		return
	}
	testModMetadata(t, mods[0].ToModMetadata())
}

func TestReadMcModInfoV2(t *testing.T) {
	mods, err := ReadMcModInfoV2(strings.NewReader(testMcModInfoV2))
	if err != nil {
		t.Error(err)
		return
	}

	if len(mods) != 1 {
		t.Errorf("There should be 1 mod not %d", len(mods))
		return
	}
	testModMetadata(t, mods[0].ToModMetadata())
}

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
	testModMetadata(t, mods[0].ToModMetadata())
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
	testModMetadata(t, mods[0].ToModMetadata())
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
