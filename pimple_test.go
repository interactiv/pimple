package pimple_test

import "testing"
import "github.com/interactiv/expect"
import "github.com/interactiv/pimple"

func TestPimple(t *testing.T){
  type Foo{
    baz int
  }
  type Bar{
    foo Foo
  }
  e:=expect.New(t)
  p:=pimple.New(map[string]{func(*pimple.Pimple)interface{}{
    "foo":func(p *pimple.Pimple){
      return &Foo{baz:1}
    },
    "bar":func(p *pimple.Pimple){
      return &Bar{foo:p.Get("foo").(Foo)}
    }
  })
  bar:=p.Get("bar").(Bar)
  e.Expect(bar.foo.baz).ToEqual(1)
  p.Value("biz","a")
  e.Expect(p.Get("biz).(string)).toEqual("a")
}
