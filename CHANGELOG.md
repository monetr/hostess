# Change Log

## [0.5.5](https://github.com/monetr/hostess/compare/v0.5.4...v0.5.5) (2026-06-30)


### Bug Fixes

* Bringing up to modern go standards ([fbbfd46](https://github.com/monetr/hostess/commit/fbbfd46a89d0d9adb2657091a9c255602382567f))
* spelling ([017b8f3](https://github.com/monetr/hostess/commit/017b8f33a48b991bb2b8ceb3b8f9a8fd693f181c))
* spelling ([6426b1e](https://github.com/monetr/hostess/commit/6426b1e7c671141a077485d8caebfac63f6257c6))
* **test:** Fixing tests in CI/CD ([2e6a2f8](https://github.com/monetr/hostess/commit/2e6a2f8a09cd4ec1b2e74229ae9008d393438af8))


### Miscellaneous

* **main:** release 0.5.3 ([#2](https://github.com/monetr/hostess/issues/2)) ([81dbe87](https://github.com/monetr/hostess/commit/81dbe87b6b9a3d213fc4455f954069d35a21cfa8))
* **main:** release 0.5.4 ([#4](https://github.com/monetr/hostess/issues/4)) ([ab81dcb](https://github.com/monetr/hostess/commit/ab81dcb51ee0828df8b137ee1c2536723985b8f9))
* Tweaking readme ([7994ad1](https://github.com/monetr/hostess/commit/7994ad148f19d1eeb4b6cb2ea98f1c4f5e027522))
* Updating paths ([0d7b015](https://github.com/monetr/hostess/commit/0d7b015d700b01e61d9cbe2e210e81605b1e71bc))


### Refactor

* Bringing into my own patterns ([513748d](https://github.com/monetr/hostess/commit/513748df568a0a3a74e5d316f8fa6dd7245f6cf9))


### Build Automation

* Upgrading to go 1.25 ([855b48c](https://github.com/monetr/hostess/commit/855b48cb4658ba68ae3e0708d1578ed4842ae2bf))

## [0.5.4](https://github.com/monetr/hostess/compare/v0.5.3...v0.5.4) (2026-06-30)


### Bug Fixes

* Bringing up to modern go standards ([fbbfd46](https://github.com/monetr/hostess/commit/fbbfd46a89d0d9adb2657091a9c255602382567f))

## [0.5.3](https://github.com/monetr/hostess/compare/v0.5.2...v0.5.3) (2026-06-30)


### Bug Fixes

* **test:** Fixing tests in CI/CD ([2e6a2f8](https://github.com/monetr/hostess/commit/2e6a2f8a09cd4ec1b2e74229ae9008d393438af8))


### Miscellaneous

* Tweaking readme ([7994ad1](https://github.com/monetr/hostess/commit/7994ad148f19d1eeb4b6cb2ea98f1c4f5e027522))
* Updating paths ([0d7b015](https://github.com/monetr/hostess/commit/0d7b015d700b01e61d9cbe2e210e81605b1e71bc))


### Refactor

* Bringing into my own patterns ([513748d](https://github.com/monetr/hostess/commit/513748df568a0a3a74e5d316f8fa6dd7245f6cf9))


### Build Automation

* Upgrading to go 1.25 ([855b48c](https://github.com/monetr/hostess/commit/855b48cb4658ba68ae3e0708d1578ed4842ae2bf))

## v0.5.2 (March 13, 2020)

Bug Fixes

- `hostess fmt -n` works properly again, and has more specific behavior:
- `hostess fmt` will replace duplicates without asking for help
- `hostess fmt -n` will *not* replace duplicates, and will exit with error if any are found (#41)
- `hostess fmt` with and without `-n` will exit with error if conflicting hostnames are found because hostess cannot fix the conflicts

## v0.5.1 (March 10, 2020)

Bug Fixes

- Format will no longer exit with an error when encountering a duplicate entry (#39)

## v0.5.0 (March 7, 2020)

Breaking changes

- Windows now has a platform-specific hosts format with one IP and hostname per line

## v0.4.1 (February 28, 2020)

Bug Fixes

- Fix hostfiles not saving on Windows #27

## v0.4.0 (February 28, 2020)

0.4.0 is a major refactor of the frontend (CLI) focused on simplifying the UI
and code, supporting newer Go tooling (i.e. go mod), and removing external
dependencies.

Breaking Changes

- Moved CLI to `github.com/cbednarski/hostess`. `go get` should now do what you probably wanted the first time.
- Moved library to `github.com/cbednarski/hostess/hostess`
- Many command aliases and flags have been removed
- `Hostlist.Enable` and `Hostlist.Disable` now return an `error` instead of `bool`. Check against `ErrHostnameNotFound`.
- Several functions will now return `ErrInvalidVersionArg` instead of panicking in that case

Improvements

- Removed `codegangsta/cli`
- Removed `aff` command
- Removed `del` command (use `rm` instead)
- Removed `list` command (use `ls` instead)
- Removed `fixed` command (just run `fmt`)
- Command `fix` renamed to `fmt`
- Removed `-s` and `-q` flags. Errors are now shown always. Redirect stderr if you don't want to see them.
- Removed `-f` from various commands. Use `fmt` instead.
- Added Go mod support
- Added AppVeyor for Windows builds
- Overhauled the Makefile for easier builds

## v0.3.0 (February 18, 2018)

Improvements

- Added `fixed` subcommand which checks whether the hosts file is already formatted by hostess

Bug Fixes

- Show an error when there is a parsing failure instead of silently truncating the hosts file
- Global flags between hostess and the subcommand are no longer ignored
- Binary should now display the correct version of the software

## v0.2.1 (May 17, 2016)

Bug Fixes

- Fix vendor path for `codegangsta/cli`

## v0.2.0 (May 10, 2016)

Improvements

- Vendor `codegangsta/cli` for more reliable builds

Bug Fixes

- Fix panic in `hostess ls` #14
- Fix incompatible API in CLI library #15

## v0.1.0 (June 6, 2015)

Initial release
