package podnetwork

import (
	corev1 "k8s.io/api/core/v1"
)

// Strategy pattern with the configurator interface
// Every configuration will have its own specific configurator
// Each configurator implements its own logic using the same interface
// with apply and reset methods
type Configurator interface {
	Apply(*corev1.Pod) error
	Delete(*corev1.Pod) error
	Get(*corev1.Pod) error
}

type Configuration struct {
	Configurator Configurator
}

// The method Apply applies a new configuration for pods
func (c *Configuration) Apply(pod *corev1.Pod) error {
	return c.Configurator.Apply(pod)
}

// The method Delete removes configuration applied to pods
func (c *Configuration) Delete(pod *corev1.Pod) error {
	return c.Configurator.Delete(pod)
}

// The method Get brings back the CR/Object instance applied to a pod
func (c *Configuration) Get(pod *corev1.Pod) error {
	return c.Configurator.Delete(pod)
}

type AdditionalNets struct {
}

func (AdditionalNets) Apply(pod *corev1.Pod) error {
	// SecondaryNetwork apply logic here
	return nil
}

func (AdditionalNets) Delete(pod *corev1.Pod) error {
	// SecondaryNetwork delete logic here
	return nil
}

func (AdditionalNets) Get(pod *corev1.Pod) error {
	// SecondaryNetwork Get logic here
	return nil
}
