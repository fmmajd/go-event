package goevent

//Listener could be any function which accepts Payload
type Listener func(payload Payload)


//Event represents the events. the main specifier is the name field, but it's private.
//See ByName for how to generate an Event
type Event struct {
	name string
	aSyncListeners []Listener
	syncListeners []Listener
}

//Returns name of the event
func (e Event) Name() string {
	return e.name
}


//AddASyncListener adds a listener who will be run asynchronously and via goroutines
func (e *Event) AddASyncListener(listener Listener) *Event {
	e.aSyncListeners = append(e.aSyncListeners, listener)
	return e
}

//ASyncListeners returns a list of all the async listeners
func (e Event) ASyncListeners() []Listener {
	return e.aSyncListeners
}

//AddSyncListener adds a synchronous listener who will block the main thread until done
func (e *Event) AddSyncListener(listener Listener) *Event {
	e.syncListeners = append(e.syncListeners, listener)
	return e
}

//SyncListeners returns a list of all the sync listeners
func (e Event) SyncListeners() []Listener {
	return e.syncListeners
}

//Dispatch is the main function of the package.
//It dispatches the event and calls all the listeners
func (e Event) Dispatch(payload *Payload) {
	var l Listener
	for _, l = range e.SyncListeners() {
		l(*payload)
	}
	for _, l = range e.ASyncListeners() {
		go l(*payload)
	}
}
