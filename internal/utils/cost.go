package utils

import (
	"github.com/go-labs/internal/logging"
	"time"
)

func TimeCost(key string) func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		if key == "" {
			key = "this"
		}
		logging.Debug().Msgf("for [%s] time cost:%v", key, tc)
	}
}
