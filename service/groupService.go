package service

import "Dependency_Injection/repository"

type BooksService struct {
	repository *repository.BooksRepository
}

func NewBooksService(repository *repository.BooksRepository) *BooksService {
	return &BooksService{repository: repository}
}

type StudentsService struct {
	repository *repository.StudentsRepository
}

func NewStudentsService(repository *repository.StudentsRepository) *StudentsService {
	return &StudentsService{repository: repository}
}
