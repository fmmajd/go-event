package goevent

//Payload is kind of a wrapper for a map with string keys and any-type values.
type Payload struct {
	data map[string]interface{}
}

//Creates a new Payload object
func CreatePayload() *Payload {
	p := Payload{
		data: make(map[string]interface{}),
	}

	return &p
}

//Add a key-value pair to the payload and returns the modified payload
func (p *Payload) Set(key string, value interface{}) *Payload{
	p.data[key] = value
	return p
}

//Returns the data of the Payload
func (p Payload) All() map[string]interface{} {
	return p.data
}
