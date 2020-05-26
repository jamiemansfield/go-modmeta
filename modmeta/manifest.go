package modmeta

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// ReadJarManifest takes a reader for a Jar file's manifest
// (META-INF/MANIFEST.mf), and returns a map of KEY: VALUE.
// If there is a format issue with the manifest, an error
// will be returned.
func ReadJarManifest(r io.Reader) (map[string]string, error) {
	manifest := map[string]string{}

	scanner := bufio.NewScanner(r)
	lineNum := 0
	var lastName string
	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		// Ignore empty lines
		if len(line) == 0 {
			lastName = ""
			continue
		}

		// Continuation line
		if strings.HasPrefix(line, " ") {
			if lastName == "" {
				return nil, fmt.Errorf("manifest[%d]: faulty line '%s'", lineNum, line)
			}
			manifest[lastName] += line[1:]
			continue
		}

		delimeter := strings.Index(line, ":")
		if delimeter == -1 {
			return nil, fmt.Errorf("manifest: faulty line '%s'", line)
		}

		key := strings.TrimSpace(line[:delimeter])
		value := strings.TrimSpace(line[delimeter+1:])

		lastName = key
		manifest[key] = value
	}

	return manifest, nil
}
