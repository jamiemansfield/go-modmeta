/*
 * Copyright (c) 2020, Jamie Mansfield <jmansfield@cadixdev.org>
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

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
	testModMetadata(t, mod.ToModMetadata())
}
