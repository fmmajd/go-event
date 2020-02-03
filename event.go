package goevent

type Listener func(payload Payload)

type Event struct {
	name string
	aSyncListeners []Listener
	syncListeners []Listener
}

func (e Event) Name() string {
	return e.name
}

func (e *Event) AddASyncListener(listener Listener) *Event {
	e.aSyncListeners = append(e.aSyncListeners, listener)
	return e
}

func (e Event) ASyncListeners() []Listener {
	return e.aSyncListeners
}

func (e *Event) AddSyncListener(listener Listener) *Event {
	e.syncListeners = append(e.syncListeners, listener)
	return e
}

func (e Event) SyncListeners() []Listener {
	return e.syncListeners
}

func (e Event) Dispatch(payload *Payload) {
	var l Listener
	for _, l = range e.SyncListeners() {
		l(*payload)
	}
	for _, l = range e.ASyncListeners() {
		go l(*payload)
	}
}
