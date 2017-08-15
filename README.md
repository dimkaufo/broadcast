### Description
Simple broadcast model with linked channels, based on [broadcasting-values-in-go-with-linked-channels](https://rogpeppe.wordpress.com/2009/12/01/concurrent-idioms-1-broadcasting-values-in-go-with-linked-channels/)

### Usage

So it has two members: **Broadcaster** and **Receiver**

To create a new *Broadcaster* and start to send messages:
``` 
    b := broadcast.NewBroadcaster()
    b.Write("Hi!")
```
To receive data you should register *Receiver* and start to read the incoming data:
```
    go func() {
        r := b.Register()
        for v := r.Read(); v != nil; v = r.Read() {
            fmt.Println(v)
        }
    }   
```
To stop the Receiver's from receiving new data (break the read loop) you should send *nil* to the Broadcaster:
``` 
    b.Write(nil)
```