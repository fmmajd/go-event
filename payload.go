package goevent

type Payload struct {
	data map[string]interface{}
}

func CreatePayload() *Payload {
	p := Payload{
		data: make(map[string]interface{}),
	}

	return &p
}

func (p *Payload) Set(key string, value interface{}) *Payload{
	p.data[key] = value
	return p
}

func (p Payload) All() map[string]interface{} {
	return p.data
}
