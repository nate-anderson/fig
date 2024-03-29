package fig

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Driver allows user-defined configuration sources
type Driver interface {
	// Should return empty string and ErrConfigNotFound if the key is simply not found
	Get(key string) (string, error)
	Name() string
}

// EnvironmentDriver supports reading from the environment and .env files
type EnvironmentDriver struct {
	env map[string]string
}

func NewEnvironmentDriver(filenames ...string) (EnvironmentDriver, error) {
	var err error
	var env = map[string]string{}
	if len(filenames) > 0 {
		env, err = godotenv.Read(filenames...)
		if err != nil {
			return EnvironmentDriver{}, err
		}
	}
	return EnvironmentDriver{env: env}, nil
}

// Get returns values from the environment, preferring real environment variables
// above those from .env files
func (d EnvironmentDriver) Get(key string) (string, error) {
	if envVal := os.Getenv(key); envVal != "" {
		return envVal, nil
	}

	val, ok := d.env[key]
	if !ok {
		return "", ErrConfigNotFound
	}

	return val, nil
}

func (d EnvironmentDriver) Name() string {
	return "env"
}

// For use when environment files may not be present in all environments
func NewOptionalFileEnvironmentDriver(filenames ...string) (EnvironmentDriver, error) {
	presentFiles := []string{}
	for _, f := range filenames {
		if _, err := os.Stat(f); err == nil {
			presentFiles = append(presentFiles, f)
		} else if !os.IsNotExist(err) {
			return EnvironmentDriver{}, fmt.Errorf("failed checking for file %s: %w", f, err)
		}
	}
	if len(presentFiles) > 0 {
		if err := godotenv.Load(presentFiles...); err != nil {
			return EnvironmentDriver{}, err
		}
	}
	return EnvironmentDriver{}, nil
}
