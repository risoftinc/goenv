# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

# [Unreleased]
### Added
- 

### Features
- 

## [1.0.0] - 2025-09-04

### Added
- ✨ **Multi-format support**: Added support for key-value (.env), JSON (.json), and YAML (.yaml/.yml) files
- ✨ **Nested values**: Support for nested data structures with dot notation (e.g., `database.host`)
- ✨ **Auto-detection**: Automatic file format detection based on file extension
- ✨ **Type conversion**: Automatic conversion to various data types (string, int, bool, float64, uint64, int64)
- ✨ **Generic functions**: Type-safe generic functions using Go generics
- ✨ **Convenience functions**: Helper functions for common data types (GetEnvString, GetEnvInt, etc.)
- ✨ **Nested key support**: GetEnvNested function for dot notation with automatic case conversion

### API
- `LoadEnv(file ...string) error` - Load environment variables with auto-detection
- `LoadEnvWithFormat(format FileFormat, file ...string) error` - Load with specific format
- `GetEnv[T any](key string, defaultVal T) T` - Generic environment variable getter
- `GetEnvNested[T any](key string, defaultVal T) T` - Nested key getter with dot notation
- `GetEnvString(key string, defaultVal string) string` - String convenience function
- `GetEnvInt(key string, defaultVal int) int` - Integer convenience function
- `GetEnvBool(key string, defaultVal bool) bool` - Boolean convenience function
- `GetEnvFloat64(key string, defaultVal float64) float64` - Float64 convenience function

### File Format Support
- **Key-Value (.env)**: Traditional environment file format with comment support
- **JSON (.json)**: JSON files with nested object flattening
- **YAML (.yaml/.yml)**: YAML files with nested object flattening

### Examples
- Complete example files for all supported formats
- Comprehensive test suite with 100% test coverage
- Documentation with usage examples and API reference
