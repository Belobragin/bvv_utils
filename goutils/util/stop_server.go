package util

import (
	"context"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	FinalCloseTime          = 50 * time.Millisecond
	GracefullShutdownTime   = time.Second * 1
	ServerReadHeaderTimeout = time.Duration(100 * time.Millisecond)
)

type ProdServer struct {
	*http.Server
}

func (p ProdServer) Stop(zapstruct *zap.Logger) {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), GracefullShutdownTime)
	defer cancel()
	if err = p.Shutdown(ctx); err != nil {
		zapstruct.Error("Could not stop server gracefully", zap.Error(err))
		if err = p.Close(); err != nil {
			zapstruct.Error("server close error", zap.Error(err))
		}
		cancel()
	}
	<-ctx.Done()
}

func Halt(
	zapstruct *zap.Logger, message string, err error, sch chan<- struct{}) {
	zapstruct.Error(message, zap.Error(err))
	sch <- struct{}{}
}

func GoodbyeWithGrpc(
	zapstruct *zap.Logger,
	stopChannel chan<- struct{},
	grpcServer *grpc.Server,
	allserver ...*http.Server,
) {
	for _, s := range allserver {
		ProdServer{s}.Stop(zapstruct)
	}
	zapstruct.Info("closed all http servers")
	var c = make(chan struct{})
	ctx, cancel := context.WithTimeout(context.Background(), FinalCloseTime)
	defer cancel()
	defer close(stopChannel)
	if grpcServer != nil {
		go func() {
			grpcServer.GracefulStop()
			zapstruct.Info("grpc server gracefully stopped")
			c <- struct{}{}
		}()
		select {
		case <-c:
			return
		case <-ctx.Done():
			grpcServer.Stop()
			zapstruct.Error("force stopped grpc server, will return now")
			os.Exit(1)
		}
	}
}

func GoodbyeHttp(
	zapstruct *zap.Logger,
	stopChannel chan<- struct{},
	allserver ...*http.Server,
) {
	for _, s := range allserver {
		ProdServer{s}.Stop(zapstruct)
	}
	zapstruct.Info("closed all http servers")
	close(stopChannel)
}
