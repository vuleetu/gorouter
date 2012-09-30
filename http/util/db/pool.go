package db

import (
    "sync"
    "container/list"
)

type Pool struct {
    capcity int
    p chan uint
    m sync.Mutex
    c *list.List
}

type Thing interface{}

type Creation func() Thing

func New(size int, m Creation) *Pool {
    var thing Thing
    c := list.New()
    for i := 0; i < size; i++ {
        thing = m()
        c.PushBack(thing)
    }
    return &Pool{
        capcity: size,
        p: make(chan uint, size),
        c: c}
}

func (p *Pool) Get() Thing {
    p.m.Lock()
    defer p.m.Unlock()
    if p.c.Len() > 0  {
        e := p.c.Front()
        p.p <- 1
        return e.Value
    }
    return nil
}

// block until pool is available
func (p *Pool) SyncGet() Thing {
    p.p <- 1
    e := p.c.Front()
    return e.Value
}

func (p *Pool) Put(thing Thing) bool {
    p.m.Lock()
    defer p.m.Unlock()
    if p.c.Len() < p.capcity  {
        <- p.p
        p.c.PushBack(thing)
        return true
    }
    return false
}

func (p *Pool) SyncPut(thing Thing) {
    <- p.p
    p.c.PushBack(thing)
}
