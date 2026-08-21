package usecase

import (
	"github.com/belobragin/bvv_utils/goutils/config"
	"github.com/belobragin/bvv_utils/goutils/database"

	"go.uber.org/zap"
)

type StandardUseCaseI interface {
	GetDb() *database.ProjectPsqlDb
	GetLog() *zap.Logger
	GetApiPort() string
	GetMetricPort() string
	GetServiceName() string
	GetUseCors() bool
	GetAllowOrigin() string
	GetMaxAge() int
	ErrCh() chan error
	GetAllStopChan() chan struct{}
	// NewServer() *http.Server
}

type StandardUseCaseRealization struct {
	log         *zap.Logger
	db          *database.ProjectPsqlDb
	errC        chan error
	serviceName string
	apiPort     string
	allStopChan chan struct{}
}

func (s *StandardUseCaseRealization) GetDb() *database.ProjectPsqlDb {
	return s.db
}

func (s *StandardUseCaseRealization) GetLog() *zap.Logger {
	return s.log
}

func (s *StandardUseCaseRealization) ErrCh() chan error {
	return s.errC
}

func (s *StandardUseCaseRealization) GetServiceName() string {
	return s.serviceName
}

func (s *StandardUseCaseRealization) GetApiPort() string {
	return s.apiPort
}

func (s *StandardUseCaseRealization) GetAllStopChan() chan struct{} {
	return s.allStopChan
}

func NewStandardUseCase(
	c config.StandardConfigI,
	l *zap.Logger,
	ec chan error,
) (*StandardUseCaseRealization, error) {
	var u = new(StandardUseCaseRealization)
	u.log = l
	u.errC = ec
	u.apiPort = c.GetApiPort()
	u.allStopChan = make(chan struct{}, 1)
	return u, nil
}

type StandardCorsRealization struct {
	useCors     bool
	allowOrigin string
	maxAge      int
}

func (s *StandardCorsRealization) GetUseCors() bool {
	return s.useCors
}

func (s *StandardCorsRealization) GetAllowOrigin() string {
	return s.allowOrigin
}

func (s *StandardCorsRealization) GetMaxAge() int {
	return s.maxAge
}

func NewStandardCors(
	c config.CorsConfigI,
) (*StandardCorsRealization, error) {
	var u = new(StandardCorsRealization)
	u.useCors = c.GetUseCors()
	u.allowOrigin = c.GetAllowOrigin()
	u.maxAge = c.GetMaxAge()
	return u, nil
}
