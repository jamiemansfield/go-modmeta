package modmeta

import (
	"encoding/json"
	"io"
	"strings"
)

// Represents a single mods' metadata from Fabric's fabric.mod.json
// standard.
type FabricModJsonMetadata struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Version string `json:"version"`
	Description string `json:"description"`
	Authors []string `json:"authors"`
	Contact struct {
		Homepage string `json:"homepage"`
	} `json:"contact"`
}

// Reads a fabric.mod.json file.
func ReadFabricModJson(reader io.Reader) (*FabricModJsonMetadata, error) {
	var mod FabricModJsonMetadata
	err := json.NewDecoder(reader).Decode(&mod)
	if err != nil {
		return nil, err
	}

	return &mod, nil
}

// Creates a ModMetadata for the fabric.mod.json metadata. The
// System will be set to "fabric".
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
