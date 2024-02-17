/*
 * Copyright (c) 2020, Jamie Mansfield <jmansfield@cadixdev.org>
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package modmeta

import "testing"

func testModMetadata(t *testing.T, mod *ModMetadata) {
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
	if mod.Authors != "Bob, Vance" {
		t.Errorf("Mod authors should be 'Bob, Vance', not '%s'", mod.Authors)
	}
}
