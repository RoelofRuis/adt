package ds

import "testing"

func TestEventBus(t *testing.T) {
	eventBus := NewEventBus[string]()

	// Test subscribing and publishing events
	eventBus.Subscribe("string_event", func(data string) {
		if data != "test" {
			t.Errorf("Expected 'test', got %s", data)
		}
	})
	eventBus.Subscribe("string_event", func(data string) {
		if data != "test" {
			t.Errorf("Expected 'test', got %s", data)
		}
	})
	eventBus.Publish("string_event", "test")

	// Test subscribing to multiple events
	eventBus.Subscribe("string_event2", func(data string) {
		if data != "test2" {
			t.Errorf("Expected 'test2', got %s", data)
		}
	})
	eventBus.Publish("string_event2", "test2")
}
