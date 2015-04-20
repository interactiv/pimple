package pimple

type Pimple struct{
  services map[string]func(pimple Pimple)interface{}
}


func New(services map[string]func(pimple Pimple)interface{})*Pimple{
  return &Pimple{services:services}
}
