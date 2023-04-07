package lock

import (
	"sync"
)

var (
	l  = sync.RWMutex{}
	lm = make(map[string]*sync.RWMutex)
)

func GetLock(id string) *sync.RWMutex {
	l.Lock()
	defer l.Unlock()
	_, ok := lm[id]
	if !ok {
		lm[id] = &sync.RWMutex{}
	}
	return lm[id]
}

func Release(id string) {
	l.Lock()
	defer l.Unlock()
	delete(lm, id)
}
