package broadcast

import "sync"

type Broadcaster interface {
	Register() Receiver
	Write(v interface{})
}

type Receiver interface {
	Read() interface{}
}

type broadcast struct {
	c chan broadcast
	v interface{}
}

type broadcaster struct {
	cc    chan broadcast
	sendc chan<- interface{}
	mut   *sync.RWMutex
}

type receiver struct {
	c chan broadcast
}

var (
	// Verify that broadcaster implements Broadcaster
	_ Broadcaster = (*broadcaster)(nil)

	// Verify that receiver implements Receiver
	_ Receiver = (*receiver)(nil)
)

func NewBroadcaster() *broadcaster {
	cc := make(chan broadcast, 1)
	sendc := make(chan interface{})
	b := &broadcaster{
		sendc: sendc,
		cc:    cc,
		mut:   &sync.RWMutex{},
	}

	go func() {
		for {
			v := <-sendc
			if v == nil {
				b.cc <- broadcast{}
				return
			}
			c := make(chan broadcast, 1)
			nb := broadcast{c: c, v: v}
			b.cc <- nb

			b.mut.Lock()
			b.cc = c
			b.mut.Unlock()
		}
	}()

	return b
}

func (b *broadcaster) Register() Receiver {
	b.mut.RLock()
	defer b.mut.RUnlock()
	return &receiver{b.cc}
}

func (b *broadcaster) Write(v interface{}) {
	b.sendc <- v
}

func (r *receiver) Read() interface{} {
	b := <-r.c
	v := b.v
	r.c <- b
	r.c = b.c
	return v
}
