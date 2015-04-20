package pimple_test

import "testing"
import "github.com/interactiv/expect"
import "github.com/interactiv/pimple"

func TestPimple(t *testing.T){
  type Foo struct{
    baz int
  }
  type Bar struct{
    foo *Foo
  }
  type Buzz struct{
    string string
  }
  p:=pimple.New(map[string]func(*pimple.Pimple)interface{}{
    "foo":func(p *pimple.Pimple)interface{}{
      return &Foo{baz:1}
    },
    "bar":func(p *pimple.Pimple)interface{}{
      return &Bar{foo:p.Get("foo").(*Foo)}
    },
  })
  bar:=p.Get("bar").(*Bar)
  p.Value("biz","a")
  p.Set("buzz",func(p *pimple.Pimple)interface{}{
    return &Buzz{string:p.Get("biz").(string)}
  })
  e:=expect.New(t)
  e.Expect(bar.foo.baz).ToEqual(1)
  e.Expect(p.Get("buzz").(*Buzz).string).ToEqual("a")
}
