package envlookup

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"
)

const separator = ","

var (
	// ErrNotFound error
	ErrNotFound = errors.New("Variable not found")

	// ErrParse error
	ErrParse = errors.New("Parse error")
)

// String retrieves the value of the environment variable named
// by the key. If the variable is present in the environment the
// value (which may be empty) is returned and the error is nil.
// If the variable is not present but a default value is supplied,
// that will be returned with an ErrNotFound. Otherwise the returned
// value will be empty and an ErrNotFound will be set.
func String(key string, def ...string) (string, error) {
	v, exists := os.LookupEnv(key)
	if !exists {
		var res string
		if len(def) > 0 {
			res = def[0]
		}
		return res, ErrNotFound
	}
	return v, nil
}

// Slice retrieves the value of the environment variable named
// by the key. If the variable is present in the environment the
// value (which may be empty) is returned and the error is nil.
// If the variable is not present but a default value is supplied,
// that will be returned with an ErrNotFound. Otherwise the returned
// value will be empty and an ErrNotFound will be set.
func Slice(key string, def ...[]string) ([]string, error) {
	v, exists := os.LookupEnv(key)
	if !exists {
		var res []string
		if len(def) > 0 {
			res = def[0]
		}
		return res, ErrNotFound
	}

	split := strings.Split(v, separator)

	return split, nil
}

// Int retrieves the value of the environment variable named
// by the key. If the variable is present in the environment the
// value (which may be empty) is returned and the error is nil.
// If the variable is not present but a default value is supplied,
// that will be returned with an ErrNotFound. If the env var could
// could not be parsed as an int, ErrParse will be returned.
// Otherwise the returned value will be empty and an ErrNotFound
// will be set.
func Int(key string, def ...int) (int, error) {
	v, exists := os.LookupEnv(key)
	if !exists {
		var res int
		if len(def) > 0 {
			res = def[0]
		}
		return res, ErrNotFound
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, ErrParse
	}
	return i, nil
}

// Bool retrieves the value of the environment variable named
// by the key. If the variable is present in the environment the
// value (which may be empty) is returned and the error is nil.
// If the variable is not present but a default value is supplied,
// that will be returned with an ErrNotFound. If the env var could
// could not be parsed as a bool, ErrParse will be returned.
// Otherwise the returned value will be empty and an ErrNotFound
// will be set.
func Bool(key string, def ...bool) (bool, error) {
	v, exists := os.LookupEnv(key)
	if !exists {
		var res bool
		if len(def) > 0 {
			res = def[0]
		}
		return res, ErrNotFound
	}

	return strToBool(v)
}

// Duration retrieves the value of the environment variable named
// by the key. If the variable is present in the environment the
// value (which may be empty) is returned and the error is nil.
// If the variable is not present but a default value is supplied,
// that will be returned with an ErrNotFound. If the env var could
// could not be parsed as a time.Duration value, ErrParse will be
// returned. Otherwise the returned value will be empty and an
// ErrNotFound will be set.
func Duration(key string, def ...time.Duration) (time.Duration, error) {
	v, exists := os.LookupEnv(key)
	if !exists {
		var res time.Duration
		if len(def) > 0 {
			res = def[0]
		}
		return res, ErrNotFound
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return 0, ErrParse
	}
	return d, nil
}

// Float64 retrieves the value of the environment variable named
// by the key. If the variable is present in the environment the
// value (which may be empty) is returned and the error is nil.
// If the variable is not present but a default value is supplied,
// that will be returned with an ErrNotFound. If the env var could
// could not be parsed as a float64, ErrParse will be returned.
// Otherwise the returned value will be empty and an ErrNotFound
// will be set.
func Float64(key string, def ...float64) (float64, error) {
	v, exists := os.LookupEnv(key)
	if !exists {
		var res float64
		if len(def) > 0 {
			res = def[0]
		}
		return res, ErrNotFound
	}

	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, ErrParse
	}

	return f, nil
}

func strToBool(v string) (bool, error) {
	switch strings.ToLower(v) {
	case "false":
		fallthrough
	case "0":
		return false, nil
	case "true":
		fallthrough
	case "1":
		return true, nil
	default:
		return false, ErrParse
	}
}
