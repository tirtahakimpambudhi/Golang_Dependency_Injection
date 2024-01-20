package controller

import "Dependency_Injection/service"

type ControllerBindingInterface struct {
	Service service.Service
}

func NewControllerBindingInterface(service service.Service) *ControllerBindingInterface {
	return &ControllerBindingInterface{Service: service}
}
