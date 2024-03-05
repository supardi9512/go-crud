package entities

type Patient struct {
	Id           int64
	Name         string `validate:"required"`
	Nik          string `validate:"required" label:"NIK"`
	Gender       string `validate:"required"`
	PlaceOfBirth string `validate:"required" label:"Place of Birth"`
	DateOfBirth  string `validate:"required" label:"Date of Birth"`
	Address      string `validate:"required"`
	PhoneNumber  string `validate:"required" label:"Phone Number"`
}
