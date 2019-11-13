package envlookup

import "time"

// MustBool is a helper that wraps a call to a function returning
// (bool, error) and panics if the error is non-nil. It is intended for
// use such as
//	b := envlookup.MustBool(envlookup.Bool("key"))
func MustBool(b bool, err error) bool {
	if err != nil {
		panic(err)
	}
	return b
}

// MustDuration is a helper that wraps a call to a function returning
// (time.Duration, error) and panics if the error is non-nil. It is
// intended for use such as
//	d := envlookup.MustDuration(envlookup.Duration("key"))
func MustDuration(d time.Duration, err error) time.Duration {
	if err != nil {
		panic(err)
	}
	return d
}

// MustFloat64 is a helper that wraps a call to a function returning
// (float64, error) and panics if the error is non-nil. It is intended
// for use such as
//	f := envlookup.MustFloat64(envlookup.Float64("key"))
func MustFloat64(f float64, err error) float64 {
	if err != nil {
		panic(err)
	}
	return f
}

// MustInt is a helper that wraps a call to a function returning
// (int, error) and panics if the error is non-nil. It is intended for
// use such as
//	i := envlookup.MustInt(envlookup.Int("key"))
func MustInt(i int, err error) int {
	if err != nil {
		panic(err)
	}
	return i
}

// MustInt64 is a helper that wraps a call to a function returning
// (int64, error) and panics if the error is non-nil. It is intended for
// use such as
//	i := envlookup.MustInt64(envlookup.Int64("key"))
func MustInt64(i int64, err error) int64 {
	if err != nil {
		panic(err)
	}
	return i
}

// MustSlice is a helper that wraps a call to a function returning
// ([]string, error) and panics if the error is non-nil. It is intended
// for use such as
//	s := envlookup.MustSlice(envlookup.Slice("key"))
func MustSlice(s []string, err error) []string {
	if err != nil {
		panic(err)
	}
	return s
}

// MustString is a helper that wraps a call to a function returning
// (string, error) and panics if the error is non-nil. It is intended
// for use such as
//	s := envlookup.MustString(envlookup.String("key"))
func MustString(s string, err error) string {
	if err != nil {
		panic(err)
	}
	return s
}

// MustUint64 is a helper that wraps a call to a function returning
// (uint64, error) and panics if the error is non-nil. It is intended
// for use such as
//	u := envlookup.MustUint64(envlookup.Uint64("key"))
func MustUint64(u uint64, err error) uint64 {
	if err != nil {
		panic(err)
	}
	return u
}
