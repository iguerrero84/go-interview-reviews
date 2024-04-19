package designpatterns

type Observer interface {
	Update()
}

type Subject struct {
	observers []Observer
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) Notify() {
	for _, o := range s.observers {
		o.Update()
	}
}

type ConcreteObserver struct{}

func (cp *ConcreteObserver) Update() {
	// Do something when subject changes
}
