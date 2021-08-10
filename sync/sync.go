package syncness

import (
	"sync"
)

type Counter struct {
	mutex sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Increase() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++

}

func (c *Counter) Value() int {
	return c.value
}
