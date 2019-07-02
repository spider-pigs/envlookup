package envlookup_test

import (
	"envlookup"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	setVars()
	defer unsetVars()
	code := m.Run()
	os.Exit(code)
}

func setVars() {
	os.Setenv("JAZZ_ARTIST", "John Coltrane")
	os.Setenv("PLAYED_WITH_MILES_DAVIES", "true")
	os.Setenv("NO_OF_STUDIO_ALBUMS", "51")
	os.Setenv("RECORD_LABELS", "Impulse!,Atlantic,Prestige,Blue Note")
	os.Setenv("LONGEST_RECORDED_TRACK", "27m32s")
	os.Setenv("LONGEST_RECORDED_TRACK_FLOAT", "27.32")
}

func unsetVars() {
	os.Unsetenv("JAZZ_ARTIST")
	os.Unsetenv("PLAYED_WITH_MILES_DAVIES")
	os.Unsetenv("NO_OF_STUDIO_ALBUMS")
	os.Unsetenv("RECORD_LABELS")
	os.Unsetenv("LONGEST_RECORDED_TRACK")
	os.Unsetenv("LONGEST_RECORDED_TRACK_FLOAT")
}

func TestEnv(t *testing.T) {
	val, err := envlookup.String("JAZZ_ARTIST")
	if len(val) == 0 {
		t.Error("value should not be empty")
	}
	if err != nil {
		t.Error("error should be nil")
	}
}

func TestEnvWithDef(t *testing.T) {
	defval := "Wayne Shorter"
	val, err := envlookup.String("JAZZ_ARTIST", defval)
	if val == defval {
		t.Error("default value should not be set")
	}
	if err != nil {
		t.Error("error should be nil")
	}
}

func TestEmptyEnv(t *testing.T) {
	val, err := envlookup.String("EMPTY_JAZZ_ARTIST")
	if len(val) != 0 {
		t.Error("value should be empty")
	}
	if err != envlookup.ErrNotFound {
		t.Error("error should be nil")
	}
}

func TestEmptyEnvWithDef(t *testing.T) {
	defval := "Wayne Shorter"
	val, err := envlookup.String("EMPTY_JAZZ_ARTIST", defval)
	if val != defval {
		t.Error("default value should be set")
	}
	if err != envlookup.ErrNotFound {
		t.Error("error should be envlookup.ErrNotFound")
	}
}

func TestSliceEnv(t *testing.T) {
	val, err := envlookup.Slice("RECORD_LABELS")
	if len(val) == 0 {
		t.Error("value should not be empty")
	}
	if err != nil {
		t.Error("error should be nil")
	}
}

func TestSliceEnvWithDef(t *testing.T) {
	defval := []string{"Impulse!", "Atlantic"}
	val, err := envlookup.Slice("RECORD_LABELS", defval)

	if reflect.DeepEqual(val, defval) {
		t.Error("default value should not be set")
	}
	if err != nil {
		t.Error("error should be nil")
	}
}

func TestEmptySliceEnv(t *testing.T) {
	val, err := envlookup.Slice("EMPTY_RECORD_LABELS")
	if len(val) != 0 {
		t.Error("value should be empty")
	}
	if err != envlookup.ErrNotFound {
		t.Error("error should be envlookup.ErrNotFound")
	}
}

func TestEmptySliceEnvWithDef(t *testing.T) {
	defval := []string{"Impulse!", "Atlantic"}
	val, err := envlookup.Slice("EMPTY_RECORD_LABELS", defval)

	if !reflect.DeepEqual(val, defval) {
		t.Error("default value should be set")
	}
	if err != envlookup.ErrNotFound {
		t.Error("error should be envlookup.ErrNotFound", err)
	}
}

func TestIntEnv(t *testing.T) {
	val, err := envlookup.Int("NO_OF_STUDIO_ALBUMS")
	if val == 0 {
		t.Error("value should not be zero")
	}
	if err != nil {
		t.Error("error should be nil")
	}
}

func TestIntEnvOrDef(t *testing.T) {
	defval := 1
	val, err := envlookup.Int("NO_OF_STUDIO_ALBUMS", defval)
	if val == defval {
		t.Error("default value should not be set")
	}
	if err != nil {
		t.Error("error should be nil")
	}
}

func TestEmptyIntEnv(t *testing.T) {
	val, err := envlookup.Int("EMPTY_NO_OF_STUDIO_ALBUMS")
	if val != 0 {
		t.Error("value should be zero")
	}
	if err != envlookup.ErrNotFound {
		t.Error("error should be envlookup.ErrNotFound")
	}
}

func TestEmptyIntEnvOrDef(t *testing.T) {
	defval := 51
	val, err := envlookup.Int("EMPTY_NO_OF_STUDIO_ALBUMS", defval)
	if val != defval {
		t.Error("default value should be set")
	}
	if err != envlookup.ErrNotFound {
		t.Error("error should be envlookup.ErrNotFound")
	}
}

func TestFloat64Env(t *testing.T) {
	val, err := envlookup.Float64("LONGEST_RECORDED_TRACK_FLOAT")
	if val == 0 {
		t.Error("value should not be zero")
	}
	if err != nil {
		t.Error("error should be nil")
	}
}

func TestFloat64EnvOrDef(t *testing.T) {
	defval := 4.43
	val, err := envlookup.Float64("LONGEST_RECORDED_TRACK_FLOAT", defval)
	if val == defval {
		t.Error("default value should not be set")
	}
	if err != nil {
		t.Error("error should be nil")
	}
}

