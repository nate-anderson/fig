package fig

import (
	"errors"
	"testing"
)

type testDriver struct {
	vals map[string]string
}

func (d testDriver) Name() string {
	return "test"
}

func (d testDriver) Get(key string) (string, error) {
	if str, ok := d.vals[key]; ok {
		return str, nil
	}
	return "", ErrConfigNotFound
}

func TestFig(t *testing.T) {
	driver := testDriver{
		vals: map[string]string{
			"A": "a",
			"B": "true",
			"C": "3",
			"F": "1.13",
		},
	}

	config := New(driver)

	t.Run("get string", func(t *testing.T) {
		t.Run("defined", func(t *testing.T) {
			t.Run("GetString returns value and no error", func(t *testing.T) {
				val, err := config.GetString("A")
				if err != nil {
					t.Errorf("unexpected non-nil error for known config key: %s", err)
				}
				if val != "a" {
					t.Errorf("unexpected value %s for config key A", val)
				}
			})

			t.Run("GetStringOr returns value instead of default", func(t *testing.T) {
				val := config.GetStringOr("A", "default")
				if val != "a" {
					t.Errorf("got unexpected value %s from GetStringOr call on key A", val)
				}
			})

			t.Run("MustGetString does not panic", func(t *testing.T) {
				val := config.MustGetString("A")
				if val != "a" {
					t.Errorf("got unexpected value %s from MustGetString call on key A", val)
				}
			})
		})

		t.Run("not defined", func(t *testing.T) {
			t.Run("GetString returns empty string and ErrConfigNotFound", func(t *testing.T) {
				val, err := config.GetString("Z")
				if !errors.Is(err, ErrConfigNotFound) {
					t.Errorf("expected ErrConfigNotFound for unknown config key: got %s", err)
				}
				if val != "" {
					t.Errorf("expected empty value for unknown config key: got %s", val)
				}
			})

			t.Run("GetStringOr returns default value", func(t *testing.T) {
				val := config.GetStringOr("Z", "default")
				if val != "default" {
					t.Errorf("got unexpected value %s from GetStringOr call on unknown key", val)
				}
			})

			t.Run("MustGetString panics", func(t *testing.T) {
				var val string
				defer func() {
					err := recover()
					if err == nil {
						t.Errorf("panicked on nil error, or no panic")
					}
					if val != "" {
						t.Errorf("config val was set despite panic")
					}
				}()
				val = config.MustGetString("Z")
			})
		})
	})

	t.Run("get int", func(t *testing.T) {
		t.Run("defined", func(t *testing.T) {
			t.Run("GetInt returns value and no error", func(t *testing.T) {
				val, err := config.GetInt("C")
				if err != nil {
					t.Errorf("unexpected non-nil error for known config key: %s", err)
				}
				if val != 3 {
					t.Errorf("unexpected value %d for config key C", val)
				}
			})

			t.Run("GetIntOr returns value instead of default", func(t *testing.T) {
				val := config.GetIntOr("C", 10)
				if val != 3 {
					t.Errorf("got unexpected value %d from GetIntOr call on key C", val)
				}
			})

			t.Run("MustGetInt does not panic", func(t *testing.T) {
				val := config.MustGetInt("C")
				if val != 3 {
					t.Errorf("got unexpected value %d from MustGetInt call on key C", val)
				}
			})
		})

		t.Run("not defined", func(t *testing.T) {
			t.Run("GetInt returns empty string and ErrConfigNotFound", func(t *testing.T) {
				val, err := config.GetInt("Z")
				if !errors.Is(err, ErrConfigNotFound) {
					t.Errorf("expected ErrConfigNotFound for unknown config key: got %s", err)
				}
				if val != 0 {
					t.Errorf("expected empty value for unknown config key: got %d", val)
				}
			})

			t.Run("GetIntOr returns default value", func(t *testing.T) {
				val := config.GetIntOr("Z", 10)
				if val != 10 {
					t.Errorf("got unexpected value %d from GetIntOr call on unknown key", val)
				}
			})

			t.Run("MustGetInt panics", func(t *testing.T) {
				var val int
				defer func() {
					err := recover()
					if err == nil {
						t.Errorf("panicked on nil error, or no panic")
					}
					if val != 0 {
						t.Errorf("config val was set despite panic")
					}
				}()
				val = config.MustGetInt("Z")
			})
		})
	})

	t.Run("get int64", func(t *testing.T) {
		t.Run("defined", func(t *testing.T) {
			t.Run("GetInt64 returns value and no error", func(t *testing.T) {
				val, err := config.GetInt64("C")
				if err != nil {
					t.Errorf("unexpected non-nil error for known config key: %s", err)
				}
				if val != 3 {
					t.Errorf("unexpected value %d for config key C", val)
				}
			})

			t.Run("GetInt64Or returns value instead of default", func(t *testing.T) {
				val := config.GetInt64Or("C", 10)
				if val != 3 {
					t.Errorf("got unexpected value %d from GetInt64Or call on key C", val)
				}
			})

			t.Run("MustGetInt64 does not panic", func(t *testing.T) {
				val := config.MustGetInt64("C")
				if val != 3 {
					t.Errorf("got unexpected value %d from MustGetInt64 call on key C", val)
				}
			})
		})

		t.Run("not defined", func(t *testing.T) {
			t.Run("GetInt64 returns empty string and ErrConfigNotFound", func(t *testing.T) {
				val, err := config.GetInt64("Z")
				if !errors.Is(err, ErrConfigNotFound) {
					t.Errorf("expected ErrConfigNotFound for unknown config key: got %s", err)
				}
				if val != 0 {
					t.Errorf("expected empty value for unknown config key: got %d", val)
				}
			})

			t.Run("GetInt64Or returns default value", func(t *testing.T) {
				val := config.GetInt64Or("Z", 10)
				if val != 10 {
					t.Errorf("got unexpected value %d from GetInt64Or call on unknown key", val)
				}
			})

			t.Run("MustGetInt64 panics", func(t *testing.T) {
				var val int64
				defer func() {
					err := recover()
					if err == nil {
						t.Errorf("panicked on nil error, or no panic")
					}
					if val != 0 {
						t.Errorf("config val was set despite panic")
					}
				}()
				val = config.MustGetInt64("Z")
			})
		})
	})

	t.Run("get float64", func(t *testing.T) {
		t.Run("defined", func(t *testing.T) {
			t.Run("GetFloat64 returns value and no error", func(t *testing.T) {
				val, err := config.GetFloat64("F")
				if err != nil {
					t.Errorf("unexpected non-nil error for known config key: %s", err)
				}
				if val != 1.13 {
					t.Errorf("unexpected value %f for config key F", val)
				}
			})

			t.Run("GetFloat64Or returns value instead of default", func(t *testing.T) {
				val := config.GetFloat64Or("F", 10)
				if val != 1.13 {
					t.Errorf("got unexpected value %f from GetFloat64Or call on key F", val)
				}
			})

			t.Run("MustGetFloat64 does not panic", func(t *testing.T) {
				val := config.MustGetFloat64("F")
				if val != 1.13 {
					t.Errorf("got unexpected value %f from MustGetFloat64 call on key F", val)
				}
			})
		})

		t.Run("not defined", func(t *testing.T) {
			t.Run("GetFloat64 returns empty string and ErrConfigNotFound", func(t *testing.T) {
				val, err := config.GetFloat64("Z")
				if !errors.Is(err, ErrConfigNotFound) {
					t.Errorf("expected ErrConfigNotFound for unknown config key: got %s", err)
				}
				if val != 0 {
					t.Errorf("expected empty value for unknown config key: got %f", val)
				}
			})

			t.Run("GetFloat64Or returns default value", func(t *testing.T) {
				val := config.GetFloat64Or("Z", 10.12)
				if val != 10.12 {
					t.Errorf("got unexpected value %f from GetFloat64Or call on unknown key", val)
				}
			})

			t.Run("MustGetFloat64 panics", func(t *testing.T) {
				var val float64
				defer func() {
					err := recover()
					if err == nil {
						t.Errorf("panicked on nil error, or no panic")
					}
					if val != 0 {
						t.Errorf("config val was set despite panic")
					}
				}()
				val = config.MustGetFloat64("Z")
			})
		})
	})

	t.Run("get Bool", func(t *testing.T) {
		t.Run("defined", func(t *testing.T) {
			t.Run("GetBool returns value and no error", func(t *testing.T) {
				val, err := config.GetBool("B")
				if err != nil {
					t.Errorf("unexpected non-nil error for known config key: %s", err)
				}
				if val != true {
					t.Errorf("unexpected value %t for config key B", val)
				}
			})

			t.Run("GetBoolOr returns value instead of default", func(t *testing.T) {
				val := config.GetBoolOr("B", false)
				if val != true {
					t.Errorf("got unexpected value %t from GetBoolOr call on key B", val)
				}
			})

			t.Run("MustGetBool does not panic", func(t *testing.T) {
				val := config.MustGetBool("B")
				if val != true {
					t.Errorf("got unexpected value %t from MustGetBool call on key B", val)
				}
			})
		})

		t.Run("not defined", func(t *testing.T) {
			t.Run("GetBool returns empty string and ErrConfigNotFound", func(t *testing.T) {
				val, err := config.GetBool("Z")
				if !errors.Is(err, ErrConfigNotFound) {
					t.Errorf("expected ErrConfigNotFound for unknown config key: got %s", err)
				}
				if val != false {
					t.Errorf("expected empty value for unknown config key: got %t", val)
				}
			})

			t.Run("GetBoolOr returns default value", func(t *testing.T) {
				val := config.GetBoolOr("Z", false)
				if val != false {
					t.Errorf("got unexpected value %t from GetBoolOr call on unknown key", val)
				}
			})

			t.Run("MustGetBool panics", func(t *testing.T) {
				var val bool
				defer func() {
					err := recover()
					if err == nil {
						t.Errorf("panicked on nil error, or no panic")
					}
					if val != false {
						t.Errorf("config val was set despite panic")
					}
				}()
				val = config.MustGetBool("Z")
			})
		})
	})
}
