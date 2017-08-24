package utils

import (
	"errors"
	"sync"
)

func NewRegistry() *Registry {
	return &Registry{
		locker:   new(sync.Mutex),
		services: make(map[string]interface{}),
	}
}

type Registry struct {
	locker   *sync.Mutex
	services map[string]interface{}
}

func (this *Registry) Register(name string, o interface{}) error {
	this.locker.Lock()
	defer this.locker.Unlock()

	if o == nil {
		return errors.New("Register service is nil")
	}

	if _, dup := this.services[name]; dup {
		return nil
	}

	this.services[name] = o
	return nil
}

func (this *Registry) Get(name string) interface{} {
	return this.services[name]
}
