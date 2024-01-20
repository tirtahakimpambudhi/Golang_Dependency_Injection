package service

import (
	"Dependency_Injection/repository"
	"errors"
)

type SimpleService struct {
	*repository.SimpleRepository
}

func NewSimpleService(repository *repository.SimpleRepository) (*SimpleService, error) {
	if repository.Error {
		return nil, errors.New("Failed Created Service")
	} else {
		return &SimpleService{
			SimpleRepository: repository,
		}, nil
	}
}
