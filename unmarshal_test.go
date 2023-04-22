package fig

import (
	"testing"
)

type optionalTestStruct struct {
	OptionalString    string   `fig:"optional_string"`
	OptionalInt       int      `fig:"optional_int"`
	OptionalInt64     int64    `fig:"optional_int64"`
	OptionalBool      bool     `fig:"optional_bool"`
	OptionalFloat     float64  `fig:"optional_float"`
	OptionalStringPtr *string  `fig:"optional_string"`
	OptionalIntPtr    *int     `fig:"optional_int"`
	OptionalInt64Ptr  *int64   `fig:"optional_int64"`
	OptionalBoolPtr   *bool    `fig:"optional_bool"`
	OptionalFloatPtr  *float64 `fig:"optional_float"`
}

type requiredTestStruct struct {
	RequiredString string  `fig:"string" required:"true"`
	RequiredInt    int     `fig:"int" required:"true"`
	RequiredInt64  int64   `fig:"int64" required:"true"`
	RequiredBool   bool    `fig:"bool" required:"true"`
	RequiredFloat  float64 `fig:"float" required:"true"`
}

type mixedTestStruct struct {
	OptionalString    string   `fig:"optional_string"`
	OptionalInt       int      `fig:"optional_int"`
	OptionalInt64     int64    `fig:"optional_int64"`
	OptionalBool      bool     `fig:"optional_bool"`
	OptionalFloat     float64  `fig:"optional_float"`
	OptionalStringPtr *string  `fig:"optional_string"`
	OptionalIntPtr    *int     `fig:"optional_int"`
	OptionalInt64Ptr  *int64   `fig:"optional_int64"`
	OptionalBoolPtr   *bool    `fig:"optional_bool"`
	OptionalFloatPtr  *float64 `fig:"optional_float"`

	RequiredString    string   `fig:"string" required:"true"`
	RequiredInt       int      `fig:"int" required:"true"`
	RequiredInt64     int64    `fig:"int64" required:"true"`
	RequiredBool      bool     `fig:"bool" required:"true"`
	RequiredFloat     float64  `fig:"float" required:"true"`
	RequiredStringPtr *string  `fig:"string" required:"true"`
	RequiredIntPtr    *int     `fig:"int" required:"true"`
	RequiredInt64Ptr  *int64   `fig:"int64" required:"true"`
	RequiredBoolPtr   *bool    `fig:"bool" required:"true"`
	RequiredFloatPtr  *float64 `fig:"float" required:"true"`
}

type testStructWithDefaults struct {
	OptionalString string `fig:"optional_string" default:"foo"`
	RequiredString string `fig:"required_string" required:"true" default:"bar"`
}

