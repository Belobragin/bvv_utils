package usecase

import (
	"database/sql"

	"go.uber.org/zap"
)

type StandardUseCaseI interface {
	GetDb() *sql.DB
	GetLog() *zap.Logger
	GetPort() string
	GetServiceName() string
	GetUseCors() bool
	GetAllowOrigin() string
	GetMaxAge() int
	ErrCh() chan error
	GetAllStopChan() chan struct{}
	// NewServer() *http.Server
}

type StandardUseCaseRealization struct {
	Log         *zap.Logger
	Db          *sql.DB
	ErrC        chan error
	ServiceName string
	Port        string
	AllStopChan chan struct{}
}

func (s *StandardUseCaseRealization) GetDb() *sql.DB {
	return s.Db
}

func (s *StandardUseCaseRealization) GetLog() *zap.Logger {
	return s.Log
}

func (s *StandardUseCaseRealization) ErrCh() chan error {
	return s.ErrC
}

func (s *StandardUseCaseRealization) GetServiceName() string {
	return s.ServiceName
}

func (s *StandardUseCaseRealization) GetPort() string {
	return s.Port
}

func (s *StandardUseCaseRealization) GetAllStopChan() chan struct{} {
	return s.AllStopChan
}

type StandardCorsRealization struct {
	UseCors     bool
	AllowOrigin string
	MaxAge      int
}

func (s *StandardCorsRealization) GetUseCors() bool {
	return s.UseCors
}

func (s *StandardCorsRealization) GetAllowOrigin() string {
	return s.AllowOrigin
}

func (s *StandardCorsRealization) GetMaxAge() int {
	return s.MaxAge
}
