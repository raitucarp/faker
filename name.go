package faker

import (
	"encoding/json"
	"encoding/xml"
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

func (n *Name) Fake() {
	n.FirstName_()
	n.LastName_()
	n.Suffix_()
	n.Prefix_()
	n.Title_()
	// fake job descriptor, area, type.
	n.JobTitle_()
}

// FirstName_ Generate first name.
func (n *Name) FirstName_() string {
	// get first name list
	list := getData("Name", "FirstName")

	// assign to firstName
	n.FirstName = sample(list)
	return n.FirstName
}

// LastName_ Generate last name.
func (name *Name) LastName_() string {
	// get last name list
	list := getData("Name", "LastName")

	// assign to
	name.LastName = sample(list)
	return name.LastName
}

// FindName_ Generate first name and last name.
func (name *Name) FindName_() string {
	firstName := name.FirstName_()
	lastName := name.LastName_()
	return firstName + " " + lastName
}

// Suffix_ Generate suffix name.
func (name *Name) Suffix_() string {
	// get last name list
	list := getData("Name", "Suffix")
	name.Suffix = sample(list)
	return name.Suffix
}

// Prefix_ Generate prefix name.
func (name *Name) Prefix_() string {
	// get last name list
	list := getData("Name", "Prefix")
	name.Prefix = sample(list)
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

// ToJSON Encode name and its fields to JSON.
func (n *Name) ToJSON() (s string, err error) {
	result, err := json.Marshal(n)

	if err != nil {
		return
	}
	return string(result), err
}

// ToJSON Encode name and its fields to JSON.
func (n *Name) ToXML() (s string, err error) {
	result, err := xml.Marshal(n)

	if err != nil {
		return
	}
	return string(result), err
}
