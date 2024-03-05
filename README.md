go-modmeta
===

go-modmeta is a Go library for reading metadata on Minecraft mods contained by
a jar file. The following data will be retrieved for each mod in a jar:
- The mod loader / system in use by the mod
- The ID, name, version, description, etc.

go-modmeta will search jars for mods from each supported mod system.

### Supported platforms

- Minecraft Forge (`mcmod.info` and `mods.toml`), identified as `forge`
- Fabric (`fabric.mod.json`), identified as `fabric`
- LiteLoader (`litemod.json`), identified as `liteloader`
- Canary (`Canary.inf`), identified as `canary`

## Usage

```go
import "github.com/jamiemansfield/go-modmeta/modmeta"
```

The most simple example is to fetch the available mod metadata from a jar
file.

```go
mods, err := modmeta.FindMetadata("example.jar")
```

## Licence

This library is distributed under the BSD 2-Clause Licence, found in the
[LICENSE.txt] file.

[LICENSE.txt]: ./LICENSE.txt
