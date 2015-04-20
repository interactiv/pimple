package pimple

import "sync"

type Pimple struct{
  services map[string]func(*Pimple)interface{}
}


func New(services map[string]func(*Pimple)interface{})*Pimple{
  return &Pimple{services:services}
}

func (p *Pimple) Value(key string,value interface{})*Pimple{
  p.services[key] = func(p *Pimple)interface{}{
    return value
  }
  return p
}

func (p *Pimple) Set(key string,fn func(*Pimple)interface{)}*Pimple{
  once :=new(sync.Once)
  p.services[key]= func(p *Pimple)interface{}{
    var result interface{}
    once.Do(func(){
      result = fn(p)
    })
    return result
  }
  return p
)
