package faker

import (
	"math"
	"reflect"
	"time"
)

type Date struct {
	time.Time
}

// Past_ Generate past date It takes two argument yearsAgo and ref date
func (d *Date) Pastʹ(params ...interface{}) time.Time {
	//years, ref := d.validate(params...)

	years := 1
	ref := time.Now()

	types := typeOf(params...)

	if len(types) >= 1 && types[0] == reflect.TypeOf(years) {
		years = params[0].(int)
	}

	if len(types) >= 2 && types[1] == reflect.TypeOf(ref) {
		ref = params[0].(time.Time)
	}

	past := ref.AddDate(-years, 0, 0)
	return past
}

func (d *Date) Futureʹ(params ...interface{}) time.Time {

	years := 1
	ref := time.Now()

	types := typeOf(params...)

	if len(types) >= 1 && types[0] == reflect.TypeOf(years) {
		years = params[0].(int)
	}

	if len(types) >= 2 && types[1] == reflect.TypeOf(ref) {
		ref = params[0].(time.Time)
	}

	future := ref.AddDate(years, 0, 0)
	return future
}

func (d *Date) Betweenʹ(from, to string, params ...interface{}) time.Time {
	layout := "02-02-2006"

	kinds := kindOf(params...)
	if len(kinds) >= 1 && kinds[0] == reflect.String {
		layout = params[0].(string)
	}

	fromTime, err := time.Parse(layout, from)
	if err != nil {
		fromTime = time.Unix(time.Now().Unix()-30, 0)
	}

	toTime, err := time.Parse(layout, to)
	if err != nil {
		toTime = time.Now()
	}

	delta := toTime.Unix() - fromTime.Unix()
	newDate := fromTime.Unix() + int64(random(0, int(delta)))

	return time.Unix(newDate, 0)
}

func (d *Date) Recentʹ(params ...interface{}) time.Time {
	days := 1
	date := time.Now()
	kinds := kindOf(params...)

	if len(kinds) >= 1 && kinds[0] == reflect.Int {
		days = params[0].(int)
	}

	min := int(math.Pow(10, 9))
	if days <= 1 {
		days = 1
	}
	max := days * 24 * 3600 * min
	rnd := random(min, max)

	future := date.UnixNano() - int64(rnd)

	return time.Unix(0, future)
}

func (d *Date) typCtx(options ...interface{}) string {
	typ := "Wide"
	context := false

	types := kindOf(options...)

	if len(types) >= 1 {
		if types[0] == reflect.String {
			v := options[0].(string)
			if v == "Abbr" {
				typ = v
			}
		}
	}

	if len(types) > 1 {
		if types[1] == reflect.Bool {
			v := options[1].(bool)
			if v {
				context = v
			}
		}
	}

	if context {
		typ += "Context"
	}

	return typ
}

func (d *Date) Monthʹ(options ...interface{}) string {
	typ := d.typCtx(options...)
	source := getData("Date", "Month", typ)

	return sample(source)
}

func (d *Date) Weekdayʹ(options ...interface{}) string {
	typ := d.typCtx(options...)
	source := getData("Date", "Weekday", typ)

	return sample(source)
}

// ToJSON Encode name and its fields to JSON.
func (d *Date) ToJSON() (string, error) {
	return ToJSON(d)
}

// ToJSON Encode name and its fields to JSON.
func (d *Date) ToXML() (string, error) {
	return ToJSON(d)
}
