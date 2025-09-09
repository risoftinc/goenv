# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

# [Unreleased]
### Added
- 

### Fixed
- 

### Features
- 

## [1.1.0] - 2025-09-09

### Added
- ✨ **Duration support**: Added `time.Duration` type conversion support
  - Generic function: `goenv.GetEnv("TIMEOUT", 30*time.Second)`
  - Convenience function: `goenv.GetEnvDuration("TIMEOUT", 30*time.Second)`
  - Supports all Go duration formats: `30s`, `5m`, `2h`, `1h30m45s`, etc.