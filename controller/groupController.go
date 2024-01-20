package controller

import "Dependency_Injection/service"

type Controller struct {
	Books    *service.BooksService
	Students *service.StudentsService
}

func NewController(books *service.BooksService, students *service.StudentsService) *Controller {
	return &Controller{
		Books:    books,
		Students: students,
	}
}
