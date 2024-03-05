/*
 * Copyright (c) 2024, Jamie Mansfield <jmansfield@cadixdev.org>
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package viutils

import (
	"bufio"
	"io"
	"strings"
)

type Properties struct {
	Props          map[string]string
	Comments       map[string][]string
	InlineComments map[string]string
	Header         []string
	Footer         []string
}

func NewProperties() *Properties {
	return &Properties{
		Props:          make(map[string]string),
		Comments:       make(map[string][]string),
		InlineComments: make(map[string]string),
		Header:         make([]string, 0),
		Footer:         make([]string, 0),
	}
}

func (p *Properties) Read(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)

	propertyComments := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, ";#") {
			p.Header = append(p.Header, line)
		} else if strings.HasPrefix(line, "#;") {
			p.Footer = append(p.Footer, line)
		} else if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			propertyComments = append(propertyComments, line)
		} else {
			// FIXME: does this behave the same as viutils?
			parts := strings.SplitN(line, "=", 2)
			key := strings.Trim(parts[0], " ")

			if len(parts) > 1 {
				valueParts := strings.SplitN(parts[1], "#!", 2)

				// value
				value := valueParts[0]
				value = strings.ReplaceAll(value, "\\#\\!", "#!")
				p.Props[key] = value

				// comments
				if len(valueParts) > 1 {
					p.InlineComments[key] = valueParts[1]
				}
			} else if strings.ContainsRune(parts[0], '=') {
				// empty value (key=)
				p.Props[key] = ""
			} else {
				// invalid (key)
				propertyComments = make([]string, 0)
				continue
			}

			if len(propertyComments) > 0 {
				p.Comments[key] = propertyComments
				propertyComments = make([]string, 0)
			}
		}
	}

	return scanner.Err()
}

func (p *Properties) GetOrDefault(key string, def string) string {
	value, ok := p.Props[key]
	if !ok {
		return def
	}
	return value
}
