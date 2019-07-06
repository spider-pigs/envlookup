package envlookup_test

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/spider-pigs/envlookup"
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
		t.Error("value should not be empty", val)
	}
	if err != nil {
		t.Error("error should be nil", err)
	}
}

func TestEnvWithDef(t *testing.T) {
	defval := "Wayne Shorter"
	val, err := envlookup.String("JAZZ_ARTIST", defval)
	if val == defval {
		t.Error("default value should not be set", defval, val)
	}
	if err != nil {
		t.Error("error should be nil", err)
	}
}

func TestEmptyEnv(t *testing.T) {
	val, err := envlookup.String("EMPTY_JAZZ_ARTIST")
	if len(val) != 0 {
		t.Error("value should be empty", val)
	}
	if _, ok := err.(*envlookup.NotFoundError); !ok {
		t.Error("error should be envlookup.NotFoundError", err)
	}
}

func TestEmptyEnvWithDef(t *testing.T) {
	defval := "Wayne Shorter"
	val, err := envlookup.String("EMPTY_JAZZ_ARTIST", defval)
	if val != defval {
		t.Error("default value should be set", defval, val)
	}
	if err != nil {
		t.Error("error should be nil", err)
	}
}

func TestSliceEnv(t *testing.T) {
	val, err := envlookup.Slice("RECORD_LABELS")
	if len(val) == 0 {
		t.Error("value should not be empty", val)
	}
	if err != nil {
		t.Error("error should be nil", err)
	}
}

func TestSliceEnvWithDef(t *testing.T) {
	defval := []string{"Impulse!", "Atlantic"}
	val, err := envlookup.Slice("RECORD_LABELS", defval)

	if reflect.DeepEqual(val, defval) {
		t.Error("default value should not be set", defval, val)
	}
	if err != nil {
		t.Error("error should be nil", err)
	}
}

func TestEmptySliceEnv(t *testing.T) {
	val, err := envlookup.Slice("EMPTY_RECORD_LABELS")
	if len(val) != 0 {
		t.Error("value should be empty", val)
	}
	if _, ok := err.(*envlookup.NotFoundError); !ok {
		t.Error("error should be envlookup.NotFoundError", err)
	}
}

func TestEmptySliceEnvWithDef(t *testing.T) {
	defval := []string{"Impulse!", "Atlantic"}
	val, err := envlookup.Slice("EMPTY_RECORD_LABELS", defval)

	if !reflect.DeepEqual(val, defval) {
		t.Error("default value should be set", defval, val)
	}
	if err != nil {
		t.Error("error should be nil", err)
	}
}

func TestIntEnv(t *testing.T) {
	val, err := envlookup.Int("NO_OF_STUDIO_ALBUMS")
	if val == 0 {
		t.Error("value should not be zero", val)
	}
	if err != nil {
		t.Error("error should be nil", err)
	}
}

func TestIntEnvOrDef(t *testing.T) {
	defval := 1
	val, err := envlookup.Int("NO_OF_STUDIO_ALBUMS", defval)
	if val == defval {
		t.Error("default value should not be set", defval, val)
	}
	if err != nil {
		t.Error("error should be nil", err)
	}
}

func TestEmptyIntEnv(t *testing.T) {
	val, err := envlookup.Int("EMPTY_NO_OF_STUDIO_ALBUMS")
	if val != 0 {
		t.Error("value should be zero", val)
	}
	if _, ok := err.(*envlookup.NotFoundError); !ok {
		t.Error("error should be envlookup.NotFoundError", err)
	}
}

func TestEmptyIntEnvOrDef(t *testing.T) {
	defval := 51
	val, err := envlookup.Int("EMPTY_NO_OF_STUDIO_ALBUMS", defval)
	if val != defval {
		t.Error("default value should be set", defval, val)
	}
	if err != nil {
		t.Error("error should be nil", err)
	}
}

func TestFloat64Env(t *testing.T) {
	val, err := envlookup.Float64("LONGEST_RECORDED_TRACK_FLOAT")
	if val == 0 {
		t.Error("value should not be zero", val)
	}
	if err != nil {
		t.Error("error should be nil", err)
	}
}

func TestFloat64EnvOrDef(t *testing.T) {
	defval := 4.43
	val, err := envlookup.Float64("LONGEST_RECORDED_TRACK_FLOAT", defval)
	if val == defval {
		t.Error("default value should not be set", defval, val)
	}
	if err != nil {
		t.Error("error should be nil", err)
	}
}

func TestFloat64ParseErr(t *testing.T) {
	_, err := envlookup.Float64("LONGEST_RECORDED_TRACK")
	if _, ok := err.(*envlookup.ParseError); !ok {
		t.Error("error should be envlookup.ParseError", err)
	}
}

func TestEmptyFloat64Env(t *testing.T) {
	val, err := envlookup.Float64("EMPTY_LONGEST_RECORDED_TRACK_FLOAT")
	if val != 0 {
		t.Error("value should be zero", val)
	}
	if _, ok := err.(*envlookup.NotFoundError); !ok {
		t.Error("error should be envlookup.NotFoundError", err)
	}
}

