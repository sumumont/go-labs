package utils

const (
	DefaultExpireSeconds = 300
)

type CacheInterface interface {
	// 获取缓存
	Get(prefix string, key string, result interface{}) error
	// 设置缓存
	Set(prefix string, key string, value interface{}, expireSeconds int) error
	// 移除缓存
	Del(prefix string, key string) (affected bool)
}
type CacheBuilder func(param interface{}) CacheInterface

var cacheBuilder CacheBuilder

func RegisterCache(builder CacheBuilder) error {
	cacheBuilder = builder
	return nil
}

func NewCache() CacheInterface {
	return cacheBuilder(nil)
}
func NewCacheParam(param interface{}) CacheInterface {
	return cacheBuilder(param)
}