func TestEmptyFloat64Env(t *testing.T) {
	val, err := envlookup.Float64("EMPTY_LONGEST_RECORDED_TRACK_FLOAT")
	if val != 0 {
		t.Error("value should be zero")
	}
	if err != envlookup.ErrNotFound {
		t.Error("error should be envlookup.ErrNotFound")
	}
}

func TestEmptyFloat64EnvOrDef(t *testing.T) {
	defval := 4.43
	val, err := envlookup.Float64("EMPTY_LONGEST_RECORDED_TRACK_FLOAT", defval)
	if val != defval {
		t.Error("default value should be set")
	}
	if err != envlookup.ErrNotFound {
		t.Error("error should be envlookup.ErrNotFound")
	}
}

func TestDurationEnv(t *testing.T) {
	val, err := envlookup.Duration("LONGEST_RECORDED_TRACK")
	if val == 0 {
		t.Error("value should not be 0")
	}
	if err != nil {
		t.Error("error should be nil")
	}
}

func TestDurationEnvOrDef(t *testing.T) {
	defval, _ := time.ParseDuration("10h2m")
	val, err := envlookup.Duration("LONGEST_RECORDED_TRACK", defval)
	if val == defval {
		t.Error("default value should not be set")
	}
	if err != nil {
		t.Error("error should be nil")
	}
}

func TestEmptyDurationEnv(t *testing.T) {
	val, err := envlookup.Duration("EMPTY_LONGEST_RECORDED_TRACK")
	if val != 0 {
		t.Error("value should be false")
	}
	if err != envlookup.ErrNotFound {
		t.Error("error should be envlookup.ErrNotFound")
	}
}

func TestEmptyDurationEnvOrDef(t *testing.T) {
	defval, _ := time.ParseDuration("10h2m")
	val, err := envlookup.Duration("EMPTY_LONGEST_RECORDED_TRACK", defval)
	if val != defval {
		t.Error("default value should be set")
	}
	if err != envlookup.ErrNotFound {
		t.Error("error should be envlookup.ErrNotFound")
	}
}

func TestBoolEnv(t *testing.T) {
	val, err := envlookup.Bool("PLAYED_WITH_MILES_DAVIES")
	if !val {
		t.Error("value should be true")
	}
	if err != nil {
		t.Error("error should be nil")
	}
}

func TestBoolEnvOrDef(t *testing.T) {
	defval := false
	val, err := envlookup.Bool("PLAYED_WITH_MILES_DAVIES", defval)
	if val == defval {
		t.Error("default value should not be set")
	}
	if err != nil {
		t.Error("error should be nil")
	}
}

func TestEmptyBoolEnv(t *testing.T) {
	val, err := envlookup.Bool("EMPTY_PLAYED_WITH_MILES_DAVIES")
	if val {
		t.Error("value should be false")
	}
	if err != envlookup.ErrNotFound {
		t.Error("error should be envlookup.ErrNotFound")
	}
}

func TestEmptyBoolEnvOrDef(t *testing.T) {
	defval := true
	val, err := envlookup.Bool("EMPTY_PLAYED_WITH_MILES_DAVIES", defval)
	if val != defval {
		t.Error("default value should be set")
	}
	if err != envlookup.ErrNotFound {
		t.Error("error should be envlookup.ErrNotFound")
	}
}

func TestBoolStrParse(t *testing.T) {
	val, err := envlookup.Bool("PLAYED_WITH_MILES_DAVIES")
	if !val {
		t.Error("value should be true")
	}
	if err != nil {
		t.Error("error should be nil")
	}

	os.Setenv("PLAYED_WITH_MILES_DAVIES", "1")
	val, err = envlookup.Bool("PLAYED_WITH_MILES_DAVIES")
	if !val {
		t.Error("value should be true")
	}
	if err != nil {
		t.Error("error should be nil")
	}

	os.Setenv("PLAYED_WITH_MILES_DAVIES", "false")
	val, err = envlookup.Bool("PLAYED_WITH_MILES_DAVIES")
	if val {
		t.Error("value should be false")
	}
	if err != nil {
		t.Error("error should be nil")
	}

	os.Setenv("PLAYED_WITH_MILES_DAVIES", "0")
	val, err = envlookup.Bool("PLAYED_WITH_MILES_DAVIES")
	if val {
		t.Error("value should be false")
	}
	if err != nil {
		t.Error("error should be nil")
	}

	os.Setenv("PLAYED_WITH_MILES_DAVIES", "f")
	val, err = envlookup.Bool("PLAYED_WITH_MILES_DAVIES")
	if val {
		t.Error("value should be false")
	}
	if err != envlookup.ErrFormat {
		t.Error("error should be nil")
	}
}

func TestIntEnvFormatErr(t *testing.T) {
	os.Setenv("NO_OF_STUDIO_ALBUMS", "ABC")
	defval := 1
	val, err := envlookup.Int("NO_OF_STUDIO_ALBUMS", defval)
	if val == defval {
		t.Error("default value should not be set", defval, val)
	}
	if err != envlookup.ErrFormat {
		t.Error("error should be nil")
	}

	val, err = envlookup.Int("NO_OF_STUDIO_ALBUMS")
	if val != 0 {
		t.Error("zero should be set")
	}
	if err != envlookup.ErrFormat {
		t.Error("error should be nil")
	}
}
