package envlookup

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const separator = ","

// NotFoundError indicates that an environment variable was not found.
type NotFoundError struct {
	Var string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("could not find environment variable \"%s\"", e.Var)
}

// ParseError indicates that an environment variable could not be
// parsed to the desired type.
type ParseError struct {
	Var string
	Err error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("could not parse environment variable \"%s\": %s", e.Var, e.Err)
}

// String retrieves the value of the environment variable named by the
// key. If the variable is present in the environment the value (which
// may be empty) is returned and the error is nil. If the variable is
// not present but a default value is supplied, that value will be
// returned. Otherwise the returned value will be empty and
// NotFoundError will be returned.
func String(key string, def ...string) (string, error) {
	v, exists := os.LookupEnv(key)
	if !exists {
		if len(def) > 0 {
			return def[0], nil
		}
		var res string
		err := &NotFoundError{key}
		return res, err
	}
	return v, nil
}

// Slice retrieves the value of the environment variable named by the
// key. If the variable is present in the environment the value (which
// may be empty) is returned and the error is nil. If the variable is
// not present but a default value is supplied, that value will be
// returned. Otherwise the returned value will be empty and
// NotFoundError will be returned.
func Slice(key string, def ...[]string) ([]string, error) {
	v, exists := os.LookupEnv(key)
	if !exists {
		if len(def) > 0 {
			return def[0], nil
		}
		var res []string
		err := &NotFoundError{key}
		return res, err
	}

	split := strings.Split(v, separator)

	return split, nil
}

// Int retrieves the value of the environment variable named by the
// key. If the variable is present in the environment the value (which
// may be empty) is returned and the error is nil. If the variable is
// not present but a default value is supplied, that value will be
// returned. If the env var could could not be parsed as an int value,
// ParseError will be returned. Otherwise the returned value will be
// empty and NotFoundError will be returned.
func Int(key string, def ...int) (int, error) {
	v, exists := os.LookupEnv(key)
	if !exists {
		if len(def) > 0 {
			return def[0], nil
		}
		var res int
		err := &NotFoundError{key}
		return res, err
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		var res int
		err := &ParseError{key, err}
		return res, err
	}
	return i, nil
}

// Int64 retrieves the value of the environment variable named by the
// key. If the variable is present in the environment the value (which
// may be empty) is returned and the error is nil. If the variable is
// not present but a default value is supplied, that value will be
// returned. If the env var could could not be parsed as an int value,
// ParseError will be returned. Otherwise the returned value will be
// empty and NotFoundError will be returned.
func Int64(key string, def ...int64) (int64, error) {
	v, exists := os.LookupEnv(key)
	if !exists {
		if len(def) > 0 {
			return def[0], nil
		}
		var res int64
		err := &NotFoundError{key}
		return res, err
	}

	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		var res int64
		err := &ParseError{key, err}
		return res, err
	}
	return i, nil
}

// Bool retrieves the value of the environment variable named by the
// key. If the variable is present in the environment the value (which
// may be empty) is returned and the error is nil. If the variable is
// not present but a default value is supplied, that value will be
// returned. If the env var could could not be parsed as a bool value,
// ParseError will be returned. Otherwise the returned value will be
// empty and NotFoundError will be returned.
func Bool(key string, def ...bool) (bool, error) {
	v, exists := os.LookupEnv(key)
	if !exists {
		if len(def) > 0 {
			return def[0], nil
		}
		var res bool
		err := &NotFoundError{key}
		return res, err
	}

	b, err := strToBool(v)
	if err != nil {
		var res bool
		err := &ParseError{key, err}
		return res, err
	}
	return b, nil
}

// Duration retrieves the value of the environment variable named by
// the key. If the variable is present in the environment the value
// (which may be empty) is returned and the error is nil. If the
// variable is not present but a default value is supplied, that value
// will be returned. If the env var could could not be parsed as a
// time.Duration value, ParseError will be returned. Otherwise the
// returned value will be empty and NotFoundError will be returned.
func Duration(key string, def ...time.Duration) (time.Duration, error) {
	v, exists := os.LookupEnv(key)
	if !exists {
		if len(def) > 0 {
			return def[0], nil
		}
		var res time.Duration
		err := &NotFoundError{key}
		return res, err
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		var res time.Duration
		err := &ParseError{key, err}
		return res, err
	}
	return d, nil
}

// Float64 retrieves the value of the environment variable named by the
// key. If the variable is present in the environment the value (which
// may be empty) is returned and the error is nil. If the variable is
// not present but a default value is supplied, that value will be
// returned. If the env var could could not be parsed as a float64,
// ParseError will be returned. Otherwise the returned value will be
// empty and NotFoundError will be returned.
func Float64(key string, def ...float64) (float64, error) {
	v, exists := os.LookupEnv(key)
	if !exists {
		if len(def) > 0 {
			return def[0], nil
		}
		var res float64
		err := &NotFoundError{key}
		return res, err
	}

	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		var res float64
		err := &ParseError{key, err}
		return res, err
	}

	return f, nil
}

// Uint64 retrieves the value of the environment variable named by the
// key. If the variable is present in the environment the value (which
// may be empty) is returned and the error is nil. If the variable is
// not present but a default value is supplied, that value will be
// returned. Otherwise the returned value will be empty and
// NotFoundError will be returned.
func Uint64(key string, def ...uint64) (uint64, error) {
	v, exists := os.LookupEnv(key)
	if !exists {
		if len(def) > 0 {
			return def[0], nil
		}
		var res uint64
		err := &NotFoundError{key}
		return res, err
	}

	u, err := strconv.ParseUint(v, 0, 64)
	if err != nil {
		var res uint64
		err := &ParseError{key, err}
		return res, err
	}

	return u, nil
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
		var res bool
		errStr := fmt.Sprintf("\"%s\" is not parseable as bool value", v)
		err := errors.New(errStr)
		return res, err
	}
}
