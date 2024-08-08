package redis

import (
	"github.com/go-labs/internal/logging"
	"github.com/go-redis/redis"
	"time"
)

// Redis cache implement
type Redis struct {
	client *redis.Client
}

// Connect connect to redis
func (r *Redis) Connect() {
	r.client = redis.NewClient(&redis.Options{
		Addr:     "192.168.0.21:30005",
		Password: "Pb9JXiq0r5XO71so",
		DB:       0,
	})
	_, err := r.client.Ping().Result()
	if err != nil {
		logging.Fatal().Err(err).Msg("Could not connected to redis")
	}
	logging.Debug().Msg("Successfully connected to redis")
}

// Get from key
func (r *Redis) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

// Set value with key and expire time
func (r *Redis) Set(key string, val string, expire int) error {
	return r.client.Set(key, val, time.Duration(expire)).Err()
}

func (r *Redis) Keys(key string) ([]string, error) {
	return r.client.Keys(key).Result()
}

func (r *Redis) Scan(cursor uint64, pat string) ([]string, uint64, error) {
	return r.client.Scan(cursor, pat, 20).Result()
}

// Del delete key in redis
func (r *Redis) Del(key string) error {
	return r.client.Del(key).Err()
}

func (r *Redis) HashSet(hk, key string, value interface{}) error {
	return r.client.HSet(hk, key, value).Err()
}

func (r *Redis) HashMSet(hk string, pairs map[string]interface{}) error {
	return r.client.HMSet(hk, pairs).Err()
}

// HashGet from key
func (r *Redis) HashGet(hk, key string) (string, error) {
	return r.client.HGet(hk, key).Result()
}

func (r *Redis) HashGetAll(hk string) (map[string]string, error) {
	return r.client.HGetAll(hk).Result()
}

// HashDel delete key in specify redis's hashtable
func (r *Redis) HashDel(hk, key string) error {
	return r.client.HDel(hk, key).Err()
}

// Increase increase value
func (r *Redis) Increase(key string) error {
	return r.client.Incr(key).Err()
}

// Expire set ttl
func (r *Redis) Expire(key string, dur time.Duration) error {
	return r.client.Expire(key, dur).Err()
}

func (r *Redis) RPush(key string, data []byte) error {
	return r.client.RPush(key, string(data)).Err()
}

func (r *Redis) LLen(key string) (int64, error) {
	return r.client.LLen(key).Result()
}
func (r *Redis) LRange(key string) ([]string, error) {
	return r.client.LRange(key, 0, -1).Result()
}

func (r *Redis) SAdd(key string, member string) error {
	return r.client.SAdd(key, member).Err()
}

func (r *Redis) SRem(key string, member string) error {
	return r.client.SRem(key, member).Err()
}

func (r *Redis) SMembers(key string) ([]string, error) {
	return r.client.SMembers(key).Result()
}
