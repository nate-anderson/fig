<img src="https://github.com/nate-anderson/fig/blob/master/figs.jpg" width="250" alt="Figs">

# fig

[![Build and test Go package](https://github.com/nate-anderson/fig/actions/workflows/ci.yml/badge.svg)](https://github.com/nate-anderson/fig/actions/workflows/ci.yml)

[![Go Report Card](https://goreportcard.com/badge/github.com/nate-anderson/fig/v2)](https://goreportcard.com/report/github.com/nate-anderson/fig/v2)

[![Go Reference](https://pkg.go.dev/badge/github.com/nate-anderson/fig/v2.svg)](https://pkg.go.dev/github.com/nate-anderson/fig/v2)

Juicy extensible config in Go

Fig wraps John Barton's [godotenv](https://github.com/joho/godotenv) with convenience methods for type validation and struct unmarshaling. I wrote this as a subpackage
for another project and extracted it to this repo for easy reuse.

`go get github.com/nate-anderson/fig/v2`

## Initializing

Calling `fig.New()` with no parameters initializes a Config object that reads from the
environment. Passing drivers to `New()` tells `fig` where to check for configuration variables
and your preferred precedence. A driver is any type that allows `fig` to read string values
by key.

A driver for reading from `.env` files is included.

```go
envDriver, err := NewEnvironmentDriver(".env", "local.env")
conf := fig.New(envDriver)
```

`fig.Config` has methods for retrieving `string`s, `int`s, `int64`s, `float64`s and `bool`s.

- GetString
- GetInt
- GetInt64
- GetBool
- GetFloat64

These methods return an `error` if the variable is not configured or if the string in the environment cannot be parsed into the appropriate type. The following methods will panic on missing or malformed variables:

- MustGetString
- MustGetInt
- MustGetInt64
- MustGetBool
- MustGetFloat64

These are useful for required configuration variables for which repeated error checks are a PITA.

For retrieving variables with hard-coded defaults, use the `...Or` set of methods:

- GetStringOr
- GetIntOr
- GetInt64Or
- GetBoolOr
- GetFloatOr

## Struct Unmarshaling

Use the `fig` and `required` struct tags to decorate your configuration structs and quickly
read from your configuration providers. This function makes use of reflection so it is not necessarily optimal for repeated runtime unmarshals.

```go
type Config struct {
    DBHost string  `fig:"DB_HOST" required:"true"`
    DBPort *int    `fig:"DB_PORT"`
    DBUser string  `fig:"DB_USER"`
    DBPass *string `fig:"DB_PASS"`
}

var appConfig Config

func init() {
    envDriver, _ := NewEnvironmentDriver(".env", "local.env")
    conf := fig.New(envDriver)
    err := conf.Unmarshal(&appConfig)
}
```

If you need more advanced struct unmarshaling for configuration, I recommend [viper](https://github.com/spf13/viper).