func TestEmptyFloat64EnvOrDef(t *testing.T) {
	defval := 4.43
	val, err := envlookup.Float64("EMPTY_LONGEST_RECORDED_TRACK_FLOAT", defval)
	if val != defval {
		t.Error("default value should be set", defval, val)
	}
	if err != nil {
		t.Error("error should be nil", err)
	}
}

func TestDurationEnv(t *testing.T) {
	val, err := envlookup.Duration("LONGEST_RECORDED_TRACK")
	if val == 0 {
		t.Error("value should not be 0", val)
	}
	if err != nil {
		t.Error("error should be nil", err)
	}
}

func TestDurationEnvOrDef(t *testing.T) {
	defval, _ := time.ParseDuration("10h2m")
	val, err := envlookup.Duration("LONGEST_RECORDED_TRACK", defval)
	if val == defval {
		t.Error("default value should not be set", defval, val)
	}
	if err != nil {
		t.Error("error should be nil", err)
	}
}

func TestEmptyDurationEnv(t *testing.T) {
	val, err := envlookup.Duration("EMPTY_LONGEST_RECORDED_TRACK")
	if val != 0 {
		t.Error("value should be false", val)
	}
	if _, ok := err.(*envlookup.NotFoundError); !ok {
		t.Error("error should be envlookup.NotFoundError", err)
	}
}

func TestEmptyDurationEnvOrDef(t *testing.T) {
	defval, _ := time.ParseDuration("10h2m")
	val, err := envlookup.Duration("EMPTY_LONGEST_RECORDED_TRACK", defval)
	if val != defval {
		t.Error("default value should be set", defval, val)
	}
	if err != nil {
		t.Error("error should be nil", err)
	}
}

func TestBoolEnv(t *testing.T) {
	val, err := envlookup.Bool("PLAYED_WITH_MILES_DAVIES")
	if !val {
		t.Error("value should be true", val)
	}
	if err != nil {
		t.Error("error should be nil", err)
	}
}

func TestBoolEnvOrDef(t *testing.T) {
	defval := false
	val, err := envlookup.Bool("PLAYED_WITH_MILES_DAVIES", defval)
	if val == defval {
		t.Error("default value should not be set", defval, val)
	}
	if err != nil {
		t.Error("error should be nil", err)
	}
}

func TestEmptyBoolEnv(t *testing.T) {
	val, err := envlookup.Bool("EMPTY_PLAYED_WITH_MILES_DAVIES")
	if val {
		t.Error("value should be false", val)
	}
	if _, ok := err.(*envlookup.NotFoundError); !ok {
		t.Error("error should be envlookup.NotFoundError", err)
	}
}

func TestEmptyBoolEnvOrDef(t *testing.T) {
	defval := true
	val, err := envlookup.Bool("EMPTY_PLAYED_WITH_MILES_DAVIES", defval)
	if val != defval {
		t.Error("default value should be set", defval, val)
	}
	if err != nil {
		t.Error("error should be nil", err)
	}
}

func TestBoolStrParse(t *testing.T) {
	val, err := envlookup.Bool("PLAYED_WITH_MILES_DAVIES")
	if !val {
		t.Error("value should be true", val)
	}
	if err != nil {
		t.Error("error should be nil", err)
	}

	os.Setenv("PLAYED_WITH_MILES_DAVIES", "1")
	val, err = envlookup.Bool("PLAYED_WITH_MILES_DAVIES")
	if !val {
		t.Error("value should be true", val)
	}
	if err != nil {
		t.Error("error should be nil", err)
	}

	os.Setenv("PLAYED_WITH_MILES_DAVIES", "false")
	val, err = envlookup.Bool("PLAYED_WITH_MILES_DAVIES")
	if val {
		t.Error("value should be false", val)
	}
	if err != nil {
		t.Error("error should be nil", err)
	}

	os.Setenv("PLAYED_WITH_MILES_DAVIES", "0")
	val, err = envlookup.Bool("PLAYED_WITH_MILES_DAVIES")
	if val {
		t.Error("value should be false", val)
	}
	if err != nil {
		t.Error("error should be nil", err)
	}

	os.Setenv("PLAYED_WITH_MILES_DAVIES", "f")
	val, err = envlookup.Bool("PLAYED_WITH_MILES_DAVIES")
	if val {
		t.Error("value should be false", val)
	}
	if _, ok := err.(*envlookup.ParseError); !ok {
		t.Error("error should be envlookup.ParseError", err)
	}
}

func TestIntEnvFormatErr(t *testing.T) {
	os.Setenv("NO_OF_STUDIO_ALBUMS", "ABC")
	defval := 1
	val, err := envlookup.Int("NO_OF_STUDIO_ALBUMS", defval)
	if val == defval {
		t.Error("default value should not be set", defval, val)
	}
	if _, ok := err.(*envlookup.ParseError); !ok {
		t.Error("error should be envlookup.ParseError", err)
	}

	val, err = envlookup.Int("NO_OF_STUDIO_ALBUMS")
	if val != 0 {
		t.Error("zero should be set", val)
	}
	if _, ok := err.(*envlookup.ParseError); !ok {
		t.Error("error should be envlookup.ParseError", err)
	}
}
