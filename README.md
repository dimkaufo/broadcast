### Discription
Simple broadcast model with linked channels, based on [broadcasting-values-in-go-with-linked-channels](https://rogpeppe.wordpress.com/2009/12/01/concurrent-idioms-1-broadcasting-values-in-go-with-linked-channels/)


### Usage

* Broadcast

        type Broadcaster struct {
            // the newest broadcast channel
        	cc    chan broadcast
            // a channel for sending messages
        	sendc chan<- interface{}
        }

`func NewBroadcaster() Broadcaster` returns a new broadcaster

`func (b Broadcaster) Register() Receiver` returns a Reciever that listens to `b`

`func (b Broadcaster) Write(v interface{})` writes a massage `v` and publishes to all receivers

* Receiver

        type Receiver struct {
        	c chan broadcast
        }

`func (r *Receiver) Read() interface{}` reads a message from broadcast which rigistered to before.
