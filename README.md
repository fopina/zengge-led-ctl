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

Basic help
```sh
$ zengge-led-ctl -h
CLI controller for Zengge LED devices

Usage:
  zengge-led-ctl [command]

Available Commands:
  scan        List discoverable Zengge LED devices
  connect     Connect to device by MAC address
  power       Power device by MAC address, 1 for ON and 0 for OFF
  color       Set strip color by MAC address, using RGB (0-255)
  version     Display version
  help        Help about any command

Flags:
  -h, --help   help for zengge-led-ctl

Use "zengge-led-ctl [command] --help" for more information about a command.
```

Common flags
- -d, --device string     BLE implementation to use (default "default")
- -w, --duration duration Scanning/connection timeout (default 5s)

Scan for devices
```sh
# scan for 10 seconds (shows duplicates by default)
$ zengge-led-ctl scan -w 10s
Scanning for 10s...
[4455ee88881111bb11aa9090ddddaa88] C -53: Name LEDnetWF02003340898A [MAC: AA:BB:CC:DD:EE:FF OFF Mode: 63 Brightness: 291 RGB: rgb(0, 0, 0) Temperature: 0 LEDs: 15]
...
```

Power on/off
```sh
# using on/off
$ zengge-led-ctl power AA:BB:CC:DD:EE:FF on
# or using 1/0
$ zengge-led-ctl power AA:BB:CC:DD:EE:FF 0
```

Set color (RGB 0-255)
```sh
$ zengge-led-ctl color AA:BB:CC:DD:EE:FF 255 64 0   # orange
$ zengge-led-ctl color AA:BB:CC:DD:EE:FF 0 0 255    # blue
```

Connect (debug/demo)
```sh
# Connects, subscribes to notifications, fetches settings and powers the device off at the end
$ zengge-led-ctl connect AA:BB:CC:DD:EE:FF
```

## Build

Check out [CONTRIBUTING.md](CONTRIBUTING.md)
