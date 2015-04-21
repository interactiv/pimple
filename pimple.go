package pimple

import (
	"fmt"

	"sync"
)

type Pimple struct {
	services map[string]func() interface{}
}

// New creates a new container
func New(services ...map[string]func(*Pimple) interface{}) (p *Pimple) {
	p = &Pimple{services: map[string]func() interface{}{}}
	for _, serviceMap := range services {
		for key, factory := range serviceMap {
			p.Set(key, factory)
		}
	}
	return
}

// Value sets a value
func (p *Pimple) Value(key string, value interface{}) *Pimple {
	p.services[key] = func() interface{} {
		return value
	}
	return p
}

// Set sets a factory for a service, mapped by a key.
// The factory is executed only once
// and the is cached
func (p *Pimple) Set(key string, fn func(*Pimple) interface{}) *Pimple {
	once := new(sync.Once)
	var result interface{}
	p.services[key] = func() interface{} {
		once.Do(func() {
			result = fn(p)
		})
		return result
	}
	return p
}

// Get gets a service result by key
func (p *Pimple) Get(key string) (result interface{}) {
	if p.services[key] != nil {
		result = p.services[key]()
	}
	return
}

func (p *Pimple) Exists(key string) bool {
	if p.services[key] != nil {
		return true
	}
	return false
}

// Extend extends(decorates) a service definition.
func (p *Pimple) Extend(serviceName string, definition func(interface{}, *Pimple) interface{}) *Pimple {
	if !p.Exists(serviceName) {
		panic(fmt.Sprintf("Pimple : Trying to extend service %s that doesn't not exist\n", serviceName))
	}
	extendedService := p.services[serviceName]
	return p.Set(serviceName, func(p *Pimple) interface{} {
		return definition(extendedService(), p)
	})
}
