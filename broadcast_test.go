package broadcast

import (
	"testing"
	"time"
)

var ls []interface{}

func listen(r Receiver) {
	for v := r.Read(); v != nil; v = r.Read() {
		go listen(r)
		ls = append(ls, v)
	}
}

func TestDo(t *testing.T) {
	var b = NewBroadcaster()
	r := b.Register()
	go listen(r)
	for i := 0; i < 10; i++ {
		b.Write(i)
	}
	b.Write(nil)

	time.Sleep(2 * 1e9)
}
