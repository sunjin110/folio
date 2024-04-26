package gdto

// Person https://developers.google.com/people/api/rest/v1/people#Person
type Person struct {
	ResourceName   string          `json:"reourceName"`
	Etag           string          `json:"etag"`
	Names          []*Name         `json:"names"`
	EmailAddresses []*EmailAddress `json:"emailAddresses"`
}

func (p *Person) GetPrimaryName() *Name {
	for _, name := range p.Names {
		if name.Metadata.Primary {
			return name
		}
	}
	return nil
}

func (p *Person) GetPrimaryEmailAddress() *EmailAddress {
	for _, emailAddress := range p.EmailAddresses {
		if emailAddress.Metadata.Primary {
			return emailAddress
		}
	}
	return nil
}

// Name https://developers.google.com/people/api/rest/v1/people#Person.Name
type Name struct {
	Metadata    *FieldMetadata `json:"metadata"`
	DisplayName string         `json:"displayName"`
	FamilyName  string         `json:"familyName"`
	GivenName   string         `json:"givenName"`
}

type EmailAddress struct {
	Metadata    *FieldMetadata `json:"metadata"`
	Value       string         `json:"value"` // email adderss
	Type        string         `json:"type"`  // home, work, other
	DisplayName string         `json:"displayName"`
}

// https://developers.google.com/people/api/rest/v1/people#Person.FieldMetadata.FIELDS
type FieldMetadata struct {
	Primary bool `json:"primary"`
}
