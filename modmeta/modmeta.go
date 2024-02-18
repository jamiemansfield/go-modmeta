/*
 * Copyright (c) 2020-2024, Jamie Mansfield <jmansfield@cadixdev.org>
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

// Package modmeta provides functionality to get mod metadata
// from mod binaries.
package modmeta

import (
	"archive/zip"
	"errors"
	"fmt"
	"io/fs"

	"git.sr.ht/~jmansfield/go-javamanifest/javamanifest"
)

var providers = make(map[string]ProviderFunction)

// ProviderFunction scans for mods within a given ZIP archive (likely
// to be a Java archive).
type ProviderFunction func(reader *zip.Reader, manifest *javamanifest.Manifest) ([]*ModMetadata, error)

// RegisterProvider registers a mod metadata provider, that is able to
// scan ZIP archives for mods.
func RegisterProvider(name string, f ProviderFunction) {
	if f == nil {
		panic(fmt.Errorf("modmeta: provider is nil"))
	}
	if _, ok := providers[name]; ok {
		panic(fmt.Errorf("modmeta: duplicate provider for %s", name))
	}

	providers[name] = f
}

// ModMetadata represents a single mod's metadata.
type ModMetadata struct {
	// The mod system/loader the mods uses.
	System string

	ID          string
	Name        string
	Version     string
	Description string
	URL         string
	Authors     string
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

	{
		file, err := reader.Open("META-INF/MANIFEST.MF")
		if err == nil {
			defer file.Close()

			if err := manifest.Read(file); err != nil {
				return nil, err
			}
		} else if !errors.Is(err, fs.ErrNotExist) {
			return nil, err
		}
	}

	var mods []*ModMetadata

	for _, function := range providers {
		provided, err := function(&reader.Reader, manifest)
		if err != nil {
			return nil, err
		}
		mods = append(mods, provided...)
	}

	return mods, nil
}
