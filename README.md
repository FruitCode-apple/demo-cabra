# demo-cabra

A small but complete CLI application written in Go that demonstrates the
[`github.com/spf13/cobra`](https://github.com/spf13/cobra) library – packaged
in Debian/Ubuntu as **`golang-github-spf13-cobra-dev`**.

Use this project as a starting point for developing Go CLI tools in a Debian
environment without relying on internet access at build time.

---

## Features

| Subcommand | Description |
|------------|-------------|
| `greet`    | Print a personalised greeting. Supports `--name` and `--loud` flags. |
| `version`  | Print the application version, Go runtime version, and OS/architecture. |

---

## Prerequisites (Debian / Ubuntu)

Install the Cobra development package and the Go toolchain:

```bash
sudo apt-get update
sudo apt-get install -y golang-go golang-github-spf13-cobra-dev
```

> **Tip:** `golang-github-spf13-cobra-dev` installs the Cobra source under
> `/usr/share/gocode/src/github.com/spf13/cobra/`.  When you set
> `GOPATH=/usr/share/gocode` the Go toolchain picks up those offline sources
> instead of downloading from the internet.

---

## Build & Run (standard `go build`)

```bash
# Clone the repository
git clone https://github.com/FruitCode-apple/demo-cabra.git
cd demo-cabra

# Fetch Go module dependencies (internet required for the first run)
go mod tidy

# Build
go build -o demo-cabra .

# Run the binary
./demo-cabra --help
./demo-cabra greet
./demo-cabra greet --name Alice
./demo-cabra greet --name Bob --loud
./demo-cabra version
```

### Build with offline Debian sources

```bash
# Point GOPATH at the system-wide gocode tree provided by Debian packages
export GOPATH=/usr/share/gocode:$GOPATH

# Vendor the Debian-installed sources into the project
go mod vendor

# Build using the vendored sources (no network required)
go build -mod=vendor -o demo-cabra .
```

### Embed the version at build time

```bash
go build -ldflags "-X github.com/FruitCode-apple/demo-cabra/cmd.Version=0.1.0" \
    -o demo-cabra .
./demo-cabra version
```

---

## Build a Debian package

The `debian/` directory contains standard Debian packaging metadata.
`dh-golang` takes care of calling `go build` and installing the binary.

```bash
# Install packaging tools
sudo apt-get install -y devscripts debhelper dh-golang

# Build the .deb package (run from the project root)
dpkg-buildpackage -us -uc -b

# Install the resulting package
sudo dpkg -i ../demo-cabra_0.1.0-1_amd64.deb

# The binary is now available system-wide
demo-cabra --help
```

---

## Project layout

```
demo-cabra/
├── main.go            # Entry point – calls cmd.Execute()
├── go.mod             # Go module definition
├── cmd/
│   ├── root.go        # Root Cobra command & Execute()
│   ├── greet.go       # "greet" subcommand with --name and --loud flags
│   └── version.go     # "version" subcommand
└── debian/
    ├── changelog      # Debian changelog
    ├── compat         # debhelper compatibility level
    ├── control        # Package metadata & build dependencies
    ├── copyright      # License information
    ├── rules          # Debian build rules (uses dh-golang)
    └── source/
        └── format     # Debian source format
```

---

## Example output

```
$ ./demo-cabra --help
demo-cabra is a small demonstration CLI application built with
the github.com/spf13/cobra library …

Usage:
  demo-cabra [command]

Available Commands:
  greet       Print a greeting message
  help        Help about any command
  version     Print the version of demo-cabra

$ ./demo-cabra greet --name Debian --loud
HELLO, Debian!!!

$ ./demo-cabra version
demo-cabra version : dev
Go version         : go1.21.x
OS/Arch            : linux/amd64
```

---

## License

Apache-2.0 – see [debian/copyright](debian/copyright).
