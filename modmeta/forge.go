package modmeta

import (
	"errors"
	"io"
	"io/ioutil"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/pelletier/go-toml"
)

var (
	FailedToReadMcModInfoVersion = errors.New("forge: failed to read mcmod.info with any supported format")
)

// Reads a mcmod.info file that uses either the V1 or V2 format. If
// modmeta is unable to read in either format, FailedToReadMcModInfoVersion
// will be returned. The System is set to "forge", though its worth noting
// that other mod systems use FML's loader - for example, Sponge plugins.
func ReadMcModInfo(reader io.Reader) ([]*ModMetadata, error) {
	raw, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	_, jsonType, _, err := jsonparser.Get(raw)
	if err != nil {
		return nil, err
	}

	var listRaw []byte

	// V1 Format
	if jsonType == jsonparser.Array {
		listRaw = raw
	} else
	// V2 Format
	if jsonType == jsonparser.Object {
		modsList, dataType, _, err := jsonparser.Get(raw, "modList")
		if err != nil {
			return nil, err
		}
		if dataType != jsonparser.Array {
			return nil, FailedToReadMcModInfoVersion
		}
		listRaw = modsList
	} else {
		return nil, FailedToReadMcModInfoVersion
	}

	var mods []*ModMetadata
	for _, mod := range getJsonArray(listRaw) {
		mods = append(mods, &ModMetadata{
			System:      "forge",
			ID:          getJsonString(mod, "modid"),
			Name:        getJsonString(mod, "name"),
			Version:     getJsonString(mod, "version"),
			Description: getJsonString(mod, "description"),
			URL:         getJsonString(mod, "url"),
			Authors:     strings.Join(getJsonStringArray(mod, "authorList"), ", "),
		})
	}

	return mods, nil
}

func getJsonString(raw []byte, key string) string {
	value, err := jsonparser.GetString(raw, key)
	if err != nil {
		return ""
	}
	return value
}

func getJsonArray(raw []byte, key ...string) [][]byte {
	var values [][]byte

	arrayRaw, _, _, err := jsonparser.Get(raw, key...)
	if err == nil {
		_, _ = jsonparser.ArrayEach(arrayRaw, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			values = append(values, value)
		})
	}

	return values
}

func getJsonStringArray(raw []byte, key ...string) []string {
	var values []string
	for _, value := range getJsonArray(raw, key...) {
		values = append(values, string(value))
	}
	return values
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
			ID:          getTomlString(modTree, "modId"),
			Name:        getTomlString(modTree, "displayName"),
			Version:     getTomlString(modTree, "version"),
			Description: getTomlString(modTree, "description"),
			URL:         getTomlString(modTree, "displayURL"),
			Authors:     getTomlString(modTree, "authors"),
		})
	}

	return mods, nil
}

func getTomlString(tree *toml.Tree, key string) string {
	if tree.Has(key) {
		return tree.Get(key).(string)
	} else {
		return ""
	}
}
