# fig
Juicy config from environment in Go

Fig wraps John Barton's [godotenv](github.com/joho/godotenv) with caching and typed config variables.

## Initializing
Calling `fig.Make()` with no parameters initializes a Config object with the contents
of `./.env`. Pass one or more filenames to read environment variables from specific files: `fig.Make("./config.env", "./example.env")`

## Retrieving config variables
Retrieving a value from the environment will cache it in the Config object to (minimally) speed up retrieval. 

``` go
conf := fig.Make()

// first call reads from environment
confStr, _ := conf.GetString("ENV_NAME")

// subsequent calls read from cache
confStrCached, _ := conf.GetString("ENV_NAME")
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