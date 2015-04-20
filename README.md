# pimple

[![Build Status](https://travis-ci.org/interactiv/pimple.svg?branch=master)](https://travis-ci.org/interactiv/pimple)


[![GoDoc](https://godoc.org/github.com/interactiv/pimple?status.png)](http://godoc.org/github.com/interactiv/pimple)

author mparaiso <mparaiso@online.fr>

license MIT

copyrights 2015 mparaiso

Pimple is an dependency injection container for Go. Inspired by Pimple from PHP, developers can use a registery that will resolve dependencies and cache the result. It's a generic factory

USAGE

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


