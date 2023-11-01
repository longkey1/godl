# godl

godl is The golang downloader.


## Usage

```
$ godl [command]
```

## Available Commands

- completion  Generate the autocompletion script for the specified shell
- help        Help about any command
- install     Install specific version
- list        Installed version list
- list-remote Downloadable version list
- path        Describe path
- remove      Remove specific version

Use "godl [command] --help" for more information about a command.

## Configration

`path/to/godl/config.toml`

```
golang_url = "https://golang.org"
goroots_dir = "/path/to/godl/goroots"
temp_dir" = "/path/to/godl/tmp"
versions = [
  "1.21",
  "1.20",
  "1.19",
  "1.18",
  "1.17",
  "1.16",
]
```
