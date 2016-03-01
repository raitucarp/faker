package faker

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"time"
)

type Date struct {
	time.Time
}

func (d *Date) validateYears(t reflect.Kind, y interface{}) int {
	years := 1
	if t == reflect.String {
		years, _ = strconv.Atoi(y.(string))
	} else {
		years = y.(int)
	}

	return years
}

func (d *Date) validateRef(t reflect.Kind, r interface{}) time.Time {
	ref := time.Now()

	if t == reflect.String {
		ref, _ = time.Parse("02-01-2006", r.(string))
	} else if t == reflect.Struct {
		if v, ok := r.(time.Time); ok {
			ref = v
		}

	} else {
		ref = time.Now()
	}
	return ref
}

func (d *Date) validate(params ...interface{}) (years int, ref time.Time) {
	types := typeof(params...)

	years = 1
	ref = time.Now()

	if len(params) >= 1 {
		years = d.validateYears(types[0], params[0])
	}

	if len(params) > 1 {
		ref = d.validateRef(types[1], params[1])
	}

	return
}

func (d *Date) randomUnix(years int, ref time.Time) int64 {
	// set range
	min := 1000
	max := years * 365 * 24 * 3600 * 1000

	rnd := random(min, max)

	return int64(rnd)
}

// Past_ Generate past date It takes two argument yearsAgo and ref date
func (d *Date) Past_(params ...interface{}) time.Time {
	years, ref := d.validate(params...)
	rnd := d.randomUnix(years, ref)

	// get past
	t := ref.Unix() - int64(rnd)
	past := time.Unix(t, 0)

	return past
}

func (d *Date) Future_(params ...interface{}) time.Time {
	years, ref := d.validate(params...)
	rnd := d.randomUnix(years, ref)

	// get past
	t := ref.Unix() + int64(rnd)
	future := time.Unix(t, 0)

	return future
}

func (d *Date) Between_(from, to, layout string) time.Time {
	fromTime, err := time.Parse(layout, from)
	if err != nil {
		fmt.Println("error go", err)
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

func (d *Date) Recent_(days int) time.Time {
	date := time.Now()

	min := int(math.Pow(10, 9))
	if days <= 1 {
		days = 1
	}
	max := days * 24 * 3600 * min
	rnd := random(min, max)

	fmt.Println(date.UnixNano(), rnd, date.UnixNano()-int64(rnd))

	future := date.UnixNano() - int64(rnd)

	return time.Unix(0, future)
}

func (d *Date) typCtx(options ...interface{}) string {
	typ := "Wide"
	context := false

	types := typeof(options...)

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

func (d *Date) Month_(options ...interface{}) string {
	typ := d.typCtx(options...)
	source := getData("Date", "Month", typ)

	return sample(source)
}

func (d *Date) Weekday_(options ...interface{}) string {
	typ := d.typCtx(options...)
	source := getData("Date", "Weekday", typ)

	return sample(source)
}
