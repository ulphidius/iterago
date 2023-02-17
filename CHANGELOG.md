# Iterago Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.0] - 2023-02-17

### Added

- Find function
- Fold function

### Changed

- Refactor Filter to stand alone function
- Refactor Map to stand alone function
- Refactor Reduce to stand alone function
- Unwrap an Option now panic

### Removed

- Iterator structure system
- Iterator helpers

## [0.1.0] - 2023-01-29

### Added

- Basic Option type for optional value
- Iterator system
- Filter iterator system
- Mapper function system
- Reduce function system
- Collect function function
- Documentation
- CI
- Generic Slice into Iterator function

[unreleased]: https://github.com/ulphidius/iterago/compare/v0.2.0...master
[0.2.0]:  https://github.com/ulphidius/iterago/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/ulphidius/iterago/releases/tag/v0.1.0
