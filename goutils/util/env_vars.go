package util

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const visibilityDefaultTimeout = 1800 * time.Second

// visibilityTimeoutFromEnv returns SQS_VISIBILITY_TIMEOUT in seconds, or the default if unset or zero.
func VisibilityTimeoutFromEnv(envName string) (int, error) {
	raw := os.Getenv(envName)
	if raw == "" {
		return int(visibilityDefaultTimeout / time.Second), nil
	}

	v, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("invalid %s %q: %w", envName, raw, err)
	}
	if v <= 0 {
		return int(visibilityDefaultTimeout / time.Second), nil
	}

	return v, nil
}
