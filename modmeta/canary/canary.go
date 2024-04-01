/*
 * Copyright (c) 2024, Jamie Mansfield <jmansfield@cadixdev.org>
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package canary

import (
	"io"
	"strings"

	"github.com/jamiemansfield/go-modmeta/modmeta"
	"github.com/jamiemansfield/go-modmeta/modmeta/canary/viutils"
)

type PluginDescriptor struct {
	Name         string
	Version      string
	Author       string
	Language     string
	EnableEarly  bool
	Dependencies []string
}

func (d *PluginDescriptor) ToModMetadata() *modmeta.ModMetadata {
	return &modmeta.ModMetadata{
		System:      "canary",
		ID:          d.Name,
		Name:        d.Name,
		Version:     d.Version,
		Description: "",
		URL:         "",
		Authors:     d.Author,
	}
}

func ReadCanaryInf(reader io.Reader) (*PluginDescriptor, error) {
	props := viutils.NewProperties()
	if err := props.Read(reader); err != nil {
		return nil, err
	}

	return &PluginDescriptor{
		Name:         props.GetOrDefault("name", props.GetOrDefault("main-class", "")),
		Version:      props.GetOrDefault("version", "UNKNOWN"),
		Author:       props.GetOrDefault("author", "UNKNOWN"),
		Language:     props.GetOrDefault("language", "java"),
		EnableEarly:  props.GetBoolOrDefault("enable-early", false),
		Dependencies: strings.Split(props.GetOrDefault("dependencies", ""), ","),
	}, nil
}
