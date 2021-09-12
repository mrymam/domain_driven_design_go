package FullName

type FullName struct {
	firstName string
	lastName  string
}

func NewFullName(firstName string, lastName string) FullName {
	return FullName{
		firstName: firstName,
		lastName:  lastName,
	}
}

func (f FullName) GetFullName() string {
	return f.lastName + " " + f.firstName
}

func (f FullName) IsEqual(other FullName) bool {
	return f.firstName == other.firstName && f.lastName == other.lastName
}
