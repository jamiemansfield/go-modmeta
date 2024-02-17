/*
 * Copyright (c) 2020, Jamie Mansfield <jmansfield@cadixdev.org>
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package modmeta

import (
	"encoding/json"
	"io"
	"strings"
)

// FabricModJsonMetadata represents a single mods' metadata from Fabric's
// fabric.mod.json standard.
type FabricModJsonMetadata struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Description string   `json:"description"`
	Authors     []string `json:"authors"`
	Contact     struct {
		Homepage string `json:"homepage"`
	} `json:"contact"`
}

// ReadFabricModJson rads a fabric.mod.json file.
func ReadFabricModJson(reader io.Reader) (*FabricModJsonMetadata, error) {
	var mod FabricModJsonMetadata
	err := json.NewDecoder(reader).Decode(&mod)
	if err != nil {
		return nil, err
	}

	return &mod, nil
}

// ToModMetadata creates a ModMetadata for the fabric.mod.json metadata.
// The System will be set to "fabric".
func (m *FabricModJsonMetadata) ToModMetadata() *ModMetadata {
	return &ModMetadata{
		System:      "fabric",
		ID:          m.ID,
		Name:        m.Name,
		Version:     m.Version,
		Description: m.Description,
		URL:         m.Contact.Homepage,
		Authors:     strings.Join(m.Authors, ", "),
	}
}
