package proto

// Profile describes user preferences.
type Profile struct {
	// Id is an internal database identifier
	Id string
	// Email
	Email string
	// FirstName
	FirstName string
	// LastName
	LastName string
}

// New returns an initialized entity with default values
func NewProfile() *Profile {
	return &Profile{}
}

// -----------------------------------------------------------------------------
