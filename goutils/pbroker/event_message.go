package pbroker

type EventMessageI interface {
	GetEventMessageKey() *uint8
	GetValue() map[string]interface{}
	// GetValue() interface{}
}
type InputEvent struct {
	Key   *uint8                 `json:"key"`
	Value map[string]interface{} `json:"value"`

	// Value interface{} `json:"value"`
}

func (i *InputEvent) GetEventMessageKey() *uint8 {
	return i.Key
}

// func (i *InputEvent) GetValue() interface{} {
// 	return i.Value
// }

func (i *InputEvent) GetValue() map[string]interface{} {
	return i.Value
}
