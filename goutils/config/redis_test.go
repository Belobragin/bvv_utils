package config

import "testing"

func TestRedisConfig_GetRedisPassw(t *testing.T) {
	// Create a new instance of RedisConfig
	redisConfig := &RedisConfig{
		RedisPsw: "password123",
	}

	// Call the GetRedisPassw method
	redisPassw := redisConfig.GetRedisPassw()

	// Check if the returned password is correct
	expectedPassw := "password123"
	if redisPassw != expectedPassw {
		t.Errorf("GetRedisPassw() returned unexpected password, expected %s, got %s", expectedPassw, redisPassw)
	}
}
