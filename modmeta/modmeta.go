// Package modmeta provides functionality to get mod metadata
// from mod binaries.
package modmeta

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
