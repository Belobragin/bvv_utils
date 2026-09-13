package mistake

import (
	"sync"

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

// selects error code upon input error parameter `err`, in case of no match return input parameter `generalErrCode`
func ErrorsSelector(err error, generalErrCode int, errMap map[error]int) int {
	var wg sync.Mutex
	wg.Lock()
	defer wg.Unlock()
	outputErrCode, ok := errMap[err]
	if !ok {
		return generalErrCode
	}
	return outputErrCode
}
