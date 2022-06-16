package storage

type Cuttable struct {
	ID          int
	Name        string
	Creator     int
	Parent      int
	Description string
	License     int
	Data        []byte
}

type User struct {
	UserID   int
	Name     string
	Password string
	Salt     string
	Email    string
}

type License struct {
	LicenseID int
	Name      string
	Text      string
}

type Files struct {
	FileID     int
	CuttableID int
	FileType   int
	Data       []byte
}

type CuttableStorage interface {
	GetAll() ([]Cuttable, error)
}

type UserStorage interface {
}
