package services

import (
	"k8s.io/client-go/pkg/util/intstr"
)
// FactoryAdapter has a method to work with Controller resources.
type FactoryAdapter interface {
	Watch()
	Sync() error
	Clean() error
}
type ServicePorts struct {
	port, nodePort int32
	targetPort intstr.IntOrString
}
