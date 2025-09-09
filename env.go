package goenv

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// FileFormat represents the supported file formats
type FileFormat int

const (
	FormatAuto     FileFormat = iota // Auto-detect based on file extension
	FormatKeyValue                   // .env format
	FormatJSON                       // .json format
	FormatYAML                       // .yaml/.yml format
)

// LoadEnv loads environment variables from files with support for multiple formats
func LoadEnv(file ...string) error {
	return LoadEnvWithFormat(FormatAuto, file...)
}

// LoadEnvWithFormat loads environment variables from files with specified format
func LoadEnvWithFormat(format FileFormat, file ...string) error {
	for _, f := range file {
		if f == "" {
			continue
		}

		// Determine format if auto-detect
		detectedFormat := format
		if format == FormatAuto {
			detectedFormat = detectFormat(f)
		}

		var err error
		switch detectedFormat {
		case FormatKeyValue:
			err = loadKeyValueFile(f)
		case FormatJSON:
			err = loadJSONFile(f)
		case FormatYAML:
			err = loadYAMLFile(f)
		default:
			err = fmt.Errorf("unsupported file format for %s", f)
		}

		if err == nil {
			return nil
		}
	}
	return fmt.Errorf("failed to load any of the specified files")
}

// detectFormat detects file format based on extension
func detectFormat(filename string) FileFormat {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".env":
		return FormatKeyValue
	case ".json":
		return FormatJSON
	case ".yaml", ".yml":
		return FormatYAML
	default:
		// Default to key-value format for unknown extensions
		return FormatKeyValue
	}
}

// loadKeyValueFile loads environment variables from key-value format (.env)
func loadKeyValueFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and full-line comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Find the first # that's not inside quotes
		commentIndex := -1
		inQuotes := false
		quoteChar := byte(0)

		for i := 0; i < len(line); i++ {
			char := line[i]
			if !inQuotes && char == '#' {
				commentIndex = i
				break
			}
			if !inQuotes && (char == '"' || char == '\'') {
				inQuotes = true
				quoteChar = char
			} else if inQuotes && char == quoteChar {
				inQuotes = false
				quoteChar = 0
			}
		}

		// Remove comment if found
		if commentIndex != -1 {
			line = strings.TrimSpace(line[:commentIndex])
		}

		// Skip if line becomes empty after removing comment
		if line == "" {
			continue
		}

		// Parse key=value pairs
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remove quotes if present
		if len(value) >= 2 && ((value[0] == '"' && value[len(value)-1] == '"') ||
			(value[0] == '\'' && value[len(value)-1] == '\'')) {
			value = value[1 : len(value)-1]
		}

		os.Setenv(key, value)
	}

	return scanner.Err()
}

// loadJSONFile loads environment variables from JSON format
func loadJSONFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return err
	}

	// Flatten nested JSON and set environment variables
	flattenAndSetEnv("", jsonData)
	return nil
}

// loadYAMLFile loads environment variables from YAML format
func loadYAMLFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var yamlData map[string]interface{}
	if err := yaml.Unmarshal(data, &yamlData); err != nil {
		return err
	}

	// Flatten nested YAML and set environment variables
	flattenAndSetEnv("", yamlData)
	return nil
}

// flattenAndSetEnv recursively flattens nested maps and sets environment variables
func flattenAndSetEnv(prefix string, data map[string]interface{}) {
	for key, value := range data {
		envKey := key
		if prefix != "" {
			envKey = prefix + "." + key
		}

		switch v := value.(type) {
		case map[string]interface{}:
			// Recursively handle nested objects
			flattenAndSetEnv(envKey, v)
		case []interface{}:
			// Handle arrays by converting to JSON string
			if jsonBytes, err := json.Marshal(v); err == nil {
				os.Setenv(envKey, string(jsonBytes))
			}
		default:
			// Convert other types to string
			os.Setenv(envKey, fmt.Sprintf("%v", v))
		}
	}
}

// GetEnv retrieves environment variable with type conversion and nested key support
func GetEnv[T any](key string, defaultVal T) T {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}

	switch any(defaultVal).(type) {
	case int:
		parsed, err := strconv.Atoi(val)
		if err != nil {
			return defaultVal
		}
		return any(parsed).(T)
	case uint64:
		parsed, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			return defaultVal
		}
		return any(parsed).(T)
	case int64:
		parsed, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return defaultVal
		}
		return any(parsed).(T)
	case bool:
		parsed, err := strconv.ParseBool(val)
		if err != nil {
			return defaultVal
		}
		return any(parsed).(T)
	case float64:
		parsed, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return defaultVal
		}
		return any(parsed).(T)
	case time.Duration:
		parsed, err := time.ParseDuration(val)
		if err != nil {
			return defaultVal
		}
		return any(parsed).(T)
	case string:
		return any(val).(T)
	default:
		return defaultVal
	}
}

// GetEnvNested retrieves nested environment variable using dot notation
// Example: GetEnvNested("db.host", "localhost") will look for DB_HOST environment variable
func GetEnvNested[T any](key string, defaultVal T) T {
	// Convert dot notation to uppercase with underscores
	envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
	return GetEnv(envKey, defaultVal)
}

// GetEnvString is a convenience function for getting string values
func GetEnvString(key string, defaultVal string) string {
	return GetEnv(key, defaultVal)
}

// GetEnvInt is a convenience function for getting int values
func GetEnvInt(key string, defaultVal int) int {
	return GetEnv(key, defaultVal)
}

// GetEnvBool is a convenience function for getting bool values
func GetEnvBool(key string, defaultVal bool) bool {
	return GetEnv(key, defaultVal)
}

// GetEnvFloat64 is a convenience function for getting float64 values
func GetEnvFloat64(key string, defaultVal float64) float64 {
	return GetEnv(key, defaultVal)
}

// GetEnvDuration is a convenience function for getting time.Duration values
func GetEnvDuration(key string, defaultVal time.Duration) time.Duration {
	return GetEnv(key, defaultVal)
}
