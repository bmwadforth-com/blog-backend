package models

// PersonModel Person model info
// @Description Person information
type PersonModel struct {
	Identifier  string `json:"id"`
	FirstName   string `json:"firstName"`
	MiddleName  string `json:"middleName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
}

// PersonCreateRequest New person model info
// @Description New person request
type PersonCreateRequest struct {
	FirstName   string `json:"firstName"`
	MiddleName  string `json:"middleName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
}
