package service

import "Dependency_Injection/repository"

type Service interface {
	Create() string
}

type ServiceImpl struct{}

func (service *ServiceImpl) Create() string {
	return "Create Service"
}

func NewServiceImpl() Service {
	return &ServiceImpl{}
}


type CRUDService struct {
	Repo repository.SampleRepository
}

func NewCRUDService(repository repository.SampleRepository) *CRUDService {
	return &CRUDService{
		Repo: repository,
	}
}