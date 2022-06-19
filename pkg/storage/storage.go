package storage

type Cuttable struct {
	ID          int    `sql:"id"`
	Name        string `sql:"name"`
	Creator     int    `sql:"creator"`
	ParentID    int    `sql:"parent_id"`
	Description string `sql:"description"`
	License     int    `sql:"license"`
	Data        []byte `sql:"data"`
}

type User struct {
	ID       int    `sql:"id"`
	Name     string `sql:"name"`
	Password string `sql:"password"`
	Salt     string `sql:"salt"`
	Email    string `sql:"email"`
}

type License struct {
	ID   int    `sql:"id"`
	Name string `sql:"name"`
	Text string `sql:"text"`
}

type File struct {
	ID         int    `sql:"id"`
	CuttableID int    `sql:"cuttable_id"`
	Type       int    `sql:"type"`
	Data       []byte `sql:"data"`
}

type CuttableStorage interface {
	GetAll() ([]Cuttable, error)
}

type UserStorage interface {
}
