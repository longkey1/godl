# godl

godl is golang downloader.


## Usage

```
$ godl [command]
```

## Available Commands

- `completion` generate the autocompletion script for the specified shell
  `goroots` describe goroots directory path
- `help` Help about any command
- `install` install specific version
- `list` installed version list
- `list-remote` downloadable version list
- `remove` remove specific version

Use "godl [command] --help" for more information about a command.

## Configration

`path/to/godl/config.toml`

```
golang_url = "https://golang.org"
goroots_dir = "goroots"
temp_dir" = "tmp"
```

