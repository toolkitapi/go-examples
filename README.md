# ToolkitAPI — Go Examples

[![Go Reference](https://pkg.go.dev/badge/github.com/toolkitapi/go-sdk.svg)](https://pkg.go.dev/github.com/toolkitapi/go-sdk)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Runnable Go examples for every [ToolkitAPI.io](https://toolkitapi.io) endpoint, using the official [`github.com/toolkitapi/go-sdk`](https://pkg.go.dev/github.com/toolkitapi/go-sdk).

## Prerequisites

- Go **1.21+**

## Setup

```bash
go mod download
export TOOLKITAPI_KEY=tk_live_...
```

## Running an example

Each file uses the `//go:build ignore` tag so they can be run individually with `go run` without being picked up by the build system:

```bash
go run dns/lookup.go
go run email/validate_email.go
go run auth/jwt_generate.go
```

## Toolkits

| Toolkit | Folder |
|---------|--------|
| Authentication & security | [`auth/`](auth/) |
| Barcode & QR codes | [`barcode/`](barcode/) |
| File & data conversion | [`convert/`](convert/) |
| Developer utilities | [`devtools/`](devtools/) |
| DNS & domain tools | [`dns/`](dns/) |
| Email validation | [`email/`](email/) |
| Geo & IP | [`geo/`](geo/) |
| Image processing | [`image/`](image/) |
| Media extraction | [`media/`](media/) |
| PDF tools | [`pdf/`](pdf/) |
| Web scraping | [`scrape/`](scrape/) |
| Text analysis (AI) | [`textanalysis/`](textanalysis/) |
| Webhooks | [`webhook/`](webhook/) |

## Links

- [ToolkitAPI docs](https://toolkitapi.io/docs)
- [Go SDK on pkg.go.dev](https://pkg.go.dev/github.com/toolkitapi/go-sdk)
- [bash examples](https://github.com/toolkitapi/bash-examples)
- [Node.js examples](https://github.com/toolkitapi/node-examples)
- [Python examples](https://github.com/toolkitapi/python-examples)
