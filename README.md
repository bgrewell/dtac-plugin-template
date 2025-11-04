# {{plugin}} – DTAC Plugin

A template for building DTAC plugins.

## Quickstart
```bash
./scripts/rename.sh github.com/you/{{repo}} {{plugin}}
go mod tidy
make build
./bin/{{plugin}}
