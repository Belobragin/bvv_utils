package config

type GRPCConfig struct {
	GrpcPort string `conf:"env:GRPC_PORT"`
}
