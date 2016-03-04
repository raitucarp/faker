package faker

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix() + int64(rand.Int()))
	return rand.Intn(max-min) + min
}

// UUID generate random uuid.
func UUID() string {
	r := regexp.MustCompile(`[xy]`)
	RFC4122Template := "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx"
	result := r.ReplaceAllStringFunc(RFC4122Template, func(placeholder string) string {
		rnd := random(0, 15)
		var val string
		if placeholder == "x" {
			val = strconv.FormatInt(int64(rnd), 16)
		} else {
			val = strconv.FormatInt(int64(rnd&0x3|0x8), 16)
		}
		return val
	})

	return result
}

// Boolean generate between truthy or falsy.
func Boolean() bool {
	rnd := rand.Int()
	if rnd%2 == 0 {
		return true
	}
	return false
}
