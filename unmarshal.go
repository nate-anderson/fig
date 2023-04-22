package fig

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	configTag   = "fig"
	requiredTag = "required"
	defaultTag  = "default"
)

// Unmarshal configuration from driver(s) into a struct. `dest“ should be a pointer to a struct
// Unmarshal will attempt to populate fields with the `fig` tag. Fields with the tag `required="true"`
// must be populated by a configured driver, else Unmarshal will error.
// Drivers will be used in configured order
// The `default` tag can specify a default value that will be used if no configured driver has the key
// `required="true"` has no effect for fields with a valid default value
func (c Config) Unmarshal(dest interface{}) error {
	refVal := reflect.ValueOf(dest)
	if refVal.Kind() != reflect.Pointer {
		return errors.New("destination in Unmarshal must be a pointer to a struct")
	}

	if refVal.Elem().Kind() != reflect.Struct {
		return errors.New("destination pointer must be to a struct")
	}

	refType := reflect.TypeOf(dest)

	under := refVal.Elem()
	for i := 0; i < under.NumField(); i++ {
		field := under.Field(i)
		fieldType := refType.Elem().Field(i)
		fieldHasBeenSet := false

		// skip unexported fields
		if !fieldType.IsExported() {
			continue
		}

		// check if this field expects to get a value from fig
		configKey, ok := fieldType.Tag.Lookup(configTag)
		if !ok {
			continue
		}

		// try each driver in configured order
		for _, driver := range c.drivers {
			configVal, err := driver.Get(configKey)
			if err != nil {
				// if this driver simply doesn't know this key, try the next one
				if errors.Is(err, ErrConfigNotFound) {
					continue
				}
				return fmt.Errorf("error reading key %s from driver %s: %w", configKey, driver.Name(), err)
			}

			if err = c.setFieldValue(field, fieldType.Type, configKey, configVal); err != nil {
				return fmt.Errorf("failed to unmarshal config key %s (value %s) into field %s: %w", configKey, configVal, fieldType.Name, err)
			}

			fieldHasBeenSet = true
		}

		// if the field wasn't set, check for a default value, then make sure it wasn't a required field
		if !fieldHasBeenSet {
			defaultVal, ok := fieldType.Tag.Lookup(defaultTag)
			if ok {
				err := c.setFieldValue(field, fieldType.Type, configKey, defaultVal)
				if err != nil {
					return fmt.Errorf("failed to unmarshal default value %s for key %s into field %s: %w", defaultVal, configKey, fieldType.Name, err)
				}
				fieldHasBeenSet = true
			} else {
				if requiredVal, ok := fieldType.Tag.Lookup(requiredTag); ok && strings.ToLower(requiredVal) == "true" {
					return fmt.Errorf("required field %s (config key %s) not found in any configured driver", fieldType.Name, configKey)
				}
			}
		}
	}

	return nil
}

// supports same primitive types supported by the Config `getFoo` methods
// int, int64, bool, string, float64
func (c Config) setFieldValue(field reflect.Value, fieldType reflect.Type, key, value string) error {
	ft := fieldType
	isPtr := false
	if fieldType.Kind() == reflect.Pointer {
		isPtr = true
		ft = fieldType.Elem()
	}

	var refVal reflect.Value

	switch ft.Kind() {
	case reflect.String:
		if isPtr {
			refVal = reflect.ValueOf(&value)
		} else {
			refVal = reflect.ValueOf(value)
		}
		field.Set(refVal)
		return nil

	case reflect.Int:
		if parsed, err := strconv.Atoi(value); err != nil {
			return errConfigWrongType(key, value, typeInt)
		} else {
			if isPtr {
				refVal = reflect.ValueOf(&parsed)
			} else {
				refVal = reflect.ValueOf(parsed)
			}
			field.Set(refVal)
			return nil
		}

	case reflect.Int64:
		if parsed, err := strconv.ParseInt(value, 10, 64); err != nil {
			return errConfigWrongType(key, value, typeInt64)
		} else {
			if isPtr {
				refVal = reflect.ValueOf(&parsed)
			} else {
				refVal = reflect.ValueOf(parsed)
			}
			field.Set(refVal)
			return nil
		}

	case reflect.Bool:
		if parsed, err := strconv.ParseBool(value); err != nil {
			return errConfigWrongType(key, value, typeBool)
		} else {
			if isPtr {
				refVal = reflect.ValueOf(&parsed)
			} else {
				refVal = reflect.ValueOf(parsed)
			}
			field.Set(refVal)
			return nil
		}

	case reflect.Float64:
		if parsed, err := strconv.ParseFloat(value, 64); err != nil {
			return errConfigWrongType(key, value, typeFloat64)
		} else {
			if isPtr {
				refVal = reflect.ValueOf(&parsed)
			} else {
				refVal = reflect.ValueOf(parsed)
			}
			field.Set(refVal)
			return nil
		}

	default:
		return fmt.Errorf("fig does not support config fields of type %s: supported values are (*)string, (*)int, (*)int64, (*)float64 and (*)bool", fieldType.Kind().String())
	}

}
