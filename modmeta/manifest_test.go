package modmeta

import (
	"strings"
	"testing"
)

const (
	testManifest = `Manifest-Version: 1.0
Implementation-Title: Example Artifact

Implementation-Version: 1.0.0
`
)

func TestReadJarManifest(t *testing.T) {
	manifest, err := ReadJarManifest(strings.NewReader(testManifest))
	if err != nil {
		t.Error(err)
		return
	}

	if manifest["Manifest-Version"] != "1.0" {
		t.Errorf("Manifest-Version should equal '1.0', not '%s'.", manifest["Manifest-Version"])
		return
	}
	if manifest["Implementation-Title"] != "Example Artifact" {
		t.Errorf("Implementation-Title should equal 'Example Artifact', not '%s'.", manifest["Implementation-Title"])
		return
	}
	if manifest["Implementation-Version"] != "1.0.0" {
		t.Errorf("Implementation-Version should equal '1.0.0', not '%s'.", manifest["Implementation-Version"])
		return
	}
}
