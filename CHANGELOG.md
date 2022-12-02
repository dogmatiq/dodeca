# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], and this project adheres to
[Semantic Versioning].

<!-- references -->

[keep a changelog]: https://keepachangelog.com/en/1.0.0/
[semantic versioning]: https://semver.org/spec/v2.0.0.html

## [1.4.2] - 2022-12-02

### Changed

- Improved the mechanism used by to `Zap()` detect whether debug-level logging is enabled

## [1.4.1] - 2022-11-29

This release simply removes the package-level deprecation notices, as they are
caught by aggressive linter configurations. The deprecation notice in `go.mod`
has been retained. We will re-add these deprecation notices when we are closer
to ceasing support for this module altogether.

## [1.4.0] - 2022-11-29

> **The Dodeca project is deprecated as of this version.**
>
> - All Dogmatiq projects that use the Dodeca `logging` package are transitioning
>   to use [`go.uber.org/zap`](https://github.com/uber-go/zap) instead.
> - The Dodeca `config` package has been superceded by
>   [`dogmatiq/ferrite`](https://github.com/dogmatiq/ferrite).
>
> This repository will be archived once it is no longer used by other Dogmatiq
> projects.

### Added

- Add `logging.Zap()` to adapt an Uber Zap logger to a `logging.Logger`

## [1.3.1] - 2022-01-11

### Fixed

- `logging.StreamWriter` now correctly handles Windows newlines (CRLF)
- `logging.StreamWriter` now treats a single CR character as a newline

## [1.3.0] - 2022-01-02

### Added

- Add `logging.Tee()`, which returns a logger that dispatches all messages to multiple other loggers

## [1.2.0] - 2021-08-26

### Added

- Add `config.NotDefined`, `InvalidValue` and `InvalidDefaultValue` errors

### Changed

- Use new `config` error types as panic values instead of arbitrary strings

## [1.1.0] - 2021-07-07

### Added

- Add `config.AsURL()` and `AsURLDefault()`

## [1.0.0] - 2020-12-21

This is the first stable release. There have been no changes to the API since
the `1.0.0-rc.1` release.

## [1.0.0-rc.1] - 2020-11-07

### Added

- Add `logging.StreamWriter`, which logs each line of text in a stream as a log message
- Add `logging.LineWriter`, which logs each call to `Write()` as a log message

### Removed

- **[BC]** Remove `logging.NewWriter()`, use a `StreamWriter` instead
- **[BC]** Remove `logging.NewDebugWriter()`, use a `StreamWriter` with `Demote()` instead

## [1.0.0-rc.0] - 2020-09-10

### Added

- Add `config.Map`, an in-memory implementation of `Bucket`
- Add `config.As[Type]()` and `As[Type]Default()` convenience functions

## [0.2.2] - 2020-09-04

### Added

- Add `logging.Prefix()`, which returns a logger that prepends a static prefix to all messages
- Add `logging.Promote()`, which returns a logger that promotes all messages to non-debug level
- Add `logging.Demote()`, which returns a logger that demotes all messages to debug level
- Add `logging.Wrapper` and `Unwrap()` for inspecting loggers that wrap other loggers

## [0.2.1] - 2020-03-08

### Fixed

- `logging.NewWriter()` and `NewDebugWriter()` now split log lines correctly

## [0.2.0] - 2020-03-05

### Added

- Add `config.Value.Bytes()`

### Changed

- **[BC]** Rename `config.Value.IsEmpty()` to `IsZero()` to better distinguish between unpopulated values and empty content
- **[BC]** `config.Value.AsString()` and `AsBytes()` now return an error if the `Value` is the zero-value
- **[BC]** `config.Value.String()` now panics if the `Value` is the zero-value

## [0.1.2] - 2019-11-20

### Added

- Add `logging.CallbackLogger` (thanks @neetle)

## [0.1.1] - 2019-07-24

### Added

- Add `logging.NewWriter()` and `NewDebugWriter()`

## [0.1.0] - 2019-05-31

- Initial release

<!-- references -->

[unreleased]: https://github.com/dogmatiq/dodeca
[0.1.0]: https://github.com/dogmatiq/dodeca/releases/tag/v0.1.0
[0.1.1]: https://github.com/dogmatiq/dodeca/releases/tag/v0.1.1
[0.1.2]: https://github.com/dogmatiq/dodeca/releases/tag/v0.1.2
[0.2.0]: https://github.com/dogmatiq/dodeca/releases/tag/v0.2.0
[0.2.1]: https://github.com/dogmatiq/dodeca/releases/tag/v0.2.1
[0.2.2]: https://github.com/dogmatiq/dodeca/releases/tag/v0.2.2
[1.0.0-rc.0]: https://github.com/dogmatiq/dodeca/releases/tag/v1.0.0-rc.0
[1.0.0-rc.1]: https://github.com/dogmatiq/dodeca/releases/tag/v1.0.0-rc.1
[1.0.0]: https://github.com/dogmatiq/dodeca/releases/tag/v1.0.0
[1.1.0]: https://github.com/dogmatiq/dodeca/releases/tag/v1.1.0
[1.2.0]: https://github.com/dogmatiq/dodeca/releases/tag/v1.2.0
[1.3.0]: https://github.com/dogmatiq/dodeca/releases/tag/v1.3.0
[1.3.1]: https://github.com/dogmatiq/dodeca/releases/tag/v1.3.1
[1.4.0]: https://github.com/dogmatiq/dodeca/releases/tag/v1.4.0
[1.4.1]: https://github.com/dogmatiq/dodeca/releases/tag/v1.4.1

<!-- version template
## [0.0.1] - YYYY-MM-DD

### Added
### Changed
### Deprecated
### Removed
### Fixed
### Security
-->
