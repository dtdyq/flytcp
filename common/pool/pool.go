package pool

import (
	"sync"
	"sync/atomic"
)

// NewPool 新建pool.
func NewPool(name string, newFunc func() interface{}) *Pool {
	p := &Pool{
		Name: name,
	}

	p.Pool = &sync.Pool{
		New: func() interface{} {
			atomic.AddUint64(&p.newCnt, 1)
			return newFunc()
		},
	}
	return p
}

// Pool sync.Pool with metrics.
type Pool struct {
	Name   string
	Pool   *sync.Pool
	getCnt uint64
	newCnt uint64
}

// Put adds x to the pool.
func (p *Pool) Put(x interface{}) {
	p.Pool.Put(x)
}

// Get get x from pool.
func (p *Pool) Get() interface{} {
	atomic.AddUint64(&p.getCnt, 1)
	return p.Pool.Get()
}
