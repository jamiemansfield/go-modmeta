/*
 * Copyright (c) 2020-2024, Jamie Mansfield <jmansfield@cadixdev.org>
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package forge

import (
	"strings"
	"testing"

	"github.com/jamiemansfield/go-modmeta/modmeta"
)

var (
	testMcModInfoV1 = `[
  {
    "modid": "example",
    "name": "Example Mod",
    "description": "Example Mod.",
    "version": "1.0.0",
    "url": "https://examplemod.com",
    "authorList": ["Example Author"]
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
      "authorList": ["Example Author"]
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
	authors = "Example Author"
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

func testModMetadata(t *testing.T, mod *modmeta.ModMetadata) {
	if mod.ID != "example" {
		t.Errorf("Mod ID should be example, not %s", mod.ID)
	}
	if mod.Name != "Example Mod" {
		t.Errorf("Mod name should be 'Example Mod', not '%s'", mod.Name)
	}
	if mod.Version != "1.0.0" {
		t.Errorf("Mod version should be '1.0.0', not '%s'", mod.Version)
	}
	if mod.Description != "Example Mod." {
		t.Errorf("Mod description should be 'Example Mod.', not '%s'", mod.Description)
	}
	if mod.URL != "https://examplemod.com" {
		t.Errorf("Mod website should be 'https://examplemod.com', not '%s'", mod.URL)
	}
	if mod.Authors != "Example Author" {
		t.Errorf("Mod authors should be 'Example Author', not '%s'", mod.Authors)
	}
}
