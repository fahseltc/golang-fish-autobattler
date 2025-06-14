package environment

import (
	"runtime"
	"time"

	"github.com/google/uuid"
)

// based on https://medium.com/@souravchoudhary0306/implementation-of-event-driven-architecture-in-go-golang-28d9a1c01f91

// Events
// StartSimulationEvent - no data associated
// StopSimulationEvent - no data associated
// FishAttackedEvent
// FishDiedEvent
// GameOverEvent
// NextEncounterEvent
// EncounterDoneEvent
// EnableUiEvent
// DisableUiEvent

type Event struct {
	Type      string
	Timestamp time.Time
	Data      interface{}
}

type FishAttackedEvent struct {
	SourceId uuid.UUID
	Owner    string
	TargetId uuid.UUID
	Type     string
	Damage   int
}
type FishDiedEvent struct {
	Id     uuid.UUID
	Killer uuid.UUID
}

type NextEncounterEvent struct {
	EncounterType string
}

type EventBus struct {
	subscribers map[string][]func(event Event)
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]func(event Event)),
	}
}

func (eb *EventBus) Subscribe(eventType string, handler func(event Event)) {
	pc, _, _, ok := runtime.Caller(1)
	caller := ""
	if ok {
		caller = runtime.FuncForPC(pc).Name()
	}
	ENV.Logger.Info("eventbus", "subscribe", eventType, "caller", caller)
	eb.subscribers[eventType] = append(eb.subscribers[eventType], handler)
}

func (eb *EventBus) GetSubscribers(eventType string) []func(event Event) {
	return eb.subscribers[eventType]
}

func (eb *EventBus) Unsubscribe(eventType string) {
	if _, exists := eb.subscribers[eventType]; exists {
		delete(eb.subscribers, eventType)
	} else {
		ENV.Logger.Error("eventbus", "error", "No subscribers found", "eventType", eventType)
	}
}

// Publish sends an event to all subscribers of a given event type
func (eb *EventBus) Publish(event Event) {
	ENV.Logger.Info("eventbus", "publish", event.Type, "handlersCount", len(eb.subscribers[event.Type]))
	handlers := eb.subscribers[event.Type]
	for _, handler := range handlers {
		handler(event)
	}
}
