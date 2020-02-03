package goevent

import (
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestEvent_Name(t *testing.T) {
	names := []string{
		"name",
		"$$$$",
		"سلام",
		"______",
		"   ",
		"a b",
		"  AAA",
		"AAA   ",
	}
	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			e := Event{name: name}
			got := e.Name()
			if got != name {
				t.Errorf("Expected event name %s, got %s", name, got)
			}
		})
	}
}

func TestEvent_AddASyncListener(t *testing.T) {
	listener := func(payload Payload) {}
	for i:=0; i<10; i++ {
		t.Run(strconv.Itoa(i)+" listeners", func(t *testing.T){
			e := Event{}
			for j:=0; j<i;j++ {
				e.AddASyncListener(listener)
			}
			listenerCount := len(e.aSyncListeners)
			if listenerCount != i {
				t.Errorf("Expected %d async listeners, got %d", i, len(e.aSyncListeners))
			}
		})
	}
}

func TestEvent_AddSyncListener(t *testing.T) {
	listener := func(payload Payload) {}
	for i:=0; i<10; i++ {
		t.Run(strconv.Itoa(i)+" listeners", func(t *testing.T){
			e := Event{}
			for j:=0; j<i;j++ {
				e.AddSyncListener(listener)
			}
			listenerCount := len(e.syncListeners)
			if listenerCount != i {
				t.Errorf("Expected %d sync listeners, got %d", i, len(e.syncListeners))
			}
		})
	}
}

func TestEvent_ASyncListeners(t *testing.T) {
	listener := func(payload Payload){}
	for i:=0; i<10; i++ {
		t.Run(strconv.Itoa(i)+" listeners", func(t *testing.T){
			listeners := []Listener{}
			for j:=0; j<i; j++ {
				listeners = append(listeners, listener)
			}
			e := Event{}
			e.aSyncListeners = listeners
			if len(e.ASyncListeners()) != i {
				t.Errorf("Expected %d async listeners, got %d", i, len(e.ASyncListeners()))
			}
		})
	}
}

func TestEvent_SyncListeners(t *testing.T) {
	listener := func(payload Payload){}
	for i:=0; i<10; i++ {
		t.Run(strconv.Itoa(i)+" listeners", func(t *testing.T){
			listeners := []Listener{}
			for j:=0; j<i; j++ {
				listeners = append(listeners, listener)
			}
			e := Event{}
			e.syncListeners = listeners
			if len(e.SyncListeners()) != i {
				t.Errorf("Expected %d async listeners, got %d", i, len(e.SyncListeners()))
			}
		})
	}
}

func TestEvent_Dispatch(t *testing.T) {
	t.Run("async", func (t *testing.T) {
		testPayload := PayloadASyncTest{}
		listener := func (payload Payload) {
			testPayload.Inc()
		}
		eventName := "test"
		n := 25
		for i:=0; i<n; i++ {
			ByName(eventName).AddASyncListener(listener)
		}
		ByName(eventName).Dispatch(CreatePayload())
		time.Sleep(time.Second*2)
		if testPayload.val != n {
			t.Errorf("Expected val to become %d, got %d", n, testPayload.val)
		}
	})
	t.Run("sync", func (t *testing.T) {
		testPayload := PayloadSyncTest{}
		listener := func (payload Payload) {
			testPayload.val++
		}
		eventName := "test"
		n := 25
		for i:=0; i<n; i++ {
			ByName(eventName).AddSyncListener(listener)
		}
		ByName(eventName).Dispatch(CreatePayload())
		time.Sleep(time.Second*2)
		if testPayload.val != n {
			t.Errorf("Expected val to become %d, got %d", n, testPayload.val)
		}
	})
}

type PayloadASyncTest struct {
	mux sync.Mutex
	val int
}

type PayloadSyncTest struct {
	val int
}

func (p *PayloadASyncTest) Inc() {
	p.mux.Lock()
	p.val = p.val+1
	p.mux.Unlock()
}