package pbroker

import (
	"sync"

	"github.com/belobragin/bvv_utils/goutils/mistake"
	"github.com/belobragin/bvv_utils/goutils/util"
)

// realization of HandleEventI interface, which matches most use cases
type ProjectEventProcessor struct {
	processor map[util.CatalogEventType]MessageProcessFoo
	sync.Mutex
}

func NewProjectEventProcessor(
	eventRoutesNamber int) (*ProjectEventProcessor, error) {
	var p = new(ProjectEventProcessor)
	c := make(map[util.CatalogEventType]MessageProcessFoo, eventRoutesNamber)
	p.processor = c
	return p, nil
}

// this function does NOT check c.processor - make sure it is not nil !
func (c *ProjectEventProcessor) SetProcessorEvent(
	event util.CatalogEventType,
	processFoo MessageProcessFoo) {
	c.Lock()
	defer c.Unlock()
	c.processor[event] = processFoo
}

func (c *ProjectEventProcessor) ProcessEvent(
	inputEvent EventMessageI,
) (MessageProcessFoo, error) {
	var err error
	eventKey := inputEvent.GetEventMessageKey()
	if eventKey == nil {
		return nil, mistake.ErrInvalidMessage
	}
	result := util.CatalogEventType(*eventKey)
	c.Lock()
	defer c.Unlock()
	if c.processor == nil {
		return nil, mistake.ErrEventProcessorNotInitialized
	}
	f, ok := c.processor[result]
	if !ok {
		err = mistake.ErrMessageTypeUnknown
		return nil, err
	}
	return f, nil
}
