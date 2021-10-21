package cli

import (
	"io"

	"github.com/alecthomas/kong"
	"github.com/joho/godotenv"
)

// EnvLoader loads environment variables from an .env file
// and passes them to Kong for CLI handling.
func EnvLoader(r io.Reader) (kong.Resolver, error) {
	mp, err := godotenv.Parse(r)
	if err != nil {
		return nil, err
	}

	return kong.ResolverFunc(func(context *kong.Context, parent *kong.Path, flag *kong.Flag) (interface{}, error) {
		raw, ok := mp[flag.Env]
		if ok {
			return raw, nil
		}

		return nil, nil
	}), nil
}
