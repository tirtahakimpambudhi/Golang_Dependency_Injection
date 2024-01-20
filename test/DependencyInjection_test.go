package test

import (
	"Dependency_Injection/injector"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDepencencyInjection(t *testing.T) {
	simpleService, err := injector.InitializedService(false)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(simpleService.SimpleRepository.Error)
}

func TestMultiBinding(t *testing.T) {
	dbRepo := injector.InitializedDatabaseRepository()

	assert.Equal(t, "POSTGRESQL", dbRepo.DatabasePostgreSQL.Name)
	assert.Equal(t, "MONGODB", dbRepo.DatabaseMongo.Name)
}

func TestBindingInterface(t *testing.T) {
	controller := injector.InitializedControllerInterface()
	t.Log(controller.Service.Create())
}
