package envlookup_test

import (
	"testing"

	"github.com/spider-pigs/envlookup"
)

func TestMustBool(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Function call should not panic")
		}
	}()
	b := envlookup.MustBool(envlookup.Bool("PLAYED_WITH_MILES_DAVIES"))
	if !b {
		t.Error("value should not be false", b)
	}
}

func TestMustBoolPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Function call should panic")
		}
	}()

	envlookup.MustBool(envlookup.Bool("PANIC_PLEASE"))
}

func TestMustDuration(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Function call should not panic")
		}
	}()

	d := envlookup.MustDuration(envlookup.Duration("LONGEST_RECORDED_TRACK"))
	if d == 0 {
		t.Error("value should not be zero", d)
	}
}

func TestMustDurationPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Function call should panic")
		}
	}()

	envlookup.MustBool(envlookup.Bool("LONGEST_RECORDED_TRACK_FLOAT"))
}

func TestMustFloat64(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Function call should not panic")
		}
	}()

	f := envlookup.MustFloat64(envlookup.Float64("LONGEST_RECORDED_TRACK_FLOAT"))
	if f == 0 {
		t.Error("value should not be zero", f)
	}
}

func TestMustFloat64Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Function call should panic")
		}
	}()

	envlookup.MustFloat64(envlookup.Float64("PANIC_PLEASE"))
}

func TestMustInt(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Function call should not panic")
		}
	}()

	i := envlookup.MustInt(envlookup.Int("NO_OF_STUDIO_ALBUMS"))
	if i == 0 {
		t.Error("value should not be zero", i)
	}
}

func TestMustIntPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Function call should panic")
		}
	}()

	envlookup.MustInt(envlookup.Int("PANIC_PLEASE"))
}

func TestMustSlice(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Function call should not panic")
		}
	}()

	s := envlookup.MustSlice(envlookup.Slice("RECORD_LABELS"))
	if len(s) == 0 {
		t.Error("value should not be zero", s)
	}
}

func TestMustSlicePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Function call should panic")
		}
	}()

	envlookup.MustSlice(envlookup.Slice("PANIC_PLEASE"))
}

func TestMustString(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Function call should not panic")
		}
	}()

	s := envlookup.MustString(envlookup.String("JAZZ_ARTIST"))
	if len(s) == 0 {
		t.Error("value should be returned", s)
	}
}

func TestMustStringPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Function call should panic")
		}
	}()

	envlookup.MustString(envlookup.String("PANIC_PLEASE"))
}
