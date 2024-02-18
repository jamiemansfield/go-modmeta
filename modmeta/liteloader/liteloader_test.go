/*
 * Copyright (c) 2020-2024, Jamie Mansfield <jmansfield@cadixdev.org>
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package liteloader

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
	"author": "Example Author",
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

	if mod.Name != "example" {
		t.Errorf("Mod ID should be example, not %s", mod.Name)
	}
	if mod.DisplayName != "Example Mod" {
		t.Errorf("Mod name should be 'Example Mod', not '%s'", mod.DisplayName)
	}
	if mod.Version != "1.0.0" {
		t.Errorf("Mod version should be '1.0.0', not '%s'", mod.Version)
	}
	if mod.McVersion != "1.12.2" {
		t.Errorf("Mod version should be '1.12.2', not '%s'", mod.McVersion)
	}
	if mod.Author != "Example Author" {
		t.Errorf("Mod author should be 'Example Author', not '%s'", mod.Author)
	}
	if mod.Description != "Example Mod." {
		t.Errorf("Mod description should be 'Example Mod.', not '%s'", mod.Description)
	}
	if mod.URL != "https://examplemod.com" {
		t.Errorf("Mod website should be 'https://examplemod.com', not '%s'", mod.URL)
	}
}
