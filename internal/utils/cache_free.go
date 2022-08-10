package utils

import (
	"encoding/json"
	"github.com/coocood/freecache"
	"github.com/go-labs/internal/logging"
	"sync"
)

var cacheInstance *cache

// 缓存结构
type cache struct {
	realCache *freecache.Cache
}

func RegisterFree() error {
	return RegisterCache(FreeCache)
}

var cacheOne sync.Once

// 获取缓存单例
func FreeCache(param interface{}) CacheInterface {
	if cacheInstance != nil {
		return cacheInstance
	}
	cacheOne.Do(func() {
		realCache := createCache()
		cacheInstance = &cache{
			realCache: realCache,
		}
	})
	return cacheInstance
}

// 缓存初始化
func createCache() *freecache.Cache {
	return freecache.NewCache(100 * 1024 * 1024)
}

func (cache *cache) Get(prefix string, key string, result interface{}) error {
	value, err := cache.realCache.Get(buildKey(prefix, key))
	if err != nil {
		logging.Error(err).Send()
		return err
	}
	err = cache.deserial(value, result)
	if err != nil {
		logging.Error(err).Send()
		return err
	}
	return nil
}
func (cache *cache) Set(prefix string, key string, value interface{}, expireSeconds int) error {
	if expireSeconds <= 0 {
		expireSeconds = DefaultExpireSeconds
	}
	v, err := cache.serial(value)
	if err != nil {
		logging.Error(err).Send()
		return err
	}
	err = cache.realCache.Set(buildKey(prefix, key), v, expireSeconds)
	return err
}

func (cache cache) Del(prefix string, key string) (affected bool) {
	return cache.realCache.Del(buildKey(prefix, key))
}

// 序列化,默认json序列化不行的时候,预留替换,result必须是一个地址
func (cache *cache) deserial(value []byte, result interface{}) error {
	return json.Unmarshal(value, result)
}

// 反序列化,默认json反序列化不行的时候,预留替换
func (cache *cache) serial(value interface{}) ([]byte, error) {
	v, err := json.Marshal(value)
	return v, err
}

// 构建key函数，预留
func buildKey(prefix string, key string) []byte {
	return []byte(prefix + "_" + key)
}
