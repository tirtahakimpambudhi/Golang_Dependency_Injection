//go:build wireinject
// +build wireinject

package injector

import (
	"Dependency_Injection/controller"
	"Dependency_Injection/repository"
	"Dependency_Injection/service"
	"github.com/google/wire"
)

func InitializedService(isError bool) (*service.SimpleService, error) {
	wire.Build(service.NewSimpleService, repository.NewRepository)
	return nil, nil
}

func InitializedDatabaseRepository() *repository.DatabaseRepository {
	wire.Build(repository.NewPostgreSQL, repository.NewMongo, repository.NewDBRepository)
	return nil
}

var booksService = wire.NewSet(repository.NewBooksRepository, service.NewBooksService)
var studentsService = wire.NewSet(repository.NewStudentsRepository, service.NewStudentsService)

func IntializedController() *controller.Controller {
	wire.Build(booksService, studentsService, controller.NewController)
	return nil
}

func IntializedControllerStruct() *controller.Controller {
	wire.Build(booksService, studentsService, wire.Struct(new(controller.Controller), "Books", "Students"))
	return nil
}
func InitializedControllerInterface() *controller.ControllerBindingInterface {
	wire.Build(service.NewServiceImpl, controller.NewControllerBindingInterface)
	return nil
}

func InitializedServiceInterfaceWith() *service. {
	wire.Build(service.New, controller.NewControllerBindingInterface)
	return nil
}
