package faker

import (
	"math"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type InternetDomain struct {
	Name   string
	Suffix string
	Word   string
}

func (id *InternetDomain) Name_(params ...interface{}) string {
	id.Word_()
	id.Suffix_()

	id.Name = strings.ToLower(id.Word + "." + id.Suffix)
	return id.Name

}

func (id *InternetDomain) Suffix_(params ...interface{}) string {
	tld := getData("Internet", "DomainSuffix", "Generic")
	name := sample(tld)

	types := typeof(params...)

	// check if all tld or not
	if len(types) >= 1 && types[0] == reflect.Bool {
		tld = getData("Internet", "DomainSuffix", "All")
		name = sample(tld)
	}

	id.Suffix = name
	return id.Suffix
}

func (id *InternetDomain) Word_(params ...interface{}) string {
	name := Name{}
	firstName := name.FirstName_()

	pattern := `(?i)[\\~#&*{}/:<>?|\"']`
	r := regexp.MustCompile(pattern)

	id.Word = r.ReplaceAllString(firstName, "")
	return strings.ToLower(id.Word)
}

type Internet struct {
	Avatar, Email, Username,
	Protocol, URL, IP, UserAgent, Color, Mac,
	Password string
	Domain InternetDomain
}

func (i *Internet) Avatar_() string {
	avatarUri := getData("Internet", "AvatarUri")
	i.Avatar = sample(avatarUri)
	return i.Avatar
}

func (i *Internet) Email_(params ...interface{}) string {
	freeEmail := getData("Internet", "FreeEmail")
	provider := sample(freeEmail)

	name := Name{}
	firstName := name.FirstName_()
	lastName := name.LastName_()

	types := typeof(params...)

	if len(types) >= 1 && types[0] == reflect.String {
		firstName = params[0].(string)
	}

	if len(types) >= 2 && types[1] == reflect.String {
		lastName = params[1].(string)
	}

	if len(types) >= 3 && types[2] == reflect.String {
		provider = params[2].(string)
	}

	i.Email = Slugify(i.Username_(firstName, lastName)) + "@" + provider

	return i.Email
}

func (i *Internet) ExampleEmail(params ...interface{}) string {
	exEmail := getData("Internet", "ExampleEmail")
	provider := sample(exEmail)

	params = append(params, provider)
	i.Email_(params...)
	return i.Email
}

func (i *Internet) Username_(params ...interface{}) string {
	name := Name{}
	firstName := name.FirstName_()
	lastName := name.LastName_()

	types := typeof(params...)
	if len(types) >= 1 && types[0] == reflect.String {
		firstName = params[0].(string)
	}

	if len(types) >= 2 && types[1] == reflect.String {
		lastName = params[1].(string)
	}

	rnd := rand.Intn(4)
	suffixNumber := strconv.Itoa(rand.Intn(99))
	separator := []string{".", "_"}
	randElement := rand.Intn(len(separator))

	var result string
	switch rnd {
	case 0:
		result = firstName + suffixNumber
	case 1:
		result = firstName + separator[randElement] + lastName
	case 2:
		result = firstName + separator[randElement] + lastName + suffixNumber
	case 3:
		result = suffixNumber + firstName + separator[randElement] + lastName
	}

	rndCase := rand.Intn(2)
	switch rndCase {
	case 0:
		result = strings.ToLower(result)
	case 1:
		result = strings.ToLower(result)
		result = strings.Title(result)
	}

	i.Username = result

	return i.Username
}

func (i *Internet) Protocol_() string {
	protocol := getData("Internet", "Protocol")
	i.Protocol = sample(protocol)
	return i.Protocol
}

func (i *Internet) URL_() string {
	i.Protocol_()
	i.Domain.Name_()

	return i.Protocol + "://" + i.Domain.Name
}

func (i *Internet) DomainName_() string {
	i.Domain.Name_()
	return i.Domain.Name
}

func (i *Internet) DomainSuffix_() string {
	i.Domain.Suffix_()
	return i.Domain.Suffix
}

func (i *Internet) DomainWord_() string {
	i.Domain.Word_()
	return i.Domain.Word
}

func (i *Internet) IP_() string {
	subnet := []string{}
	for i := 0; i < 4; i++ {
		rnd := rand.Intn(256)
		subnet = append(subnet, strconv.Itoa(rnd))
	}

	i.IP = strings.Join(subnet, ".")
	return i.IP
}

func (i *Internet) UserAgent_() string {
	return ""
}

func (i *Internet) Color_(params ...interface{}) string {
	base := map[string]int{}
	base["red"] = 0
	base["green"] = 0
	base["blue"] = 0

	rand.Seed(time.Now().Unix() + rand.Int63())

	types := typeof(params...)
	if len(types) >= 1 && types[0] == reflect.Int {
		base["red"] = params[0].(int)
	}

	if len(types) >= 2 && types[1] == reflect.Int {
		base["green"] = params[1].(int)
	}

	if len(types) >= 3 && types[2] == reflect.Int {
		base["blue"] = params[2].(int)
	}

	red := math.Floor(float64(rand.Intn(256)+base["red"]) / 2)
	green := math.Floor(float64(rand.Intn(256)+base["green"]) / 2)
	blue := math.Floor(float64(rand.Intn(256)+base["blue"]) / 2)

	redHex := strconv.FormatInt(int64(red), 16)
	greenHex := strconv.FormatInt(int64(green), 16)
	blueHex := strconv.FormatInt(int64(blue), 16)

	i.Color = "#" + strings.Join([]string{redHex, greenHex, blueHex}, "")
	return i.Color
}

func (i *Internet) Mac_() string {
	var mac string
	for i := 0; i < 12; i++ {
		mac += strconv.FormatInt(int64(rand.Intn(17)), 16)
		if i%2 == 1 && i != 11 {
			mac += ":"
		}
	}

	i.Mac = mac
	return i.Mac
}

// TODO: password
func (i *Internet) Password_() string {
	return ""
}
