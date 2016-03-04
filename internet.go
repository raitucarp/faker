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

func (id *InternetDomain) Nameʹ(params ...interface{}) string {
	id.Wordʹ()
	id.Suffixʹ()

	id.Name = strings.ToLower(id.Word + "." + id.Suffix)
	return id.Name

}

func (id *InternetDomain) Suffixʹ(params ...interface{}) string {
	tld := getData("Internet", "DomainSuffix", "Generic")
	name := sample(tld)

	types := kindOf(params...)

	// check if all tld or not
	if len(types) >= 1 && types[0] == reflect.Bool {
		tld = getData("Internet", "DomainSuffix", "All")
		name = sample(tld)
	}

	id.Suffix = name
	return id.Suffix
}

func (id *InternetDomain) Wordʹ(params ...interface{}) string {
	name := Name{}
	firstName := name.FirstNameʹ()

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

func (i *Internet) Avatarʹ() string {
	avatarUri := getData("Internet", "AvatarUri")
	i.Avatar = sample(avatarUri)
	return i.Avatar
}

func (i *Internet) Emailʹ(params ...interface{}) string {
	freeEmail := getData("Internet", "FreeEmail")
	provider := sample(freeEmail)

	name := Name{}
	firstName := name.FirstNameʹ()
	lastName := name.LastNameʹ()

	types := kindOf(params...)

	if len(types) >= 1 && types[0] == reflect.String {
		firstName = params[0].(string)
	}

	if len(types) >= 2 && types[1] == reflect.String {
		lastName = params[1].(string)
	}

	if len(types) >= 3 && types[2] == reflect.String {
		provider = params[2].(string)
	}

	i.Email = Slugify(i.Usernameʹ(firstName, lastName)) + "@" + provider

	return i.Email
}

func (i *Internet) ExampleEmail(params ...interface{}) string {
	exEmail := getData("Internet", "ExampleEmail")
	provider := sample(exEmail)

	params = append(params, provider)
	i.Emailʹ(params...)
	return i.Email
}

func (i *Internet) Usernameʹ(params ...interface{}) string {
	name := Name{}
	firstName := name.FirstNameʹ()
	lastName := name.LastNameʹ()

	types := kindOf(params...)
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

func (i *Internet) Protocolʹ() string {
	protocol := getData("Internet", "Protocol")
	i.Protocol = sample(protocol)
	return i.Protocol
}

func (i *Internet) URLʹ() string {
	i.Protocolʹ()
	i.Domain.Nameʹ()

	return i.Protocol + "://" + i.Domain.Name
}

func (i *Internet) DomainNameʹ() string {
	i.Domain.Nameʹ()
	return i.Domain.Name
}

func (i *Internet) DomainSuffixʹ() string {
	i.Domain.Suffixʹ()
	return i.Domain.Suffix
}

func (i *Internet) DomainWordʹ() string {
	i.Domain.Wordʹ()
	return i.Domain.Word
}

func (i *Internet) IPʹ() string {
	subnet := []string{}
	for i := 0; i < 4; i++ {
		rnd := rand.Intn(256)
		subnet = append(subnet, strconv.Itoa(rnd))
	}

	i.IP = strings.Join(subnet, ".")
	return i.IP
}

func (i *Internet) UserAgentʹ() string {
	return ""
}

func (i *Internet) Colorʹ(params ...interface{}) string {
	base := map[string]int{}
	base["red"] = 0
	base["green"] = 0
	base["blue"] = 0

	rand.Seed(time.Now().Unix() + rand.Int63())

	types := kindOf(params...)
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

func (i *Internet) Macʹ() string {
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
func (i *Internet) Passwordʹ() string {
	return ""
}

// Fake Generate random data for all field
func (i *Internet) Fake() {
	i.Avatarʹ()
	i.Emailʹ()
	i.ExampleEmail()
	i.Usernameʹ()
	i.Protocolʹ()
	i.URLʹ()
	i.DomainNameʹ()
	i.DomainSuffixʹ()
	i.DomainWordʹ()
	i.IPʹ()
	i.UserAgentʹ()
	i.Colorʹ()
	i.Macʹ()
	i.Passwordʹ()
}

// ToJSON Encode name and its fields to JSON.
func (i *Internet) ToJSON() (string, error) {
	return ToJSON(i)
}

// ToXML Encode name and its fields to XML.
func (i *Internet) ToXML() (string, error) {
	return ToJSON(i)
}
