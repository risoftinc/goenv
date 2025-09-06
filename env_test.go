package goenv

import (
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{
			name:     "non-existent file",
			filename: "non_existent.env",
			wantErr:  true,
		},
		{
			name:     "empty filename",
			filename: "",
			wantErr:  true, // Our implementation returns error for empty filename
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := LoadEnv(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadEnv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetEnv_String(t *testing.T) {
	tests := []struct {
		name       string
		key        string
		value      string
		defaultVal string
		want       string
		clearEnv   bool
	}{
		{
			name:       "existing env var",
			key:        "TEST_STRING",
			value:      "test_value",
			defaultVal: "default_value",
			want:       "test_value",
		},
		{
			name:       "non-existing env var",
			key:        "NON_EXISTENT",
			value:      "",
			defaultVal: "default_value",
			want:       "default_value",
			clearEnv:   true,
		},
		{
			name:       "empty env var",
			key:        "EMPTY_STRING",
			value:      "",
			defaultVal: "default_value",
			want:       "default_value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.clearEnv {
				os.Unsetenv(tt.key)
			} else {
				os.Setenv(tt.key, tt.value)
			}

			got := GetEnv(tt.key, tt.defaultVal)
			if got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnv_Int(t *testing.T) {
	tests := []struct {
		name       string
		key        string
		value      string
		defaultVal int
		want       int
		clearEnv   bool
	}{
		{
			name:       "valid integer",
			key:        "TEST_INT",
			value:      "42",
			defaultVal: 0,
			want:       42,
		},
		{
			name:       "invalid integer",
			key:        "INVALID_INT",
			value:      "not_a_number",
			defaultVal: 100,
			want:       100,
		},
		{
			name:       "non-existing env var",
			key:        "NON_EXISTENT_INT",
			value:      "",
			defaultVal: 50,
			want:       50,
			clearEnv:   true,
		},
		{
			name:       "negative integer",
			key:        "NEGATIVE_INT",
			value:      "-10",
			defaultVal: 0,
			want:       -10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.clearEnv {
				os.Unsetenv(tt.key)
			} else {
				os.Setenv(tt.key, tt.value)
			}

			got := GetEnv(tt.key, tt.defaultVal)
			if got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnv_Uint64(t *testing.T) {
	tests := []struct {
		name       string
		key        string
		value      string
		defaultVal uint64
		want       uint64
		clearEnv   bool
	}{
		{
			name:       "valid uint64",
			key:        "TEST_UINT64",
			value:      "18446744073709551615",
			defaultVal: 0,
			want:       18446744073709551615,
		},
		{
			name:       "invalid uint64",
			key:        "INVALID_UINT64",
			value:      "not_a_number",
			defaultVal: 100,
			want:       100,
		},
		{
			name:       "negative number",
			key:        "NEGATIVE_UINT64",
			value:      "-10",
			defaultVal: 50,
			want:       50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.clearEnv {
				os.Unsetenv(tt.key)
			} else {
				os.Setenv(tt.key, tt.value)
			}

			got := GetEnv(tt.key, tt.defaultVal)
			if got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnv_Int64(t *testing.T) {
	tests := []struct {
		name       string
		key        string
		value      string
		defaultVal int64
		want       int64
		clearEnv   bool
	}{
		{
			name:       "valid int64",
			key:        "TEST_INT64",
			value:      "9223372036854775807",
			defaultVal: 0,
			want:       9223372036854775807,
		},
		{
			name:       "negative int64",
			key:        "NEGATIVE_INT64",
			value:      "-9223372036854775808",
			defaultVal: 0,
			want:       -9223372036854775808,
		},
		{
			name:       "invalid int64",
			key:        "INVALID_INT64",
			value:      "not_a_number",
			defaultVal: 100,
			want:       100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.clearEnv {
				os.Unsetenv(tt.key)
			} else {
				os.Setenv(tt.key, tt.value)
			}

			got := GetEnv(tt.key, tt.defaultVal)
			if got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnv_Bool(t *testing.T) {
	tests := []struct {
		name       string
		key        string
		value      string
		defaultVal bool
		want       bool
		clearEnv   bool
	}{
		{
			name:       "true value",
			key:        "TEST_BOOL_TRUE",
			value:      "true",
			defaultVal: false,
			want:       true,
		},
		{
			name:       "false value",
			key:        "TEST_BOOL_FALSE",
			value:      "false",
			defaultVal: true,
			want:       false,
		},
		{
			name:       "1 as true",
			key:        "TEST_BOOL_ONE",
			value:      "1",
			defaultVal: false,
			want:       true,
		},
		{
			name:       "0 as false",
			key:        "TEST_BOOL_ZERO",
			value:      "0",
			defaultVal: true,
			want:       false,
		},
		{
			name:       "invalid bool",
			key:        "INVALID_BOOL",
			value:      "maybe",
			defaultVal: true,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.clearEnv {
				os.Unsetenv(tt.key)
			} else {
				os.Setenv(tt.key, tt.value)
			}

			got := GetEnv(tt.key, tt.defaultVal)
			if got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnv_Float64(t *testing.T) {
	tests := []struct {
		name       string
		key        string
		value      string
		defaultVal float64
		want       float64
		clearEnv   bool
	}{
		{
			name:       "valid float",
			key:        "TEST_FLOAT",
			value:      "3.14159",
			defaultVal: 0.0,
			want:       3.14159,
		},
		{
			name:       "negative float",
			key:        "NEGATIVE_FLOAT",
			value:      "-2.718",
			defaultVal: 0.0,
			want:       -2.718,
		},
		{
			name:       "integer as float",
			key:        "INT_AS_FLOAT",
			value:      "42",
			defaultVal: 0.0,
			want:       42.0,
		},
		{
			name:       "invalid float",
			key:        "INVALID_FLOAT",
			value:      "not_a_number",
			defaultVal: 1.5,
			want:       1.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.clearEnv {
				os.Unsetenv(tt.key)
			} else {
				os.Setenv(tt.key, tt.value)
			}

			got := GetEnv(tt.key, tt.defaultVal)
			if got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnv_UnknownType(t *testing.T) {
	// Test with an unknown type (should return default value)
	type CustomType struct {
		Value string
	}

	defaultVal := CustomType{Value: "default"}

	// Set an environment variable
	os.Setenv("TEST_CUSTOM", "some_value")

	got := GetEnv("TEST_CUSTOM", defaultVal)
	if got != defaultVal {
		t.Errorf("GetEnv() = %v, want %v", got, defaultVal)
	}
}

func TestGetEnv_Cleanup(t *testing.T) {
	// Clean up environment variables set during tests
	envVars := []string{
		"TEST_STRING", "EMPTY_STRING", "TEST_INT", "INVALID_INT", "NEGATIVE_INT",
		"TEST_UINT64", "INVALID_UINT64", "NEGATIVE_UINT64", "TEST_INT64", "NEGATIVE_INT64", "INVALID_INT64",
		"TEST_BOOL_TRUE", "TEST_BOOL_FALSE", "TEST_BOOL_ONE", "TEST_BOOL_ZERO", "INVALID_BOOL",
		"TEST_FLOAT", "NEGATIVE_FLOAT", "INT_AS_FLOAT", "INVALID_FLOAT", "TEST_CUSTOM",
	}

	for _, envVar := range envVars {
		os.Unsetenv(envVar)
	}
}

// Test new functionality for multiple file formats

func TestDetectFormat(t *testing.T) {
	tests := []struct {
		filename string
		expected FileFormat
	}{
		{"config.env", FormatKeyValue},
		{"config.json", FormatJSON},
		{"config.yaml", FormatYAML},
		{"config.yml", FormatYAML},
		{"config.txt", FormatKeyValue}, // Default fallback
		{"config", FormatKeyValue},     // No extension
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			result := detectFormat(tt.filename)
			if result != tt.expected {
				t.Errorf("detectFormat(%s) = %v, want %v", tt.filename, result, tt.expected)
			}
		})
	}
}

func TestLoadKeyValueFile(t *testing.T) {
	// Create a temporary .env file
	content := `# This is a comment
APP_NAME=MyApp
APP_VERSION=1.0.0
DEBUG=true
PORT=8080
EMPTY_VALUE=
QUOTED_VALUE="quoted string"
SINGLE_QUOTED='single quoted'
INVALID_LINE=value=extra=equals
CONNECTION=MYSQL  #Mysql, Postgres, Etc
DATABASE_URL=postgres://user:pass@localhost/db  # Database connection string
API_KEY="secret-key"  # API key for external services
`

	tmpFile := createTempFile(t, ".env", content)
	defer os.Remove(tmpFile)

	// Clear any existing environment variables
	envVars := []string{"APP_NAME", "APP_VERSION", "DEBUG", "PORT", "EMPTY_VALUE", "QUOTED_VALUE", "SINGLE_QUOTED", "CONNECTION", "DATABASE_URL", "API_KEY"}
	for _, envVar := range envVars {
		os.Unsetenv(envVar)
	}

	err := loadKeyValueFile(tmpFile)
	if err != nil {
		t.Fatalf("loadKeyValueFile() error = %v", err)
	}

	// Test loaded values
	tests := []struct {
		key   string
		value string
	}{
		{"APP_NAME", "MyApp"},
		{"APP_VERSION", "1.0.0"},
		{"DEBUG", "true"},
		{"PORT", "8080"},
		{"EMPTY_VALUE", ""},
		{"QUOTED_VALUE", "quoted string"},
		{"SINGLE_QUOTED", "single quoted"},
		{"CONNECTION", "MYSQL"},
		{"DATABASE_URL", "postgres://user:pass@localhost/db"},
		{"API_KEY", "secret-key"},
	}

	for _, tt := range tests {
		if got := os.Getenv(tt.key); got != tt.value {
			t.Errorf("Environment variable %s = %v, want %v", tt.key, got, tt.value)
		}
	}

	// Clean up
	for _, envVar := range envVars {
		os.Unsetenv(envVar)
	}
}

func TestLoadKeyValueFileWithInlineComments(t *testing.T) {
	// Test various inline comment scenarios
	content := `# Test inline comments
VALUE_WITH_COMMENT=test  # This is a comment
QUOTED_WITH_COMMENT="test value"  # Comment after quoted value
SINGLE_QUOTED_WITH_COMMENT='test value'  # Comment after single quoted value
HASH_IN_QUOTES="test#value"  # This should not be treated as comment
SINGLE_HASH_IN_QUOTES='test#value'  # This should not be treated as comment
MULTIPLE_HASH="test#value#here"  # Only the last # should be treated as comment
EMPTY_AFTER_COMMENT=  # This should result in empty value
`

	tmpFile := createTempFile(t, ".env", content)
	defer os.Remove(tmpFile)

	// Clear any existing environment variables
	envVars := []string{"VALUE_WITH_COMMENT", "QUOTED_WITH_COMMENT", "SINGLE_QUOTED_WITH_COMMENT",
		"HASH_IN_QUOTES", "SINGLE_HASH_IN_QUOTES", "MULTIPLE_HASH", "EMPTY_AFTER_COMMENT"}
	for _, envVar := range envVars {
		os.Unsetenv(envVar)
	}

	err := loadKeyValueFile(tmpFile)
	if err != nil {
		t.Fatalf("loadKeyValueFile() error = %v", err)
	}

	// Test loaded values
	tests := []struct {
		key   string
		value string
	}{
		{"VALUE_WITH_COMMENT", "test"},
		{"QUOTED_WITH_COMMENT", "test value"},
		{"SINGLE_QUOTED_WITH_COMMENT", "test value"},
		{"HASH_IN_QUOTES", "test#value"},
		{"SINGLE_HASH_IN_QUOTES", "test#value"},
		{"MULTIPLE_HASH", "test#value#here"},
		{"EMPTY_AFTER_COMMENT", ""},
	}

	for _, tt := range tests {
		if got := os.Getenv(tt.key); got != tt.value {
			t.Errorf("Environment variable %s = %v, want %v", tt.key, got, tt.value)
		}
	}

	// Clean up
	for _, envVar := range envVars {
		os.Unsetenv(envVar)
	}
}

func TestLoadJSONFile(t *testing.T) {
	// Create a temporary JSON file
	content := `{
		"app": {
			"name": "MyApp",
			"version": "1.0.0",
			"debug": true,
			"port": 8080
		},
		"database": {
			"host": "localhost",
			"port": 5432,
			"name": "mydb"
		},
		"features": ["auth", "logging", "metrics"],
		"timeout": 30.5
	}`

	tmpFile := createTempFile(t, ".json", content)
	defer os.Remove(tmpFile)

	// Clear any existing environment variables
	envVars := []string{"APP_NAME", "APP_VERSION", "APP_DEBUG", "APP_PORT", "DATABASE_HOST", "DATABASE_PORT", "DATABASE_NAME", "FEATURES", "TIMEOUT"}
	for _, envVar := range envVars {
		os.Unsetenv(envVar)
	}

	err := loadJSONFile(tmpFile)
	if err != nil {
		t.Fatalf("loadJSONFile() error = %v", err)
	}

	// Test loaded values
	tests := []struct {
		key   string
		value string
	}{
		{"app.name", "MyApp"},
		{"app.version", "1.0.0"},
		{"app.debug", "true"},
		{"app.port", "8080"},
		{"database.host", "localhost"},
		{"database.port", "5432"},
		{"database.name", "mydb"},
		{"features", `["auth","logging","metrics"]`},
		{"timeout", "30.5"},
	}

	for _, tt := range tests {
		if got := os.Getenv(tt.key); got != tt.value {
			t.Errorf("Environment variable %s = %v, want %v", tt.key, got, tt.value)
		}
	}

	// Clean up
	for _, envVar := range envVars {
		os.Unsetenv(envVar)
	}
}

func TestLoadYAMLFile(t *testing.T) {
	// Create a temporary YAML file
	content := `app:
  name: MyApp
  version: 1.0.0
  debug: true
  port: 8080

database:
  host: localhost
  port: 5432
  name: mydb

features:
  - auth
  - logging
  - metrics

timeout: 30.5
`

	tmpFile := createTempFile(t, ".yaml", content)
	defer os.Remove(tmpFile)

	// Clear any existing environment variables
	envVars := []string{"APP_NAME", "APP_VERSION", "APP_DEBUG", "APP_PORT", "DATABASE_HOST", "DATABASE_PORT", "DATABASE_NAME", "FEATURES", "TIMEOUT"}
	for _, envVar := range envVars {
		os.Unsetenv(envVar)
	}

	err := loadYAMLFile(tmpFile)
	if err != nil {
		t.Fatalf("loadYAMLFile() error = %v", err)
	}

	// Test loaded values
	tests := []struct {
		key   string
		value string
	}{
		{"app.name", "MyApp"},
		{"app.version", "1.0.0"},
		{"app.debug", "true"},
		{"app.port", "8080"},
		{"database.host", "localhost"},
		{"database.port", "5432"},
		{"database.name", "mydb"},
		{"features", `["auth","logging","metrics"]`},
		{"timeout", "30.5"},
	}

	for _, tt := range tests {
		if got := os.Getenv(tt.key); got != tt.value {
			t.Errorf("Environment variable %s = %v, want %v", tt.key, got, tt.value)
		}
	}

	// Clean up
	for _, envVar := range envVars {
		os.Unsetenv(envVar)
	}
}

func TestLoadEnvWithFormat(t *testing.T) {
	// Create temporary files
	envContent := `APP_NAME=MyApp
APP_VERSION=1.0.0`

	jsonContent := `{
		"app": {
			"name": "MyAppJSON",
			"version": "2.0.0"
		}
	}`

	yamlContent := `app:
  name: MyAppYAML
  version: 3.0.0`

	envFile := createTempFile(t, ".env", envContent)
	jsonFile := createTempFile(t, ".json", jsonContent)
	yamlFile := createTempFile(t, ".yaml", yamlContent)

	defer os.Remove(envFile)
	defer os.Remove(jsonFile)
	defer os.Remove(yamlFile)

	// Clear environment variables
	envVars := []string{"APP_NAME", "APP_VERSION", "app.name", "app.version"}
	for _, envVar := range envVars {
		os.Unsetenv(envVar)
	}

	// Test loading with specific format
	t.Run("LoadEnv with FormatKeyValue", func(t *testing.T) {
		err := LoadEnvWithFormat(FormatKeyValue, envFile)
		if err != nil {
			t.Fatalf("LoadEnvWithFormat() error = %v", err)
		}

		if got := os.Getenv("APP_NAME"); got != "MyApp" {
			t.Errorf("APP_NAME = %v, want MyApp", got)
		}
	})

	// Clean up and test JSON
	for _, envVar := range envVars {
		os.Unsetenv(envVar)
	}

	t.Run("LoadEnv with FormatJSON", func(t *testing.T) {
		err := LoadEnvWithFormat(FormatJSON, jsonFile)
		if err != nil {
			t.Fatalf("LoadEnvWithFormat() error = %v", err)
		}

		if got := os.Getenv("app.name"); got != "MyAppJSON" {
			t.Errorf("app.name = %v, want MyAppJSON", got)
		}
	})

	// Clean up and test YAML
	for _, envVar := range envVars {
		os.Unsetenv(envVar)
	}

	t.Run("LoadEnv with FormatYAML", func(t *testing.T) {
		err := LoadEnvWithFormat(FormatYAML, yamlFile)
		if err != nil {
			t.Fatalf("LoadEnvWithFormat() error = %v", err)
		}

		if got := os.Getenv("app.name"); got != "MyAppYAML" {
			t.Errorf("app.name = %v, want MyAppYAML", got)
		}
	})

	// Clean up
	for _, envVar := range envVars {
		os.Unsetenv(envVar)
	}
}

func TestGetEnvNested(t *testing.T) {
	// Set up nested environment variables
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("API_TIMEOUT", "30")
	os.Setenv("API_RETRY_COUNT", "3")

	defer func() {
		os.Unsetenv("DB_HOST")
		os.Unsetenv("DB_PORT")
		os.Unsetenv("API_TIMEOUT")
		os.Unsetenv("API_RETRY_COUNT")
	}()

	tests := []struct {
		name       string
		key        string
		defaultVal interface{}
		want       interface{}
	}{
		{
			name:       "nested string",
			key:        "db.host",
			defaultVal: "default_host",
			want:       "localhost",
		},
		{
			name:       "nested int",
			key:        "db.port",
			defaultVal: 3306,
			want:       5432,
		},
		{
			name:       "nested int with dot notation",
			key:        "api.timeout",
			defaultVal: 60,
			want:       30,
		},
		{
			name:       "non-existent nested key",
			key:        "non.existent",
			defaultVal: "default",
			want:       "default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.defaultVal.(type) {
			case string:
				got := GetEnvNested(tt.key, v)
				if got != tt.want {
					t.Errorf("GetEnvNested() = %v, want %v", got, tt.want)
				}
			case int:
				got := GetEnvNested(tt.key, v)
				if got != tt.want {
					t.Errorf("GetEnvNested() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestConvenienceFunctions(t *testing.T) {
	// Set up test environment variables
	os.Setenv("TEST_STRING", "test_value")
	os.Setenv("TEST_INT", "42")
	os.Setenv("TEST_BOOL", "true")
	os.Setenv("TEST_FLOAT", "3.14")

	defer func() {
		os.Unsetenv("TEST_STRING")
		os.Unsetenv("TEST_INT")
		os.Unsetenv("TEST_BOOL")
		os.Unsetenv("TEST_FLOAT")
	}()

	t.Run("GetEnvString", func(t *testing.T) {
		got := GetEnvString("TEST_STRING", "default")
		if got != "test_value" {
			t.Errorf("GetEnvString() = %v, want test_value", got)
		}
	})

	t.Run("GetEnvInt", func(t *testing.T) {
		got := GetEnvInt("TEST_INT", 0)
		if got != 42 {
			t.Errorf("GetEnvInt() = %v, want 42", got)
		}
	})

	t.Run("GetEnvBool", func(t *testing.T) {
		got := GetEnvBool("TEST_BOOL", false)
		if got != true {
			t.Errorf("GetEnvBool() = %v, want true", got)
		}
	})

	t.Run("GetEnvFloat64", func(t *testing.T) {
		got := GetEnvFloat64("TEST_FLOAT", 0.0)
		if got != 3.14 {
			t.Errorf("GetEnvFloat64() = %v, want 3.14", got)
		}
	})
}

// Helper function to create temporary files for testing
func createTempFile(t *testing.T, suffix, content string) string {
	tmpFile, err := os.CreateTemp("", "goenv_test_*"+suffix)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	if err := tmpFile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	return tmpFile.Name()
}
