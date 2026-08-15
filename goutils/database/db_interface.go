package database

type ProjectDbI interface {
	GetVersion() string
	CloseDb() error
	// method ConfigureDb demands environment variables
	// these variables are different for different database types
	ConfigureDb() error
}
