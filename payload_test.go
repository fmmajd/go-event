package goevent

import (
	"strconv"
	"testing"
)

func TestCreatePayload(t *testing.T) {
	p := CreatePayload()
	t.Run("empty_date", func (t *testing.T) {
		if len(p.data) != 0 {
			t.Errorf("Expected the payload data to be empty, got %d elements", len(p.data))
		}
	})
}

func TestPayload_Set(t *testing.T) {
	t.Run("new_pair_added", func (t *testing.T) {
		p := CreatePayload()
		key := "k1"
		p.Set(key, "")
		if len(p.data) != 1 {
			t.Errorf("Expected payload to have 1 pair of key-value, got %d", len(p.data))
		}
	})
	t.Run("new_key_added", func (t *testing.T) {
		p := CreatePayload()
		key := "k1"
		p.Set(key, "")
		if _, ok := p.data[key]; !ok {
			t.Errorf("Expected key to be set in payload map")
		}
	})
	t.Run("new_value_added", func (t *testing.T) {
		p := CreatePayload()
		key := "k1"
		val := ""
		p.Set(key, val)
		if p.data[key] != val {
			t.Errorf("Expected payload data value to be %v, got %v", val, p.data[key])
		}
	})
	t.Run("different_kinds_of_values_added", func (t *testing.T) {
		p := CreatePayload()
		key := "k1"
		values := []interface{} {
			"",
			nil,
			1,
			"A",
			"a",
			" ",
			"ا",//persian alef
			"ِ",
			"|",
			"|",//line
			0,
			Payload{},
			Event{},
			1.90220,
			true,
			"     ",
			"\t",
		}
		for val := range values {
			p.Set(key, val)
			if p.data[key] != val {
				t.Errorf("Expected payload data value to be %v, got %v", val, p.data[key])
			}
		}
	})
	t.Run("existing_key_updates", func (t *testing.T) {
		p := CreatePayload()
		key := "k1"
		val1 := "A"
		val2 := "B"
		p.Set(key, val1)
		p.Set(key, val2)
		if p.data[key] != val2 {
			t.Errorf("Expected payload data value to update to %v, got %v", val2, p.data[key])
		}
	})
	t.Run("multiple_keys_works", func (t *testing.T) {
		for i:=0; i<10; i++ {
			t.Run(strconv.Itoa(i)+" keys", func (t *testing.T) {
				p := CreatePayload()
				for j:=0;j<i;j++ {
					key := "k"+strconv.Itoa(j)
					v := 34*(j+8)
					p.Set(key, v)
				}
				if len(p.data) != i {
					t.Errorf("Expected %d pairs of data in the paylaod, got %d", i, len(p.data))
				}
				for j:=0;j<i;j++ {
					key := "k"+strconv.Itoa(j)
					if _, ok := p.data[key]; !ok {
						t.Errorf("Expected key %s to exist in the paylaod, found nothing", key)
					}
					v := 34*(j+8)
					if p.data[key] != v {
						t.Errorf("Expected values of %s in payload to be %v, got %v", key, v, p.data[key])
					}
				}
			})
		}
	})
	t.Run("chaining_works", func (t *testing.T) {
		p := CreatePayload()
		p.Set("k1", 1).
			Set("k2", 2).
			Set("k3", 3).
			Set("k4", 4)

		if len(p.data) != 4 {
			t.Errorf("Expected %d chain calls to work, got %d elements instead", 4, len(p.data))
		}
	})
}

func TestPayload_All(t *testing.T) {
	p := CreatePayload().Set("k1", 1).Set("k2", 2)
	if len(p.data) != len(p.All()) {
		t.Errorf("Expected %d arguments returned, got %d", len(p.data), len(p.All()))
	}
}

func TestPayload_Get(t *testing.T) {
	pairs := map[string]interface{} {
		"k1": 91209,
		"k2": nil,
		"k3": "AAAAAAAAA",
		"k4": true,
		"k5": make(chan int),
	}
	payload := CreatePayload()
	for k, v := range pairs {
		payload.Set(k, v)
	}
	for k, v := range pairs {
		t.Run(k, func (t *testing.T) {
			getV, _ := payload.Get(k)
			if getV != v {
				t.Errorf("Expected value %v, got %v", v, getV)
			}
		})
	}
}