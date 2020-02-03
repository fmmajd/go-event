[![](https://godoc.org/github.com/fmmajd/goevent?status.svg)](https://godoc.org/github.com/fmmajd/goevent)

# Observer system

this package is a minimal effort to have a simple synchronous/asynchronous event bus

## How to install
always check the last version in case I made a mistake here
```bash
go get github.com/fmmajd/goevent@v0.0.3
``` 

## Simple Usage
functions in this collection can be chain-called
```go
import (
    "github.com/fmmajd/goevent"
    "log"
)

goevent.ByName("my_event_name").
    AddASyncListener(func(payload goevent.Payload) {
        log.Println("event async listener")  
    }).
    AddSyncListener(func(payload goevent.Payload){
        log.Println("event sync listener")
    }).
    Dispatch(goevent.CreatePayload().Set("payload_field_1", 123))
```

the function `ByName` returns the event by name, and if it does not exist, creates it

on any event object, you can call two methods, `AddASyncListener` for adding an async listener(DUH!) and `AddSyncListener` for sync listeners.
both these functions return the event itself, so you can chain as many of those as you want

to dispatch an event, you need a Payload object. Payload is kind of a wrapper for a map and you can add any key, value pairs you need the listener to have access to.
to create a Payload object, call `CreatePayload` function and chain it with any number of `Set(key, value)` functions.

after creating or getting the event and setting the listeners, you can dispatch the event by name anytime you wants:
```go
e := goevent.ByName("previously_chosen_name")
p := goevent.CreatePaylod().Set("a", 1).Set("b", "B").Set("c", true)
e.Dispatch(p)
``` 

