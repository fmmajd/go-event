package goevent

import (
	"log"
	"sync"
)

var events map[string]Event
var once sync.Once

func init() {
	log.Println("Initializing events")
	once.Do(func(){
		events = make(map[string]Event)
	})
}

func ByName(eventName string) *Event {
	e, exists := events[eventName]
	if !exists {
		e = Event{name:eventName}
		events[eventName] = e
	}
	return &e
}


