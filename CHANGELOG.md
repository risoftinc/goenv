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

## [1.0.1] - 2025-09-06

### Fixed
- 🐛 **Inline comments in .env files**: Fixed parser to properly handle inline comments in key-value format files
  - Parser now correctly distinguishes between `#` inside quotes vs `#` as comment delimiter
  - Supports both single and double quotes
  - Handles multiple `#` characters in a single line correctly
  - Examples that now work: `CONNECTION=MYSQL  #Mysql, Postgres, Etc`
