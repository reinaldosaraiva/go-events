package events

import "time"


type EventInterface interface {	
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	Handle(event EventInterface)
}

type EventDispatcherInterface interface {
	Dispatch(event EventInterface) error
	Register(eventName string, handler EventHandlerInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string) bool
	Clear() error
}