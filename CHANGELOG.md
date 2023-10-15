# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.3.1] - 2023-10-15
### Added
* The version command

### Changed
* The target container registry to icr global

### Removed
* Now the test relies only to the standard tool

### Fixed


## [0.3.0] - 2023-10-15
### Added
* Minimal support for GitHub actions

### Changed
* Updates the project base to Go `1.21`
* Better way to deal with the endpoint configuration from the container
* Default log level to debug

### Removed

### Fixed


## [0.2.2] - 2023-02-27
### Added
* Complete commmand line
* Additional label

### Changed
* Improve the log messages when a connection fails

### Removed

### Fixed

  
## [0.2.1] - 2023-02-09
### Added
* Adds a new way to test the endpoints in parallel mode

### Changed
* Reduces the model.Endpoint deleting the kind field

### Removed

### Fixed


## [0.2.0] - 2023-02-09
### Added
* Logging support using [logrus](https://github.com/sirupsen/logrus)

### Changed
* Minor changes to adjust the version generation

### Removed

### Fixed
* The exit from all the commands should be controlled


## [0.1.0] - 2023-02-08
### Added
* Docker support
* helm deployment support

### Changed
* Minor changes to adjust the version generation

### Removed

### Fixed
* The Docker releases having `PROJECT_VERSION`


## [0.0.0] - 2023-02-08
### Added
* Add everything!

### Changed

### Removed

### Fixed


[Unreleased]: https://github.com/jjuarez/simple-prober/compare/0.3.1...HEAD
[0.3.1]: https://github.com/jjuarez/simple-prober/compare/0.3.0...0.3.1
[0.3.0]: https://github.com/jjuarez/simple-prober/compare/0.2.2...0.3.0
[0.2.2]: https://github.com/jjuarez/simple-prober/compare/0.2.1...0.2.2
[0.2.1]: https://github.com/jjuarez/simple-prober/compare/0.2.0...0.2.1
[0.2.0]: https://github.com/jjuarez/simple-prober/compare/0.1.0...0.2.0
[0.1.0]: https://github.com/jjuarez/simple-prober/compare/0.0.0...0.1.0
[0.0.0]: https://github.com/jjuarez/simple-prober/tree/0.0.0
