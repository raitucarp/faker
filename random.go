package faker

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

var seed int64

func StaticSeed(s int64) {
	seed = s
}

func ResetStaticSeed() {
	seed = 0
}

func GetStaticSeed() int64 {
	return seed
}

func random(min, max int) int {
	if seed > 0 {
		rand.Seed(seed)
	} else {
		rand.Seed(time.Now().Unix() + int64(rand.Int()))
	}

	intn := rand.Intn(max-min) + min
	return intn
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
