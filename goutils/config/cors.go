package config

type CorsConfigI interface {
	GetUseCors() bool
	GetAllowOrigin() string
	GetMaxAge() int
}
type CorsConfig struct {
	useCors     bool   `conf:"default:false,env:USE_CORS"`
	allowOrigin string `conf:"env:ACCESS_CONTROL_ALLOW_ORIGIN"`
	maxAge      int    `conf:"default:3600,env:ACCESS_CONTROL_MAX_AGE"`
}

func (s *CorsConfig) GetUseCors() bool {
	return s.useCors
}
func (s *CorsConfig) GetAllowOrigin() string {
	return s.allowOrigin
}
func (s *CorsConfig) GetMaxAge() int {
	return s.maxAge
}

func (c *CorsConfig) ValidateCorsConfig() error {
	return nil
}
