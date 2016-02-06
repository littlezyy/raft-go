package raft

type Mutex struct {
	ch chan int
}

func (m *Mutex) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
		return false
	}
}

func (m *Mutex) Lock() {
	<-m.ch
}

func (m *Mutex) UnLock() {
	m.ch <- 0
}

func NewMutex() *Mutex {
	m := &Mutex{}
	m.ch = make(chan int, 1)
	m.ch <- 0
	return m
}
