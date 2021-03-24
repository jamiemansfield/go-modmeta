// Package modmeta provides functionality to get mod metadata
// from mod binaries.
package modmeta

import (
	"archive/zip"

	"git.sr.ht/~jmansfield/go-javamanifest/javamanifest"
)

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

	// Find MANIFEST first
	manifest := javamanifest.NewManifest()
	for _, file := range reader.File {
		if file.Name != "META-INF/MANIFEST.MF" {
			continue
		}
		fc, err := file.Open()
		if err != nil {
			return nil, err
		}

		if err := manifest.Read(fc); err != nil {
			return nil, err
		}
	}

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

			mods = append(mods, forgeMods...)
		}

		// Minecraft Forge / mods.toml
		if file.Name == "META-INF/mods.toml" {
			fc, err := file.Open()
			if err != nil {
				return nil, err
			}

			forgeMods, err := ReadForgeModsToml(fc)
			if err != nil {
				return nil, err
			}
			fc.Close()

			// Minecraft Forge supports substitutions in mods.toml files,
			// with data populated from the Jar's MANIFEST.
			// Substitutions are in the form ${file.KEY}.
			for _, mod := range forgeMods {
				if mod.Version == "${file.jarVersion}" {
					manifestVersion := manifest.Main["Implementation-Version"]
					if manifestVersion == "" {
						// This matches Minecraft Forge's behaviour
						manifestVersion = "NONE"
					}

					mod.Version = manifestVersion
				}
			}

			mods = append(mods, forgeMods...)
		}

		// Fabric / fabric.mod.json
		if file.Name == "fabric.mod.json" {
			fc, err := file.Open()
			if err != nil {
				return nil, err
			}

			fabricMod, err := ReadFabricModJson(fc)
			if err != nil {
				return nil, err
			}
			fc.Close()

			mods = append(mods, fabricMod.ToModMetadata())
		}

		// LiteLoader / litemod.json
		if file.Name == "litemod.json" {
			fc, err := file.Open()
			if err != nil {
				return nil, err
			}

			liteMod, err := ReadLiteModJson(fc)
			if err != nil {
				return nil, err
			}
			fc.Close()

			mods = append(mods, liteMod.ToModMetadata())
		}
	}

	return mods, nil
}
