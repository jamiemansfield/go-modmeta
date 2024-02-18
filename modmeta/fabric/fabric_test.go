/*
 * Copyright (c) 2020-2024, Jamie Mansfield <jmansfield@cadixdev.org>
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package fabric

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
    "authors": ["Example Author"]
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
	if mod.Name != "Example Mod" {
		t.Errorf("Mod name should be 'Example Mod', not '%s'", mod.Name)
	}
	if mod.Version != "1.0.0" {
		t.Errorf("Mod version should be '1.0.0', not '%s'", mod.Version)
	}
	if mod.Description != "Example Mod." {
		t.Errorf("Mod description should be 'Example Mod.', not '%s'", mod.Description)
	}
	if mod.Contact.Homepage != "https://examplemod.com" {
		t.Errorf("Mod website should be 'https://examplemod.com', not '%s'", mod.Contact.Homepage)
	}
	if len(mod.Authors) != 1 {
		t.Errorf("Mod authors should have a length of 1, not '%d'", len(mod.Authors))
	} else {
		if mod.Authors[0] != "Example Author" {
			t.Errorf("Mod authors should be 'Example Author', not '%s'", mod.Authors[0])
		}
	}
}
