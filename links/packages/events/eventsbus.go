package eventsbus

type Event struct {
	Type string
	Data any
}

type EventBus struct {
	bus chan Event
}

func NewEventBus() *EventBus {
	return &EventBus{
		bus: make(chan Event),
	}
}

func (eventBus *EventBus) Publish(event Event) {
	eventBus.bus <- event
}

func (eventBus *EventBus) Subscribe() <-chan Event {
	return eventBus.bus
}
