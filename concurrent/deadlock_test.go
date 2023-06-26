package concurrent

import (
	"sync"
	"testing"
)

func TestDeadlock(t *testing.T) {
	y()
	x()

	m()
	n()
}

var a sync.RWMutex

var b sync.Mutex
var c sync.Mutex

// 死锁 Case1 - 测试下来没死锁
func x() {
	a.RLock()
	defer a.RUnlock()

	y()
}

func y() {
	a.RLock()
	defer a.RUnlock()

	//
}

// 死锁 Case2
func m() {
	b.Lock()
	c.Lock()
}

func n() {
	c.Lock()
	b.Lock()
}
