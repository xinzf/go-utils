package utils

import (
	"sync/atomic"
)

func NewObjectPool(initSize, maxSize int, fun func() interface{}) *ObjectPool {
	p := &ObjectPool{
		c:      make(chan interface{}, maxSize),
		create: fun,
	}

	if initSize > 0 {
		for i := 0; i < initSize; i++ {
			p.Put(fun())
		}
	}

	return p
}

type ObjectPool struct {
	c      chan interface{}
	create func() interface{}

	hit  int64
	miss int64
}

func (this *ObjectPool) Get() (o interface{}) {
	select {
	case o = <-this.c:
		atomic.AddInt64(&this.hit, 1)
	default:
		o = this.create()
		atomic.AddInt64(&this.miss, 1)
	}
	return
}

func (this *ObjectPool) Put(o interface{}) {
	select {
	case this.c <- o:
	default:
	}
}

func (this *ObjectPool) Hit() int64 {
	return this.hit
}

func (this *ObjectPool) Miss() int64 {
	return this.miss

}
