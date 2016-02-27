package company

type businessSlogan struct {
	Template []string
	Adjective []string
	Buzz []string
	Noun []string
}

var bsTemplate = []string{
	"#{.Company.BusinessSlogan.Adjective} #{.Company.BusinessSlogan.Buzz} #{.Company.BusinessSlogan.Noun}",
}
var BusinessSlogan = businessSlogan{
	Template: bsTemplate,
	Adjective: bsAdjective,
	Buzz: bsVerb,
	Noun: bsNoun,
}