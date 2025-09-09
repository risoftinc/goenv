# GoEnv - Multi-Format Environment Variable Loader

GoEnv is a Go package that allows you to load environment variables from multiple file formats with nested values support. This package supports JSON and YAML formats in addition to traditional key-value files.

## Features

- ✅ **Multi-format support**: Key-value (.env), JSON (.json), and YAML (.yaml/.yml)
- ✅ **Nested values**: Support for nested data structures with dot notation
- ✅ **Auto-detection**: Automatically detects file format based on extension
- ✅ **Type conversion**: Automatic conversion to various data types (string, int, bool, float64, time.Duration, etc.)
- ✅ **Generic functions**: Uses Go generics for type safety
- ✅ **Convenience functions**: Helper functions for common data types
- ✅ **Easy migration**: Simple API for loading environment variables

## Installation

```bash
go get github.com/risoftinc/goenv
```

## Usage

### 1. Key-Value Format (.env)

```go
package main

import (
    "fmt"
    "github.com/risoftinc/goenv"
)

func main() {
    // Load from .env file
    err := goenv.LoadEnv("config.env")
    if err != nil {
        panic(err)
    }

    // Get values with default values
    appName := goenv.GetEnvString("APP_NAME", "DefaultApp")
    port := goenv.GetEnvInt("PORT", 8080)
    debug := goenv.GetEnvBool("DEBUG", false)

    fmt.Printf("App: %s, Port: %d, Debug: %t\n", appName, port, debug)
}
```

**config.env:**
```env
APP_NAME=MyApp
PORT=8080
DEBUG=true
```

### 2. JSON Format

```go
package main

import (
    "fmt"
    "github.com/risoftinc/goenv"
)

func main() {
    // Load from JSON file
    err := goenv.LoadEnv("config.json")
    if err != nil {
        panic(err)
    }

    // Get nested values with dot notation
    dbHost := goenv.GetEnvString("database.host", "localhost")
    dbPort := goenv.GetEnvInt("database.port", 5432)
    timeout := goenv.GetEnvFloat64("timeout", 30.0)

    fmt.Printf("DB: %s:%d, Timeout: %.1f\n", dbHost, dbPort, timeout)
}
```

**config.json:**
```json
{
  "database": {
    "host": "localhost",
    "port": 5432,
    "name": "mydb"
  },
  "timeout": 30.5,
  "features": ["auth", "logging"]
}
```

### 3. YAML Format

```go
package main

import (
    "fmt"
    "github.com/risoftinc/goenv"
)

func main() {
    // Load from YAML file
    err := goenv.LoadEnv("config.yaml")
    if err != nil {
        panic(err)
    }

    // Get nested values
    apiHost := goenv.GetEnvString("api.host", "localhost")
    apiPort := goenv.GetEnvInt("api.port", 3000)

    fmt.Printf("API: %s:%d\n", apiHost, apiPort)
}
```

**config.yaml:**
```yaml
api:
  host: localhost
  port: 3000
  timeout: 30

database:
  host: localhost
  port: 5432
  name: mydb

features:
  - auth
  - logging
  - metrics
```

### 4. Using Specific Format

```go
// Load with explicitly specified format
err := goenv.LoadEnvWithFormat(goenv.FormatJSON, "config.json")
err := goenv.LoadEnvWithFormat(goenv.FormatYAML, "config.yaml")
err := goenv.LoadEnvWithFormat(goenv.FormatKeyValue, "config.env")
```

### 5. Generic Functions

```go
// Using generic function for type safety
appName := goenv.GetEnv("APP_NAME", "DefaultApp")        // string
port := goenv.GetEnv("PORT", 8080)                       // int
debug := goenv.GetEnv("DEBUG", false)                    // bool
timeout := goenv.GetEnv("TIMEOUT", 30.5)                 // float64
duration := goenv.GetEnv("CACHE_DURATION", 5*time.Minute) // time.Duration
```

### 6. Duration Support

```go
package main

import (
    "fmt"
    "time"
    "github.com/risoftinc/goenv"
)

func main() {
    err := goenv.LoadEnv("config.env")
    if err != nil {
        panic(err)
    }

    // Get duration values
    timeout := goenv.GetEnvDuration("TIMEOUT", 30*time.Second)
    retryInterval := goenv.GetEnvDuration("RETRY_INTERVAL", 5*time.Minute)
    cleanupInterval := goenv.GetEnvDuration("CLEANUP_INTERVAL", 1*time.Hour)

    fmt.Printf("Timeout: %v\n", timeout)
    fmt.Printf("Retry Interval: %v\n", retryInterval)
    fmt.Printf("Cleanup Interval: %v\n", cleanupInterval)
}
```

**config.env:**
```env
TIMEOUT=45s
RETRY_INTERVAL=2m30s
CLEANUP_INTERVAL=6h
```

### 7. Nested Keys with GetEnvNested

```go
// GetEnvNested converts dot notation to UPPER_CASE with underscore
// "db.host" -> "DB_HOST"
dbHost := goenv.GetEnvNested("db.host", "localhost")
dbPort := goenv.GetEnvNested("db.port", 5432)
```

## API Reference

### Functions

#### LoadEnv
```go
func LoadEnv(file ...string) error
```
Loads environment variables from files with auto-detection format.

#### LoadEnvWithFormat
```go
func LoadEnvWithFormat(format FileFormat, file ...string) error
```
Loads environment variables with specified format.

#### GetEnv (Generic)
```go
func GetEnv[T any](key string, defaultVal T) T
```
Retrieves environment variable with automatic type conversion.

#### GetEnvNested
```go
func GetEnvNested[T any](key string, defaultVal T) T
```
Retrieves environment variable with dot notation (converts to UPPER_CASE).

#### Convenience Functions
```go
func GetEnvString(key string, defaultVal string) string
func GetEnvInt(key string, defaultVal int) int
func GetEnvBool(key string, defaultVal bool) bool
func GetEnvFloat64(key string, defaultVal float64) float64
func GetEnvDuration(key string, defaultVal time.Duration) time.Duration
```

### Types

#### FileFormat
```go
type FileFormat int

const (
    FormatAuto FileFormat = iota // Auto-detect based on extension
    FormatKeyValue               // .env format
    FormatJSON                   // .json format
    FormatYAML                   // .yaml/.yml format
)
```

## Nested Values

This package supports nested values with dot notation. Example:

**JSON/YAML:**
```json
{
  "database": {
    "host": "localhost",
    "port": 5432
  }
}
```

**Environment Variables created:**
- `database.host` = "localhost"
- `database.port` = "5432"

**Usage:**
```go
host := goenv.GetEnvString("database.host", "localhost")
port := goenv.GetEnvInt("database.port", 5432)
```

## Basic Usage

Here's a simple example of how to use GoEnv:

**Traditional approach:**
```go
import "os"

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
}
```

**With GoEnv:**
```go
import "github.com/risoftinc/goenv"

func main() {
    err := goenv.LoadEnv("config.env")
    if err != nil {
        panic(err)
    }
    
    port := goenv.GetEnvInt("PORT", 8080)
}
```

## Testing

```bash
go test -v
```

## Complete Examples

See the `example/` folder for complete usage examples with various file formats.

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for a detailed list of changes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
