# Iterago Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.6.1] - 2023-03-22

### Fixed

- Remove debug which flood stdout

## [0.6.0] - 2023-03-22

### Added

- Multithreading for many functions. The number of go routine spawned by iterago function is defined by the environment variable **ITERAGO_THREADS**

## [0.5.0] - 2023-03-05

### Added

- Foreach function to execute code without return
- FilterMap function which merges the behavior of filter and map
- FilterReduce function which merges the behavior of filter and reduce
- FilterFold function which merges the behavior of filter and fold
- MapReduce function which merges the behavior of map and reduce
- PartitionForeach function which merges the behavior of partition and foreach

## [0.4.0] - 2023-02-19

### Added

- Sort function. Use Merge Sort algorithm
- IsSorted function
- Enumerate function
- Reverse function
- Chunks function. To split an array into sub arrays
- Any function. To check if all values are invalid
- All function. To check if all values are valid

## [0.3.0] - 2023-02-18

### Added

- Partition function
- Zip function
- Examples for each functions

## [0.2.0] - 2023-02-17

### Added

- Find function
- Fold function

### Changed

- Refactor Filter into stand alone function
- Refactor Map into stand alone function
- Refactor Reduce into stand alone function
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

[unreleased]: https://github.com/ulphidius/iterago/compare/v0.6.1...master
[0.6.1]:  https://github.com/ulphidius/iterago/compare/v0.6.0...v0.6.1
[0.6.0]:  https://github.com/ulphidius/iterago/compare/v0.5.0...v0.6.0
[0.5.0]:  https://github.com/ulphidius/iterago/compare/v0.4.0...v0.5.0
[0.4.0]:  https://github.com/ulphidius/iterago/compare/v0.3.0...v0.4.0
[0.3.0]:  https://github.com/ulphidius/iterago/compare/v0.2.0...v0.3.0
[0.2.0]:  https://github.com/ulphidius/iterago/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/ulphidius/iterago/releases/tag/v0.1.0
