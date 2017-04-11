package broadcast

type broadcast struct {
	c chan broadcast
	v interface{}
}

type Broadcaster struct {
	cc    chan broadcast
	sendc chan<- interface{}
}

type Receiver struct {
	c chan broadcast
}

func NewBroadcaster() *Broadcaster {
	cc := make(chan broadcast, 1)
	sendc := make(chan interface{})
	b := &Broadcaster{
		sendc: sendc,
		cc:    cc,
	}

	go func() {
		for {
			select {
			case v := <-sendc:
				if v == nil {
					b.cc <- broadcast{}
					return
				}
				c := make(chan broadcast, 1)
				nb := broadcast{c: c, v: v}
				b.cc <- nb
				b.cc = c
			}
		}
	}()

	return b
}

func (b *Broadcaster) Register() Receiver {
	return Receiver{b.cc}
}

func (b *Broadcaster) Write(v interface{}) {
	b.sendc <- v
}

func (r *Receiver) Read() interface{} {
	b := <-r.c
	v := b.v
	r.c <- b
	r.c = b.c
	return v
}
