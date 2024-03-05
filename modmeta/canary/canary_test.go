/*
 * Copyright (c) 2024, Jamie Mansfield <jmansfield@cadixdev.org>
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package canary

import (
	"strings"
	"testing"
)

func TestReadCanaryInf(t *testing.T) {
	plugin, err := ReadCanaryInf(strings.NewReader(`name=Demo Plugin
main-class=com.example.demo
author=Demo Author
version=0.1`))
	if err != nil {
		t.Error(err)
		return
	}

	if plugin.Name != "Demo Plugin" {
		t.Errorf("Plugin name should be example, not %s", plugin.Name)
	}
	if plugin.Author != "Demo Author" {
		t.Errorf("Plugin name should be example, not %s", plugin.Name)
	}
	if plugin.Version != "0.1" {
		t.Errorf("Plugin name should be example, not %s", plugin.Name)
	}
}
