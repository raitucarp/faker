package faker

import (
	"bytes"
	"github.com/raitucarp/faker/locales"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func Fake(s string) (result string, err error) {
	s = strings.Replace(s, "#{", "#{sample ", -1)
	t := template.New("fake")
	t = t.Funcs(template.FuncMap{
		"sample": sample,
	})
	t = t.Delims("#{", "}")
	t, err = t.Parse(s)
	if err != nil {
		return
	}

	var b bytes.Buffer

	err = t.Execute(&b, data)
	if err != nil {
		return
	}

	return b.String(), nil
}

// get data from EN, and return error if not found
func defaultLocaleData(step ...string) (values []string) {
	immutable := reflect.ValueOf(defaultData)
	data := immutable.FieldByName(locales.EN).Interface()

	// field name length
	fieldLength := len(step) - 1

	for index, name := range step {

		immutable := reflect.ValueOf(data)
		val := immutable.FieldByName(name)

		if val.Kind().String() == "invalid" {
			return
		}

		if index == fieldLength {
			return val.Interface().([]string)
		} else {
			data = val.Interface()
		}
	}

	return
}

func getData(field ...string) (values []string) {
	// clone of data
	dat := data

	// field name length
	fieldLength := len(field) - 1

	// so far step
	var step []string

	// iterate field args
	for index, name := range field {
		step = append(step, name)
		immutable := reflect.ValueOf(dat)
		val := immutable.FieldByName(name)

		// try to get data from en locale
		if val.Kind().String() == "invalid" {
			d := defaultLocaleData(step...)

			// final decision, if not found
			// return zero slice
			if len(d) <= 0 {
				return
			} else {
				return d
			}
		}

		if index == fieldLength {
			return val.Interface().([]string)
		} else {
			dat = val.Interface()
		}
	}

	return
}

func sample(list []string) string {
	if len(list) <= 0 {
		return ""
	}

	listLength := len(list)
	nano := time.Now().UTC().UnixNano()
	rand.Seed(nano)
	firstName := rand.Perm(listLength)
	chosen := firstName[0]
	return list[chosen]
}

func replaceSymbolsWithNumber(s string, symbol rune) (result string) {
	for _, v := range s {
		if v == symbol {
			r := rand.Intn(9)
			result += strconv.Itoa(r)
		} else {
			result += string(v)
		}
	}
	return
}

func replaceSymbols(symbol string) (s string) {
	alpha := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i := 0; i < len(symbol); i++ {
		if string(symbol[i]) == "#" {
			s += strconv.Itoa(rand.Intn(9))
		} else if string(symbol[i]) == "?" {
			list := strings.Split(alpha, "")
			s += sample(list)
		} else {
			s += string(symbol[i])
		}
	}
	return s
}
