package behavioral

import "fmt"

type (
	Event struct {
		Data int64
	}

	Observer interface {
		OnNotify(Event)
	}

	Notifier interface {
		Register(Observer)
		Deregister(Observer)
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

func (o *eventObserver) OnNotify(e Event) {
	fmt.Printf("*** Observer %d received: %d\n", o.id, e.Data)
}

func (o *eventNotifier) Register(l Observer) {
	o.observers[l] = struct{}{}
}

func (o *eventNotifier) Deregister(l Observer) {
	delete(o.observers, l)
}

func (o *eventNotifier) Notify(e Event) {
	for p := range o.observers {
		p.OnNotify(e)
	}
}
