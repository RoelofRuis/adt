package ds

// EventBus is a generic type representing a publish/subscribe mechanism for events of type E.
type EventBus[E any] struct {
	subscriptions map[string][]func(E)
}

// NewEventBus creates a new instance of EventBus for events of type E.
func NewEventBus[E any]() *EventBus[E] {
	return &EventBus[E]{
		subscriptions: make(map[string][]func(E)),
	}
}

// Subscribe adds a callback function to the list of subscribers for a given event key.
func (e *EventBus[E]) Subscribe(key string, callback func(E)) {
	e.subscriptions[key] = append(e.subscriptions[key], callback)
}

// Publish triggers all subscribed callback functions for a given event key with the provided data.
func (e *EventBus[E]) Publish(key string, data E) {
	for _, callback := range e.subscriptions[key] {
		callback(data)
	}
}
