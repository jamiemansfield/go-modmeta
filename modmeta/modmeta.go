// Package modmeta provides functionality to get mod metadata
// from mod binaries.
package modmeta

import "archive/zip"

// Represents a single mod's metadata.
type ModMetadata struct {
	// The mod system/loader the mods uses.
	System string

	ID string
	Name string
	Version string
	Description string
	URL string
	Authors string
}

// FindMetadata attempts to find mod information from a Java binary,
// looking for metadata from Minecraft Forge.
func FindMetadata(archive string) ([]*ModMetadata, error) {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	var mods []*ModMetadata
	for _, file := range reader.File {
		// Minecraft Forge / mcmod.info
		if file.Name == "mcmod.info" {
			fc, err := file.Open()
			if err != nil {
				return nil, err
			}

			forgeMods, err := ReadMcModInfo(fc)
			if err != nil {
				return nil, err
			}
			fc.Close()

			for _, mod := range forgeMods {
				mods = append(mods, mod.ToModMetadata())
			}
		}
	}

	return mods, nil
}
