package settings

type Database struct {
	Name string
	URI  string
}

type Settings struct {
	Database *Database
	Address  string
}
