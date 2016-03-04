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

func (n *Name) withGender(params ...interface{}) bool {
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

// FirstName_ Generate first name.
func (name *Name) FirstName_(params ...interface{}) string {
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

// LastName_ Generate last name.
func (name *Name) LastName_(params ...interface{}) string {
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

// FindName_ Generate first name and last name.
func (name *Name) FindName_(params ...interface{}) string {
	var gender interface{}
	firstName := name.FirstName_()
	lastName := name.LastName_()
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
			firstName = name.FirstName_(gender)
			lastName = name.LastName_(gender)
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
		prefix = name.Prefix_(gender)
		if prefix != "" {
			finalName = []string{prefix, firstName, lastName}
			return strings.Join(finalName, " ")
		}
	}

	if r == 1 {
		suffix = name.Suffix_()
		finalName = []string{firstName, lastName, suffix}
		return strings.Join(finalName, " ")
	}

	return strings.Join(finalName, " ")
}

// Suffix_ Generate suffix name.
func (name *Name) Suffix_() string {
	// get last name list
	list := getData("Name", "Suffix")
	name.Suffix = sample(list)
	return name.Suffix
}

// Prefix_ Generate prefix name.
func (name *Name) Prefix_(params ...interface{}) string {
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

// Title_ Generate title name.
func (name *Name) Title_() string {
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

// JobTitle_ Generate job title.
func (name *Name) JobTitle_() string {
	title := []string{
		name.JobDescriptor_(),
		name.JobArea_(),
		name.JobType_(),
	}
	return strings.Join(title, " ")
}

// JobDescriptor_ Generate job descriptor
func (name *Name) JobDescriptor_() string {
	// get last name list
	descriptors := getData("Name", "Title", "Descriptor")
	name.Job.Descriptor = sample(descriptors)
	return name.Job.Descriptor
}

// JobArea_ Generate job area.
func (name *Name) JobArea_() string {
	areas := getData("Name", "Title", "Level")
	name.Job.Area = sample(areas)
	return name.Job.Area
}

// JobType_ Generate job type.
func (name *Name) JobType_() string {
	types := getData("Name", "Title", "Job")
	name.Job.Type = sample(types)
	return name.Job.Type
}

// Fake Generate random data for all field
func (this *Name) Fake() {
	this.FirstName_()
	this.LastName_()
	this.FindName_()
	this.Suffix_()
	this.Prefix_()
	this.Title_()
	this.JobTitle_()
	this.JobDescriptor_()
	this.JobArea_()
	this.JobType_()
}

// ToJSON Encode name and its fields to JSON.
func (n *Name) ToJSON() (string, error) {
	return ToJSON(n)
}

// ToJSON Encode name and its fields to JSON.
func (n *Name) ToXML() (string, error) {
	return ToJSON(n)
}
