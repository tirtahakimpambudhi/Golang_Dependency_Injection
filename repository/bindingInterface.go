package repository

import "fmt"

type SampleRepository interface {
	CRUD() string
}

type SampleRepositoryImpl struct {
	DB string
}

func (repository *SampleRepositoryImpl) CRUD() string {
	return fmt.Sprintf("%s Create , Read , Update , Delete",repository.DB)
}


func NewSampleRepository (DBNAME string) *SampleRepositoryImpl {
	return &SampleRepositoryImpl{}
}