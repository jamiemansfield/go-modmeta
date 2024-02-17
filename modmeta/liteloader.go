/*
 * Copyright (c) 2020, Jamie Mansfield <jmansfield@cadixdev.org>
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package modmeta

import (
	"encoding/json"
	"io"
)

// LiteModMetadata represents a single mods' metadata from LiteLoaders's
// litemod.json standard.
type LiteModMetadata struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Version     string `json:"version"`
	McVersion   string `json:"mcversion"`
	Author      string `json:"author"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// ReadLiteModJson reads a litemod.json file.
func ReadLiteModJson(reader io.Reader) (*LiteModMetadata, error) {
	var mod LiteModMetadata
	err := json.NewDecoder(reader).Decode(&mod)
	if err != nil {
		return nil, err
	}

	return &mod, nil
}

// ToModMetadata creates a ModMetadata for the litemod.json metadata.
// The System will be set to "liteloader".
func (m *LiteModMetadata) ToModMetadata() *ModMetadata {
	return &ModMetadata{
		System: "liteloader",
		ID:     m.Name,
		// See http://develop.liteloader.com/liteloader/LiteLoader/blob/master/src/main/java/com/mumfrey/liteloader/core/api/LoadableModFile.java#L182
		Name: getDefaultedString(m.DisplayName, m.Name),
		// See http://develop.liteloader.com/liteloader/LiteLoader/blob/master/src/main/java/com/mumfrey/liteloader/core/api/LoadableModFile.java#L183
		Version:     getDefaultedString(m.Version, "Unknown"),
		Description: m.Description,
		URL:         m.URL,
		// See http://develop.liteloader.com/liteloader/LiteLoader/blob/master/src/main/java/com/mumfrey/liteloader/core/api/LoadableModFile.java#L184
		Authors: getDefaultedString(m.Author, "Unknown"),
	}
}

func getDefaultedString(value string, otherwise string) string {
	if value == "" {
		return otherwise
	}
	return value
}
