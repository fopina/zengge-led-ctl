# zengge-led-ctl

[![goreference](https://pkg.go.dev/badge/github.com/fopina/zengge-led-ctl.svg)](https://pkg.go.dev/github.com/fopina/zengge-led-ctl)
[![release](https://img.shields.io/github/v/release/fopina/zengge-led-ctl)](https://github.com/fopina/zengge-led-ctl/releases)
[![downloads](https://img.shields.io/github/downloads/fopina/zengge-led-ctl/total.svg)](https://github.com/fopina/zengge-led-ctl/releases)
[![ci](https://github.com/fopina/zengge-led-ctl/actions/workflows/publish-main.yml/badge.svg)](https://github.com/fopina/zengge-led-ctl/actions/workflows/publish-main.yml)
[![test](https://github.com/fopina/zengge-led-ctl/actions/workflows/test.yml/badge.svg)](https://github.com/fopina/zengge-led-ctl/actions/workflows/test.yml)
[![codecov](https://codecov.io/github/fopina/zengge-led-ctl/graph/badge.svg)](https://codecov.io/github/fopina/zengge-led-ctl)


CLI controller for Zengge LED devices.

> Only possible thanks to the research done by [@8none1](https://github.com/8none1): [zengge_lednetwf](https://github.com/8none1/zengge_lednetwf/)

## Usage

```sh
➜  golang-template -h
golang project template demo application

Usage:
  golang-template [flags]
  golang-template [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  example     example subcommand which adds or multiplies two given integers
  help        Help about any command
  version     Display version

Flags:
  -h, --help   help for golang-template

Use "golang-template completion [command] --help" for more information about a command.
```

```sh
➜  golang-template example 2 3 --add
5
➜  golang-template example 2 3 --multiply
6
```

## Build

Check out [CONTRIBUTING.md](CONTRIBUTING.md)

### Makefile Targets
```sh
➜  make
bootstrap                      install build deps
build                          build golang binary
clean                          clean up environment
help                           list makefile targets
install                        install golang binary
race                           display test coverage with race
run                            run the app
snapshot                       goreleaser snapshot
test                           display test coverage
```
