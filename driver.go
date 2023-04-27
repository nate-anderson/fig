package fig

import (
	"errors"
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
type EnvironmentDriver struct{}

func NewEnvironmentDriver(filenames ...string) (EnvironmentDriver, error) {
	var err error
	if len(filenames) > 0 {
		err = godotenv.Load(filenames...)
	}
	return EnvironmentDriver{}, err
}

func (d EnvironmentDriver) Get(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
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
		} else if !errors.Is(err, os.ErrNotExist) {
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
