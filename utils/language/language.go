package language

import (
	"errors"
	"os"
	"strings"
)

func Get() (string, error) {
	// Check the LANG environment variable, common on UNIX.
	envlang, ok := os.LookupEnv("LANG")
	if ok {
		return strings.Split(envlang, ".")[0], nil
	}
	return "", errors.New("Can't retrieve envvar $LANG")
}
