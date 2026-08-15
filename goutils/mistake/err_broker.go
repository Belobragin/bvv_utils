package mistake

import "errors"

// general broker mistakes:
var (
	ErrNoNatsServer              = errors.New("nats server not running")
	ErrNoSuchEvent               = errors.New("invalid event key - no such event")
	ErrInvalidMessage            = errors.New("event message invalid")
	ErrInvalidEventHandlerRoute  = errors.New("no handler - route invalid")
	ErrEventConrextFormatInvalid = errors.New("message context field is invalid")
	ErrNoEventHandlerData        = errors.New("no event handler for the route")
	ErrNoEventApiKeyData         = errors.New("no event api key for the route")
	ErrMessageTypeUnknown        = errors.New("message type unknown")
)
