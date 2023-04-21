package fig

import (
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
	err := godotenv.Load(filenames...)
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
