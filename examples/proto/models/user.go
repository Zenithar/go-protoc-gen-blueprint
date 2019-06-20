package proto

// User describes user account entity.
type User struct {
	// Id
	Id string
	// Principal
	Principal string
	// Secret
	Secret string
}

// New returns an initialized entity with default values
func NewUser() *User {
	return &User{}
}

// -----------------------------------------------------------------------------
