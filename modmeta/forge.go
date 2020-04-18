package modmeta

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/pelletier/go-toml"
	"io"
	"io/ioutil"
	"strings"
)

var (
	FailedToReadMcModInfoVersion = errors.New("forge: failed to read mcmod.info with any supported format")
)

// Represents a single mods' metadata from Minecraft Forge's
// mcmod.info standard.
type McModInfoMetadata struct {
	ID string `json:"modid"`
	Name string `json:"name"`
	Version string `json:"version"`
	Description string `json:"description"`
	URL string `json:"url"`
	Authors []string `json:"authorList"`
}

type mcModInfoMetadataV1 []*McModInfoMetadata
type mcModInfoMetadataV2 struct {
	Mods []*McModInfoMetadata `json:"modList"`
}

// Reads a mcmod.info file that uses the V1 specification.
func ReadMcModInfoV1(reader io.Reader) ([]*McModInfoMetadata, error) {
	var mods mcModInfoMetadataV1
	err := json.NewDecoder(reader).Decode(&mods)
	if err != nil {
		return nil, err
	}

	return mods, nil
}

// Reads a mcmod.info file that uses the V2 specification.
func ReadMcModInfoV2(reader io.Reader) ([]*McModInfoMetadata, error) {
	var mods mcModInfoMetadataV2
	err := json.NewDecoder(reader).Decode(&mods)
	if err != nil {
		return nil, err
	}

	return mods.Mods, nil
}

// Reads a mcmod.info file that uses either the V1 or V2 format. If
// modmeta is unable to read in either format, FailedToReadMcModInfoVersion
// will be returned.
func ReadMcModInfo(reader io.Reader) ([]*McModInfoMetadata, error) {
	// We need to copy the bytes so we can read it twice
	raw, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	// Try V1 format
	mods, err := ReadMcModInfoV1(bytes.NewReader(raw))
	if err == nil {
		return mods, nil
	}

	// Try V2 format
	mods, err = ReadMcModInfoV2(bytes.NewReader(raw))
	if err == nil {
		return mods, nil
	}

	return nil, FailedToReadMcModInfoVersion
}

// Creates a ModMetadata for the mcmod.info metadata. The System
// is set to "forge", though its worth noting that other mod
// systems use FML's loader - for example, Sponge plugins.
func (m *McModInfoMetadata) ToModMetadata() *ModMetadata {
	return &ModMetadata{
		System:      "forge",
		ID:          m.ID,
		Name:        m.Name,
		Version:     m.Version,
		Description: m.Description,
		URL:         m.URL,
		Authors:     strings.Join(m.Authors, ", "),
	}
}

// Reads a mods.toml file.
func ReadForgeModsToml(reader io.Reader) ([]*ModMetadata, error) {
	tree, err := toml.LoadReader(reader)
	if err != nil {
		return nil, err
	}
	modsTree := tree.Get("mods").([]*toml.Tree)

	var mods []*ModMetadata
	for _, modTree := range modsTree {
		mods = append(mods, &ModMetadata{
			System:      "forge",
			ID:          modTree.Get("modId").(string),
			Name:        modTree.Get("displayName").(string),
			Version:     modTree.Get("version").(string),
			Description: modTree.Get("description").(string),
			URL:         modTree.Get("displayURL").(string),
			Authors:     modTree.Get("authors").(string),
		})
	}

	return mods, nil
}
