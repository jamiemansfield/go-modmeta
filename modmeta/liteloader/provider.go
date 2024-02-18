/*
 * Copyright (c) 2020-2024, Jamie Mansfield <jmansfield@cadixdev.org>
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package liteloader

import (
	"archive/zip"
	"errors"
	"io/fs"

	"git.sr.ht/~jmansfield/go-javamanifest/javamanifest"

	"github.com/jamiemansfield/go-modmeta/modmeta"
)

func init() {
	modmeta.RegisterProvider("liteloader", findMods)
}

func findMods(reader *zip.Reader, manifest *javamanifest.Manifest) ([]*modmeta.ModMetadata, error) {
	file, err := reader.Open("litemod.json")
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}
	defer file.Close()

	mod, err := ReadLiteModJson(file)
	if err != nil {
		return nil, err
	}

	return []*modmeta.ModMetadata{mod.ToModMetadata()}, nil
}
