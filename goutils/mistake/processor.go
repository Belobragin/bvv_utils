package mistake

import (
	"go.uber.org/zap"
)

func ErrorsProcessor(
	zapstruct *zap.Logger,
	stopC <-chan struct{},
	errC ...<-chan error,
) {
	agg := make(chan error)
	for _, ch := range errC {
		go func(c <-chan error) {
			for msg := range c {
				agg <- msg
			}
		}(ch)
	}
	for {
		select {
		case <-stopC:
			zapstruct.Info("general shutdown - errors processor")
			return
		case msg := <-agg:
			zapstruct.Error("processor error:", zap.Error(msg))
		}
	}
}
