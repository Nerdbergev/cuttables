package service

import "github.com/nerdbergev/cuttables/pkg/storage"

type CuttableService interface {
	GetAll() ([]storage.Cuttable, error)
}

func NewCuttableService(storage storage.CuttableStorage) CuttableService {
	return cuttableService{
		storage: storage,
	}
}

type cuttableService struct {
	storage storage.CuttableStorage
}

func (s cuttableService) GetAll() ([]storage.Cuttable, error) {
	return s.storage.GetAll()
}

type UserService interface{}

type userService struct {
	storage storage.UserStorage
}

func NewUserService() UserService {
	return userService{}
}
