package fig

import (
	"errors"
	"fmt"
	"strconv"
)

// Config caches and retrieves configurations from the environment
type Config struct {
	drivers []Driver
}

const (
	typeInt     string = "int"
	typeInt64   string = "int64"
	typeBool    string = "bool"
	typeFloat64 string = "float64"
)

// Sentinel error for undefined config variable
var ErrConfigNotFound = errors.New("config variable not found")

// build error for malformed variable
func errConfigWrongType(key, value, expType string) error {
	return fmt.Errorf("Configuration variable %s (value '%s') not of requested type %s", key, value, expType)
}

// New initializes a config object
func New(drivers ...Driver) Config {
	return Config{
		drivers: drivers,
	}
}

// get string or cache
func (c Config) get(key string) (string, error) {
	for _, driver := range c.drivers {
		if val, err := driver.Get(key); err == nil {
			return val, nil
		} else if errors.Is(err, ErrConfigNotFound) {
			continue
		} else {
			return "", fmt.Errorf("error ")
		}
	}
	return "", fmt.Errorf("%w: config key %s not found", ErrConfigNotFound, key)
}

// GetString retrieves the configured string
func (c Config) GetString(key string) (string, error) {
	return c.get(key)
}

// GetInt retrieves the configured int
func (c Config) GetInt(key string) (int, error) {
	val, err := c.get(key)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		return i, errConfigWrongType(key, val, typeInt)
	}

	return i, nil
}

// GetInt64 retrieves the configured int64
func (c Config) GetInt64(key string) (int64, error) {
	val, err := c.get(key)
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return i, errConfigWrongType(key, val, typeInt64)
	}

	return i, nil
}

// GetBool retrieves the configured bool
func (c Config) GetBool(key string) (bool, error) {
	val, err := c.get(key)
	if err != nil {
		return false, err
	}

	i, err := strconv.ParseBool(val)
	if err != nil {
		return i, errConfigWrongType(key, val, typeBool)
	}

	return i, nil
}

// GetFloat64 retrieves the configured float64
func (c Config) GetFloat64(key string) (float64, error) {
	val, err := c.get(key)
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return i, errConfigWrongType(key, val, typeFloat64)
	}

	return i, nil
}

// MustGetString retrieves the configured string or panics if undefined
func (c Config) MustGetString(key string) string {
	value, err := c.get(key)
	if err != nil {
		panic(err)
	}

	return value
}

// MustGetInt retrieves the configured int or panics
func (c Config) MustGetInt(key string) int {
	val, err := c.get(key)
	if err != nil {
		panic(err)
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		panic(errConfigWrongType(key, val, typeInt))
	}

	return i
}

// MustGetInt64 retrieves the configured int64 or panics if missing or malformed
func (c Config) MustGetInt64(key string) int64 {
	val, err := c.get(key)
	if err != nil {
		panic(err)
	}

	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		panic(errConfigWrongType(key, val, typeInt64))
	}

	return i
}

// MustGetBool retrieves the configured bool or panics if missing or malformed
func (c Config) MustGetBool(key string) bool {
	val, err := c.get(key)
	if err != nil {
		panic(err)
	}

	i, err := strconv.ParseBool(val)
	if err != nil {
		panic(errConfigWrongType(key, val, typeBool))
	}

	return i
}

// MustGetFloat64 retrieves the configured float64 or panics if missing or malformed
func (c Config) MustGetFloat64(key string) float64 {
	val, err := c.get(key)
	if err != nil {
		panic(err)
	}

	i, err := strconv.ParseFloat(val, 64)
	if err != nil {
		panic(errConfigWrongType(key, val, typeFloat64))
	}

	return i
}

// GetStringOr retrieves the configured string or the provided default
func (c Config) GetStringOr(key, defaultString string) string {
	value, err := c.get(key)
	if err != nil {
		return defaultString
	}

	return value
}

// GetIntOr retrieves the configured int or the provide ddefault
func (c Config) GetIntOr(key string, defaultInt int) int {
	val, err := c.get(key)
	if err != nil {
		return defaultInt
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		panic(errConfigWrongType(key, val, typeInt))
	}

	return i
}

// GetInt64Or retrieves the configured int64 or the provided default
func (c Config) GetInt64Or(key string, defaultInt64 int64) int64 {
	val, err := c.get(key)
	if err != nil {
		return defaultInt64
	}

	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		panic(errConfigWrongType(key, val, typeInt64))
	}

	return i
}

// GetBoolOr retrieves the configured bool or the provided default
func (c Config) GetBoolOr(key string, defaultBool bool) bool {
	val, err := c.get(key)
	if err != nil {
		return defaultBool
	}

	i, err := strconv.ParseBool(val)
	if err != nil {
		panic(errConfigWrongType(key, val, typeBool))
	}

	return i
}

// GetFloat64Or retrieves the configured float64 or the provided default
func (c Config) GetFloat64Or(key string, defaultFloat64 float64) float64 {
	val, err := c.get(key)
	if err != nil {
		return defaultFloat64
	}

	i, err := strconv.ParseFloat(val, 64)
	if err != nil {
		panic(errConfigWrongType(key, val, typeFloat64))
	}

	return i
}
