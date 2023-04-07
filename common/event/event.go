package event

import (
	"flytcp/common/gort"
	"sync"
)

type EventType string
type Consumer func(interface{})

var eventConsumers = make(map[EventType][]Consumer)
var ecLock = sync.RWMutex{}

func Pub(event EventType, data interface{}) {
	ecLock.Lock()
	defer ecLock.Unlock()
	cs, ok := eventConsumers[event]
	if !ok {
		eventConsumers[event] = make([]Consumer, 0)
	} else {
		for _, i := range cs {
			gort.QuietGo(func() {
				i(data)
			})
		}
	}
}

func Sub(event EventType, consumer Consumer) {
	ecLock.Lock()
	defer ecLock.Unlock()
	_, ok := eventConsumers[event]
	if !ok {
		eventConsumers[event] = make([]Consumer, 0)
	}
	eventConsumers[event] = append(eventConsumers[event], consumer)
}
