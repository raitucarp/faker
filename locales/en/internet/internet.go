package internet

type Definitions struct {
	AvatarUri    []string
	DomainSuffix domainTld
	FreeEmail    []string
	ExampleEmail []string
	Protocol     []string
}

func Export() Definitions {
	def := Definitions{
		AvatarUri:    AvatarUri,
		DomainSuffix: DomainSuffix,
		FreeEmail:    FreeEmail,
		ExampleEmail: ExampleEmail,
		Protocol:     Protocol,
	}
	return def
}
