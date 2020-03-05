# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], and this project adheres to
[Semantic Versioning].

<!-- references -->
[Keep a Changelog]: https://keepachangelog.com/en/1.0.0/
[Semantic Versioning]: https://semver.org/spec/v2.0.0.html

## [Unreleased]

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
[0.2.1]: https://github.com/dogmatiq/dogma/releases/tag/v0.1.2

<!-- version template
## [0.0.1] - YYYY-MM-DD

### Added
### Changed
### Deprecated
### Removed
### Fixed
### Security
-->
