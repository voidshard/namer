package name

// Namer is something that chooses or randomly generates names.
//
// A Namer chooses names based on tags, where a tag represents some
// category of names (eg. human, elven, dwarvern, french, english etc).
//
// If a naming function is not given a tag then one will be chosen
// randomly internally.
type Namer interface {
	nameChooser

	// Get a list of valid tags
	Tags() []string

	// Choose a name based on a tag
	Tag(string) nameChooser
}

type nameChooser interface {
	// Get a town
	Town() string

	// Get character name & surname
	// Returns name, surname
	Male() (string, string)
	Female() (string, string)

	// Get place name
	Place() string

	// Get River name
	// Returns name, type
	// Type here is a descriptor like 'brook' 'stream' 'creek' etc
	River() (string, string)
}

// New readies a new `Namer` for uh, naming things
func New() (Namer, error) {
	return NewEmbeddedNamer()
}