func TestUnmarshal(t *testing.T) {
	t.Run("defined variables", func(t *testing.T) {
		driver := testDriver{
			vals: map[string]string{
				"string":          "string",
				"int":             "10",
				"int64":           "10000",
				"bool":            "true",
				"float":           "4592.1",
				"optional_string": "opt_str",
				"optional_int":    "11",
				"optional_int64":  "2002",
				"optional_bool":   "true",
				"optional_float":  "456.7",
			},
		}

		conf := New(driver)

		var ts mixedTestStruct
		err := conf.Unmarshal(&ts)
		if err != nil {
			t.Errorf("unexpected error from Unmarshal: %s", err)
		}

		t.Run("required are populated", func(t *testing.T) {
			if ts.RequiredString != "string" {
				t.Errorf("required string has incorrect value %s", ts.RequiredString)
			}

			if *ts.RequiredStringPtr != "string" {
				t.Errorf("required *string has incorrect value %s", *ts.RequiredStringPtr)
			}

			if ts.RequiredInt != 10 {
				t.Errorf("required int has incorrect value %d", ts.RequiredInt)
			}

			if *ts.RequiredIntPtr != 10 {
				t.Errorf("required *int has incorrect value %d", *&ts.RequiredInt)
			}

			if ts.RequiredInt64 != 10000 {
				t.Errorf("required int has incorrect value %d", ts.RequiredInt64)
			}

			if *ts.RequiredInt64Ptr != 10000 {
				t.Errorf("required *int has incorrect value %d", *ts.RequiredInt64Ptr)
			}

			if ts.RequiredBool != true {
				t.Errorf("required bool has incorrect value %t", ts.RequiredBool)
			}

			if *ts.RequiredBoolPtr != true {
				t.Errorf("required *bool has incorrect value %t", *ts.RequiredBoolPtr)
			}

			if ts.RequiredFloat != 4592.1 {
				t.Errorf("required float has incorrect value %f", ts.RequiredFloat)
			}

			if *ts.RequiredFloatPtr != 4592.1 {
				t.Errorf("required *float has incorrect value %f", *ts.RequiredFloatPtr)
			}
		})

		t.Run("optional are populated", func(t *testing.T) {
			if ts.OptionalString != "opt_str" {
				t.Errorf("optional string has incorrect value %s", ts.OptionalString)
			}

			if *ts.OptionalStringPtr != "opt_str" {
				t.Errorf("optional *string has incorrect value %s", *ts.OptionalStringPtr)
			}

			if ts.OptionalInt != 11 {
				t.Errorf("optional int has incorrect value %d", ts.OptionalInt)
			}

			if *ts.OptionalIntPtr != 11 {
				t.Errorf("optional *int has incorrect value %d", *ts.OptionalIntPtr)
			}

			if ts.OptionalInt64 != 2002 {
				t.Errorf("optional int64 has incorrect value %d", ts.OptionalInt64)
			}

			if *ts.OptionalInt64Ptr != 2002 {
				t.Errorf("optional *int64 has incorrect value %d", *ts.OptionalInt64Ptr)
			}

			if ts.OptionalBool != true {
				t.Errorf("optional bool has incorrect value %t", ts.OptionalBool)
			}

			if *ts.OptionalBoolPtr != true {
				t.Errorf("optional *bool has incorrect value %t", *ts.OptionalBoolPtr)
			}

			if ts.OptionalFloat != 456.7 {
				t.Errorf("optional float has incorrect value %f", ts.OptionalFloat)
			}

			if *ts.OptionalFloatPtr != 456.7 {
				t.Errorf("optional *float has incorrect value %f", *ts.OptionalFloatPtr)
			}
		})
	})

	t.Run("undefined variables", func(t *testing.T) {
		t.Run("optionals default to zero value", func(t *testing.T) {
			var ts optionalTestStruct
			driver := testDriver{vals: map[string]string{}}
			conf := New(driver)

			err := conf.Unmarshal(&ts)
			if err != nil {
				t.Errorf("config with all optional fields errored on unmarshal: %s", err)
			}

			if ts.OptionalString != "" {
				t.Errorf("unmarshal on optional field resulted in non-zero string value %v", ts.OptionalString)
			}

			if ts.OptionalStringPtr != nil {
				t.Errorf("unmarshal on optional field resulted in non-zero *string value %v", *ts.OptionalStringPtr)
			}

			if ts.OptionalInt != 0 {
				t.Errorf("unmarshal on optional field resulted in non-zero int value %v", ts.OptionalInt)
			}

			if ts.OptionalIntPtr != nil {
				t.Errorf("unmarshal on optional field resulted in non-zero *int value %v", *ts.OptionalIntPtr)
			}

			if ts.OptionalInt64 != 0 {
				t.Errorf("unmarshal on optional field resulted in non-zero int64 value %v", ts.OptionalInt64)
			}

			if ts.OptionalInt64Ptr != nil {
				t.Errorf("unmarshal on optional field resulted in non-zero *int64 value %v", *ts.OptionalInt64Ptr)
			}

			if ts.OptionalBool != false {
				t.Errorf("unmarshal on optional field resulted in non-zero bool value %v", ts.OptionalBool)
			}

			if ts.OptionalBoolPtr != nil {
				t.Errorf("unmarshal on optional field resulted in non-zero bool value %v", *ts.OptionalBoolPtr)
			}

			if ts.OptionalFloat != 0 {
				t.Errorf("unmarshal on optional field resulted in non-zero bool value %v", ts.OptionalFloat)
			}

			if ts.OptionalFloatPtr != nil {
				t.Errorf("unmarshal on optional field resulted in non-zero bool value %v", *ts.OptionalFloatPtr)
			}
		})

		t.Run("required values error when missing", func(t *testing.T) {
			var ts requiredTestStruct
			driver := testDriver{vals: map[string]string{}}
			conf := New(driver)

			err := conf.Unmarshal(&ts)
			if err == nil {
				t.Errorf("missing required fields should error on unmarshal")
			}
		})
	})

	t.Run("default values are populated", func(t *testing.T) {
		ts := testStructWithDefaults{}
		driver := testDriver{vals: map[string]string{}}
		conf := New(driver)
		err := conf.Unmarshal(&ts)
		if err != nil {
			t.Errorf("unexpected error from unmarshal with defaults: %s", err)
		}

		if ts.OptionalString != "foo" {
			t.Errorf("expected default value 'foo' for optional string, got %s", ts.OptionalString)
		}

		if ts.RequiredString != "bar" {
			t.Errorf("expected default value 'bar' for required string, got %s", ts.RequiredString)
		}
	})
}
