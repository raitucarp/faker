package name

type definitions struct {
	FirstName []string
	LastName  []string
	Prefix    []string
	Suffix    []string
	Title     title
}

func Export() interface{} {
	def := definitions{
		FirstName: FirstName,
		LastName:  LastName,
		Prefix:    Prefix,
		Suffix:    Suffix,
		Title:     Title,
	}

	return def
}

/*type definitions struct {
	FirstName []string
	LastName  []string
	Prefix    []string
	Suffix    []string
	Title     map[string][]string
}

func Definitions() definitions {
	d := definitions{}
	d.FirstName = firstName
	d.LastName = lastName
	d.Prefix = prefix
	d.Suffix = suffix
	d.Title = title
}

func Export() interface{} {
	def := definitions{}
	def.FirstName = firstName
	def.LastName = lastName
	def.Prefix = prefix
	def.Suffix = suffix
	def.Title = title
	return def
}

var Definitions = map[string]interface{}{
	"FirstName": firstName,
	"LastName": lastName,
	"Prefix": prefix,
	"Suffix": suffix,
	"Title": title,
}

type NameMethod struct {}

func (m *NameMethod) FirstName() []string{
	return firstName
}

func (m *NameMethod) LastName() []string {
	return lastName
}

func (m *NameMethod) Prefix() []string {
	return prefix
}

func (m *NameMethod) Suffix() []string {
	return suffix
}

func (m *NameMethod) Title() map[string][]string {
	return title
}*/
