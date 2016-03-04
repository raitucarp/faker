package faker

import (
	"math/rand"
	"reflect"
	"strings"
)

// Job
type Job struct {
	Descriptor string `json:"descriptor"`
	Area       string `json:"area"`
	Type       string `json:"type"`
}

// Name
type Name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
	Prefix    string `json:"prefix"`
	Suffix    string `json:"suffix"`
	Job       Job    `json:"job"`
}

func genderData(gender Gender, base string, dataName string) []string {
	var list []string

	if gender == Male {
		dataName = "Male" + dataName
	}

	if gender == Female {
		dataName = "Female" + dataName
	}

	list = getData(base, dataName)
	return list

}

func (name *Name) withGender(params ...interface{}) bool {
	types := typeOf(params...)

	// check type
	if len(types) >= 1 && types[0] == reflect.TypeOf(Male) {
		return true
	}
	return false
}

func getGender(g interface{}) Gender {
	t := typeOf(g)

	var gender Gender
	if t[0] == reflect.TypeOf(Male) {
		gender = g.(Gender)
	} else {
		gender = randomGender()
	}
	return gender
}

func randomGender() Gender {
	gender := []Gender{Male, Female}
	rnd := rand.Intn(len(gender))
	return gender[rnd]
}

// FirstNameʹ Generate first name.
func (name *Name) FirstNameʹ(params ...interface{}) string {
	gender := Male

	var firstNames []string
	if name.withGender(params...) {
		gender = getGender(params[0])
		firstNames = genderData(gender, "Name", "FirstName")
	} else {
		// get first name list
		firstNames = getData("Name", "FirstName")
	}
	// assign to firstName
	name.FirstName = sample(firstNames)
	return name.FirstName
}

// LastNameʹ Generate last name.
func (name *Name) LastNameʹ(params ...interface{}) string {
	gender := Male

	var lastNames []string
	if name.withGender(params...) {
		gender = getGender(params[0])
		lastNames = genderData(gender, "Name", "LastName")
	} else {
		// get first name list
		lastNames = getData("Name", "LastName")
	}

	// get last name list
	// assign to
	name.LastName = sample(lastNames)
	return name.LastName
}

// FindNameʹ Generate first name and last name.
func (name *Name) FindNameʹ(params ...interface{}) string {
	var gender interface{}
	firstName := name.FirstNameʹ()
	lastName := name.LastNameʹ()
	var prefix, suffix string
	r := rand.Intn(8)

	/*prefix := name.Prefix_()
	suffix := name.Suffix_()*/

	kinds := kindOf(params...)
	types := typeOf(params...)

	// gender
	if len(kinds) >= 3 {
		if types[2] == reflect.TypeOf(Male) {
			gender = getGender(gender)
			firstName = name.FirstNameʹ(gender)
			lastName = name.LastNameʹ(gender)
		}
	}

	// set first name if defined in params
	if len(kinds) >= 1 && kinds[0] == reflect.String {
		firstName = params[0].(string)
	}
	// set last name if defined in params
	if len(kinds) >= 2 && kinds[1] == reflect.String {
		lastName = params[1].(string)
	}

	finalName := []string{firstName, lastName}

	if r == 0 && len(params) >= 3 {
		prefix = name.Prefixʹ(gender)
		if prefix != "" {
			finalName = []string{prefix, firstName, lastName}
			return strings.Join(finalName, " ")
		}
	}

	if r == 1 {
		suffix = name.Suffixʹ()
		finalName = []string{firstName, lastName, suffix}
		return strings.Join(finalName, " ")
	}

	return strings.Join(finalName, " ")
}

// Suffixʹ Generate suffix name.
func (name *Name) Suffixʹ() string {
	// get last name list
	list := getData("Name", "Suffix")
	name.Suffix = sample(list)
	return name.Suffix
}

// Prefixʹ Generate prefix name.
func (name *Name) Prefixʹ(params ...interface{}) string {
	var prefix []string
	if name.withGender(params...) {
		gender = getGender(params[0])
		prefix = genderData(gender, "Name", "Prefix")
	} else {
		// get first name list
		prefix = getData("Name", "LastName")
	}

	name.Prefix = sample(prefix)
	return name.Prefix
}

// Titleʹ Generate title name.
func (name *Name) Titleʹ() string {
	descriptor := getData("Name", "Title", "Descriptor")
	level := getData("Name", "Title", "Level")
	job := getData("Name", "Title", "Job")

	separator := " "
	name.Title = strings.Join([]string{
		sample(descriptor),
		sample(level),
		sample(job),
	}, separator)
	return name.Title
}

// JobTitleʹ Generate job title.
func (name *Name) JobTitleʹ() string {
	title := []string{
		name.JobDescriptorʹ(),
		name.JobAreaʹ(),
		name.JobTypeʹ(),
	}
	return strings.Join(title, " ")
}

// JobDescriptorʹ Generate job descriptor
func (name *Name) JobDescriptorʹ() string {
	// get last name list
	descriptors := getData("Name", "Title", "Descriptor")
	name.Job.Descriptor = sample(descriptors)
	return name.Job.Descriptor
}

// JobAreaʹ Generate job area.
func (name *Name) JobAreaʹ() string {
	areas := getData("Name", "Title", "Level")
	name.Job.Area = sample(areas)
	return name.Job.Area
}

// JobTypeʹ Generate job type.
func (name *Name) JobTypeʹ() string {
	types := getData("Name", "Title", "Job")
	name.Job.Type = sample(types)
	return name.Job.Type
}

// Fake Generate random data for all field
func (name *Name) Fake() {
	name.FirstNameʹ()
	name.LastNameʹ()
	name.FindNameʹ()
	name.Suffixʹ()
	name.Prefixʹ()
	name.Titleʹ()
	name.JobTitleʹ()
	name.JobDescriptorʹ()
	name.JobAreaʹ()
	name.JobTypeʹ()
}

// ToJSON Encode name and its fields to JSON.
func (name *Name) ToJSON() (string, error) {
	return ToJSON(name)
}

// ToXML Encode name and its fields to XML.
func (name *Name) ToXML() (string, error) {
	return ToXML(name)
}
