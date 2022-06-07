package memory

import "github.com/nerdbergev/cuttables/pkg/storage"

func NewCuttableStorage() storage.CuttableStorage {
	return cuttableStorage{
		data: map[int]storage.Cuttable{
			1: {ID: 1, Name: "Dino"},
			2: {ID: 2, Name: "Rocket"},
		},
	}
}

type cuttableStorage struct {
	data map[int]storage.Cuttable
}

func (c cuttableStorage) GetAll() ([]storage.Cuttable, error) {
	cs := make([]storage.Cuttable, 0, len(c.data))
	for _, c := range c.data {
		cs = append(cs, c)
	}
	return cs, nil
}
