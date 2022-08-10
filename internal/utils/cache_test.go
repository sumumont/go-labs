package utils

import (
	"github.com/go-labs/internal/logging"
	"testing"
	"time"
)

func TestBbc(t *testing.T) {
	_ = RegisterCache(FreeCache)
	var inferCache = map[string]int{}
	tm := time.NewTicker(time.Second * 10)
	prefix := "INFER_COUNTS"
	key := "query.CacheKey"
	if _, ok := inferCache[key]; ok {
		//续期
		inferCache[key] = 12
	} else {
		inferCache[key] = 12
		go func() {
			for {
				select {
				case <-tm.C: //更新缓存
					_ = NewCache().Set(prefix, key, 321321, 11)
					inferCache[key] = inferCache[key] - 1

					logging.Info().Msg("cache update")
				default:
					//logging.Info().Interface("inferCache.count",inferCache[key]).Send()
					if inferCache[key] <= 0 {
						logging.Info().Interface("inferCache.count", inferCache[key]).Msg("infer cache stop")
						tm.Stop()
						delete(inferCache, key)
						return
					}
				}
			}
		}()
	}
	_ = NewCache().Set(prefix, key, 321321, 11)
	for {
		rsp := 0
		_ = NewCache().Get(prefix, key, &rsp)
		logging.Info().Interface("cache.value", rsp).Interface("inferCache.count", inferCache[key]).Send()
		time.Sleep(time.Second * 3)
	}
}
