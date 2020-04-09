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
    "authors": ["Bob", "Vance"]
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
      "authors": ["Bob", "Vance"]
    }
  ]
}`
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
}
