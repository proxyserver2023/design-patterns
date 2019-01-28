package observer

import (
	"fmt"
	"time"
)

type (
	Event struct {
		Data int64
	}

	Observer interface {
		OnNotify(Event)
	}

	Notifier interface {
		Register(Observer)
		DeRegister(Observer)
		Notify(Event)
	}
)

type (
	eventObserver struct {
		id int
	}

	eventNotifier struct {
		observers map[Observer]struct{}
	}
)

// Observer Implementation
func (o *eventObserver) OnNotify(e Event) {
	fmt.Printf("*** Observer %d received: %d\n", o.id, e.Data)
}

// Notifier Implementation
func (n *eventNotifier) Register(o Observer) {
	n.observers[o] = struct{}{}
}

func (n *eventNotifier) DeRegister(o Observer) {
	delete(n.observers, o)
}

func (n *eventNotifier) Notify(e Event) {
	for o := range n.observers {
		o.OnNotify(e)
	}
}

func Run() {
	n := eventNotifier{
		observers: map[Observer]struct{}{},
	}

	n.Register(&eventObserver{id: 1})
	n.Register(&eventObserver{id: 2})
	n.Register(&eventObserver{id: 3})

	stop := time.NewTimer(time.Second * 10).C
	tick := time.NewTicker(time.Second).C

	for {
		select {
		case <-stop:
			fmt.Println("Done With Observing. I am out.")
			return
		case t := <-tick:
			n.Notify(Event{Data: t.UnixNano()})
		}
	}
}

/*
Output
-----------
*** Observer 1 received: 1548663777657324471
*** Observer 2 received: 1548663777657324471
*** Observer 3 received: 1548663777657324471

*** Observer 1 received: 1548663778657266465
*** Observer 2 received: 1548663778657266465
*** Observer 3 received: 1548663778657266465

*** Observer 1 received: 1548663779657159246
*** Observer 2 received: 1548663779657159246
*** Observer 3 received: 1548663779657159246

*** Observer 1 received: 1548663780657376036
*** Observer 2 received: 1548663780657376036
*** Observer 3 received: 1548663780657376036

*** Observer 2 received: 1548663781657399776
*** Observer 3 received: 1548663781657399776
*** Observer 1 received: 1548663781657399776

*** Observer 1 received: 1548663782657289452
*** Observer 2 received: 1548663782657289452
*** Observer 3 received: 1548663782657289452

*** Observer 2 received: 1548663783657345876
*** Observer 3 received: 1548663783657345876
*** Observer 1 received: 1548663783657345876

*** Observer 1 received: 1548663784657361894
*** Observer 2 received: 1548663784657361894
*** Observer 3 received: 1548663784657361894

*** Observer 1 received: 1548663785657371681
*** Observer 2 received: 1548663785657371681
*** Observer 3 received: 1548663785657371681

---------------------------------------------
Done With Observing. I am out.
------------------------------------------*/
