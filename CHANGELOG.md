# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], and this project adheres to
[Semantic Versioning].

<!-- references -->
[Keep a Changelog]: https://keepachangelog.com/en/1.0.0/
[Semantic Versioning]: https://semver.org/spec/v2.0.0.html

## [1.0.0] - 2020-12-21

This is the first stable release. There have been no changes to the API since
the `1.0.0-rc.1` release.

## [1.0.0-rc.1] - 2020-11-07

### Added

- Add `StreamWriter`, which logs each line of text in a stream as a log message
- Add `LineWriter`, which logs each call to `Write()` as a log message

### Removed

- **[BC]** Remove `NewWriter()`, use a `StreamWriter` instead
- **[BC]** Remove `NewDebugWriter()`, use a `StreamWriter` with `Demote()` instead

## [1.0.0-rc.0] - 2020-09-10

### Added

- Add `Map`, an in-memory implementation of `Bucket`
- Add `As[Type]()` and `As[Type]Default()` convenience functions

## [0.2.2] - 2020-09-04

### Added

- Add `Prefix()`, which returns a logger that prepends a static prefix to all messages
- Add `Promote()`, which returns a logger that promotes all messages to non-debug level
- Add `Demote()`, which returns a logger that demotes all messages to debug level
- Add `Wrapper` and `Unwrap()` for inspecting loggers that wrap other loggers

## [0.2.1] - 2020-03-08

### Fixed

- `NewWriter()` and `NewDebugWriter()` now split log lines correctly

## [0.2.0] - 2020-03-05

### Added

- Add `Value.Bytes()`

### Changed

- **[BC]** Rename `Value.IsEmpty()` to `IsZero()` to better distinguish between unpopulated values and empty content
- **[BC]** `Value.AsString()` and `AsBytes()` now return an error if the `Value` is the zero-value
- **[BC]** `Value.String()` now panics if the `Value` is the zero-value

## [0.1.2] - 2019-11-20

### Added

- Add `CallbackLogger` (thanks @neetle)

## [0.1.1] - 2019-07-24

### Added

- Add `NewWriter()` and `NewDebugWriter()`

## [0.1.0] - 2019-05-31

- Initial release

<!-- references -->
[Unreleased]: https://github.com/dogmatiq/dogma
[0.1.0]: https://github.com/dogmatiq/dogma/releases/tag/v0.1.0
[0.1.1]: https://github.com/dogmatiq/dogma/releases/tag/v0.1.1
[0.1.2]: https://github.com/dogmatiq/dogma/releases/tag/v0.1.2
[0.2.0]: https://github.com/dogmatiq/dogma/releases/tag/v0.2.0
[0.2.1]: https://github.com/dogmatiq/dogma/releases/tag/v0.2.1
[0.2.2]: https://github.com/dogmatiq/dogma/releases/tag/v0.2.2
[1.0.0-rc.0]: https://github.com/dogmatiq/dogma/releases/tag/v1.0.0-rc.0
[1.0.0-rc.1]: https://github.com/dogmatiq/dogma/releases/tag/v1.0.0-rc.1
[1.0.0]: https://github.com/dogmatiq/dogma/releases/tag/v1.0.0

<!-- version template
## [0.0.1] - YYYY-MM-DD

### Added
### Changed
### Deprecated
### Removed
### Fixed
### Security
-->
