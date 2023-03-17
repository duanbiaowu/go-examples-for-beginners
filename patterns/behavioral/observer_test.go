package behavioral

import (
	"math"
	"testing"
)

func Test_Observer(t *testing.T) {
	n := eventNotifier{
		observers: make(map[Observer]struct{}),
	}

	n.Register(&eventObserver{1})
	n.Register(&eventObserver{2})
	n.Register(&eventObserver{3})

	n.Notify(Event{Data: math.MaxInt64})
}
