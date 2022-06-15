package storage

type Cuttable struct {
	ID   int
	Name string
}

type CuttableStorage interface {
	GetAll() ([]Cuttable, error)
}

type UserStorage interface {
}
