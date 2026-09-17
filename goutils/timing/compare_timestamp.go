package timing

import (
	"time"

	"github.com/belobragin/bvv_utils/goutils/mistake"
)

func OutputCompareTimestamp(timestamps ...time.Time) (time.Time, error) {
	if len(timestamps) == 0 {
		return time.Time{}, mistake.ErrNoTimestamps
	}

	latest := timestamps[0]
	distinct := false
	for _, timestamp := range timestamps[1:] {
		if !timestamp.Equal(timestamps[0]) {
			distinct = true
		}
		if timestamp.After(latest) {
			latest = timestamp
		}
	}

	if distinct {
		return latest, mistake.ErrDistinctTimestamp
	}

	return latest, nil
}
