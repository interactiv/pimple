package pimple

import "sync"

type Pimple struct{
  services map[string]func()interface{}
}

// New creates a new container
func New(services... map[string]func(*Pimple)interface{})(p *Pimple){
  p= &Pimple{services:map[string]func()interface{}{}}
  for _,serviceMap:=range services{
    for key,factory:=range serviceMap{
      p.Set(key,factory)
    }
  }
  return
}
// Value sets a value
func (p *Pimple) Value(key string,value interface{})*Pimple{
  p.services[key] = func()interface{}{
    return value
  }
  return p
}
// Set sets a factory for a service, mapped by a key.
// The factory is executed only once
// and the is cached
func (p *Pimple) Set(key string,fn func(*Pimple)interface{})*Pimple{
  once :=new(sync.Once)
  p.services[key]= func()interface{}{
    var result interface{}
    once.Do(func(){
      result = fn(p)
    })
    return result
  }
  return p
}
// Get gets a service result by key
func (p *Pimple) Get(key string)(result interface{}){
    if p.services[key]!=nil{
      result = p.services[key]()
    }
    return
}
