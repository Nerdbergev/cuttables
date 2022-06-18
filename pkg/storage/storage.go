package storage

type Cuttable struct {
	ID          int    `sql:"cuttable_id"`
	Name        string `sql:"cuttable_name"`
	Creator     int    `sql:"cuttable_creator"`
	ParentID    int    `sql:"cuttable_parent_id"`
	Description string `sql:"cuttable_description"`
	License     int    `sql:"cuttable_license"`
	Data        []byte `sql:"cuttable_data"`
}

type User struct {
	ID       int    `sql:"user_id"`
	Name     string `sql:"user_name"`
	Password string `sql:"user_password"`
	Salt     string `sql:"user_salt"`
	Email    string `sql:"user_email"`
}

type License struct {
	ID   int    `sql:"license_id"`
	Name string `sql:"license_name"`
	Text string `sql:"license_text"`
}

type File struct {
	ID         int    `sql:"file_id"`
	CuttableID int    `sql:"file_cuttable_id"`
	Type       int    `sql:"file_type"`
	Data       []byte `sql:"file_data"`
}

type CuttableStorage interface {
	GetAll() ([]Cuttable, error)
}

type UserStorage interface {
}
