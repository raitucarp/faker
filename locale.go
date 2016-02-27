package faker

import (
	"errors"
	"github.com/raitucarp/faker/locales"
	"reflect"
)

// this is for rollback
var defaultData = locales.Export()
var locale string

// SetLocale Set data to specific locale.
func SetLocale(l string) error {
	locale = l
	immutable := reflect.ValueOf(defaultData)
	val := immutable.FieldByName(locale)

	if val.Kind().String() == "invalid" {
		return errors.New("Locale not available")
	}

	data = val.Interface()
	return nil
}
