package config

type CorsConfig struct {
	UseCors     bool   `conf:"default:false,env:USE_CORS"`
	AllowOrigin string `conf:"env:ACCESS_CONTROL_ALLOW_ORIGIN"`
	MaxAge      int    `conf:"default:3600,env:ACCESS_CONTROL_MAX_AGE"`
}
