/*
 * Copyright (c) 2020-2024, Jamie Mansfield <jmansfield@cadixdev.org>
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package forge

import (
	"archive/zip"
	"errors"
	"io/fs"

	"git.sr.ht/~jmansfield/go-javamanifest/javamanifest"

	"github.com/jamiemansfield/go-modmeta/modmeta"
)

func init() {
	modmeta.RegisterProvider("forge", findMods)
}

func findMods(reader *zip.Reader, manifest *javamanifest.Manifest) ([]*modmeta.ModMetadata, error) {
	var mods []*modmeta.ModMetadata

	// mcmod.info
	{
		file, err := reader.Open("mcmod.info")
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				goto modernFormat
			}
			return nil, err
		}
		defer file.Close()

		oldMods, err := ReadMcModInfo(file)
		if err != nil {
			return nil, err
		}

		mods = append(mods, oldMods...)
	}

modernFormat:
	// mods.toml
	{
		file, err := reader.Open("META-INF/mods.toml")
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				goto finish
			}
			return nil, err
		}
		defer file.Close()

		newMods, err := ReadForgeModsToml(file)
		if err != nil {
			return nil, err
		}

		// Minecraft Forge supports substitutions in mods.toml files,
		// with data populated from the Jar's MANIFEST.
		// Substitutions are in the form ${file.KEY}.
		for _, mod := range newMods {
			if mod.Version == "${file.jarVersion}" {
				manifestVersion := manifest.Main["Implementation-Version"]
				if manifestVersion == "" {
					// This matches Minecraft Forge's behaviour
					manifestVersion = "NONE"
				}

				mod.Version = manifestVersion
			}
		}

		mods = append(mods, newMods...)
	}

finish:
	return mods, nil
}
