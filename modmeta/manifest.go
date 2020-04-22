package modmeta

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

var (
	FaultyManifest = errors.New("faulty manifest provided")
)

// ReadJarManifest takes a reader for a Jar file's manifest
// (META-INF/MANIFEST.mf), and returns a map of KEY: VALUE.
// If there is a format issue with the manifest, an error
// will be returned.
func ReadJarManifest(r io.Reader) (map[string]string, error) {
	manifest := map[string]string{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		delimeter := strings.Index(line, ":")
		if delimeter == -1 {
			return nil, fmt.Errorf("manifest: faulty line '%s' %w", line, FaultyManifest)
		}

		key := strings.TrimSpace(line[:delimeter])
		value := strings.TrimSpace(line[delimeter+1:])

		manifest[key] = value
	}

	return manifest, nil
}
