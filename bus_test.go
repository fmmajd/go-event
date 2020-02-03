package goevent

import (
	"math/rand"
	"testing"
)

//this test needs to be separated from other cases to actually test INIT state
func TestEmptyEventsAtInitialization(t *testing.T) {
	if len(events) != 0 {
		t.Log(events)
		t.Errorf("Expected the bus to be empty at initialization, got len: %d", len(events))
	}
}

func TestByName(t *testing.T) {
	t.Run("name_checking", func (t *testing.T){
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
			t.Run(name, func (t *testing.T){
				e := ByName(name)
				if e.name != name {
					t.Errorf("Expected event named %s, got %s", name, e.name)
				}
			})
		}
	})
	t.Run("adds_to_bus", func (t *testing.T) {
		events = make(map[string]*Event)
		name := "test_name"
		e := ByName(name)
		if len(events) != 1 {
			t.Errorf("Expected 1 events in the bus, got %d", len(events))
		}
		if e.name != name {
			t.Errorf("Expected bus to contain event named %s, got %s", name, e.name)
		}
	})
	t.Run("no_duplicates", func (t *testing.T) {
		events = make(map[string]*Event)
		name := "test_name"
		ByName(name)
		for i:=0; i<rand.Intn(20); i++ {
			ByName(name)
		}
		if len(events) != 1 {
			t.Errorf("Expected 1 events in the bus, got %d", len(events))
		}
	})
}
