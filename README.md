# hostess [![GoDoc](https://godoc.org/github.com/monetr/hostess?status.svg)](http://godoc.org/github.com/monetr/hostess)

This is a fork of https://github.com/cbednarski/hostess

An **idempotent** command-line utility for managing your `/etc/hosts`* file.

    hostess add local.example.com 127.0.0.1
    hostess add staging.example.com 10.0.2.16

Why? Because you edit `/etc/hosts` for development, testing, and debugging.
Because sometimes DNS doesn't work in production. And because editing
`/etc/hosts` by hand is a pain. Put hostess in your `Makefile` or deploy scripts
and call it a day.

\* And `C:\Windows\System32\drivers\etc\hosts` on Windows.

**Note: 0.4.0 has backwards incompatible changes in the API and CLI.** See
`CHANGELOG.md` for details.

## Installation

hostess is distributed as pre-built binaries attached to each [GitHub
release](https://github.com/monetr/hostess/releases). They are statically linked
(CGO disabled), so there are no runtime dependencies, and each one is published
alongside a `.sha256` checksum and a [Sigstore](https://www.sigstore.dev/)
build-provenance attestation you can verify.

Binaries are built for `linux-amd64`, `linux-arm64`, `linux-arm`, `darwin-amd64`,
`darwin-arm64`, `windows-amd64`, and `windows-arm64`. Each asset is named
`hostess-<version>-<os>-<arch>`.

### macOS and Linux

With the [GitHub CLI](https://cli.github.com) you can grab the newest release for
your platform in one step:

    PLATFORM="linux-amd64" # or linux-arm64, linux-arm, darwin-amd64, darwin-arm64
    gh release download --repo monetr/hostess --pattern "hostess-*-${PLATFORM}"
    chmod +x hostess-*-"${PLATFORM}"
    sudo mv hostess-*-"${PLATFORM}" /usr/local/bin/hostess

Without it, download straight from the [latest
release](https://github.com/monetr/hostess/releases/latest) (fill in the version
shown there):

    VERSION="v0.5.3" # the latest tag from the releases page
    PLATFORM="linux-amd64"
    curl -fsSL -o hostess \
      "https://github.com/monetr/hostess/releases/download/${VERSION}/hostess-${VERSION}-${PLATFORM}"
    chmod +x hostess
    sudo mv hostess /usr/local/bin/hostess

### Windows

Download `hostess-<version>-windows-amd64.exe` (or `windows-arm64`) from the
[latest release](https://github.com/monetr/hostess/releases/latest), or use the
[GitHub CLI](https://cli.github.com):

    gh release download --repo monetr/hostess --pattern "hostess-*-windows-amd64.exe"

Rename it to `hostess.exe` and put it somewhere on your `PATH`. Remember the hosts
file is protected, so run it from an elevated (_Run as administrator_) prompt.

### Verifying a download

Every release binary ships with a matching `.sha256` checksum and a
build-provenance attestation, so you can confirm a download is intact and was
actually built by this repo's release workflow:

    # Checksum: download the .sha256 asset next to the binary, then
    sha256sum -c hostess-v0.5.3-linux-amd64.sha256

    # Provenance, using the GitHub CLI:
    gh attestation verify hostess-v0.5.3-linux-amd64 --repo monetr/hostess

### Build from source

Build from source with a [recent version of Go](https://golang.org/dl):

    git clone https://github.com/monetr/hostess
    cd hostess
    make install

## Usage

Run `hostess` or `hostess -h` to see a full list of commands.

**Note** The hosts file is protected. On unixes you will need to use `sudo` or
run the `hostess` command as root. On Windows, you will need to run `hostess`
from an elevated prompt (right click and _Run as administrator_).

## Format

On unixes, hostess follows the format specified by `man hosts`, with one line
per IP address:

    127.0.0.1 localhost hostname2 hostname3
    127.0.1.1 machine.name
    # 10.10.20.30 some.host

On Windows, hostess writes each hostname on its own line.

    127.0.0.1 localhost
    127.0.0.1 hostname2
    127.0.0.1 hostname3

## Configuration

hostess may be configured via environment variables.

- `HOSTESS_FMT` may be set to `windows` or `unix` to override platform detection
  for the hosts file format. See Behavior, above, for details

- `HOSTESS_PATH` may be set to override platform detection for the location of
  the hosts file. By default this is `C:\Windows\System32\drivers\etc\hosts` on
  Windows and `/etc/hosts` everywhere else.

## IPv4 and IPv6

It's possible for your hosts file to include overlapping entries for IPv4 and
IPv6. This is an uncommon case so the CLI ignores this distinction. The hostess
library includes logic that differentiates between these cases.

## Contributing

I hope my software is useful, readable, fun to use, and helps you learn
something new. I maintain this software in my spare time. I rarely merge PRs
because I am both lazy and a snob. Bug reports, fixes, questions, and comments
are welcome but expect a delayed response. No refunds. 👻
