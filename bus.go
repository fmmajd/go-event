//This package handles a simple minimalist event bus, letting handlers observe the events
package goevent

import (
	"sync"
)

var events map[string]*Event
var once sync.Once

func init() {
	once.Do(func(){
		events = make(map[string]*Event)
	})
}

//ByName returns the event specified by the name
//and if the event does not exist, creates a new one and saves it in the bus
func ByName(eventName string) *Event {
	e, exists := events[eventName]
	if !exists {
		e = &Event{name:eventName}
		events[eventName] = e
	}
	return e
}


