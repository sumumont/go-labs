package redis

import (
	"fmt"
	"github.com/go-labs/internal/logging"
	"testing"
)

func TestRedis(t *testing.T) {
	redis := &Redis{}
	redis.Connect()
	result, err := redis.Keys("token:12:*")
	if err != nil {
		logging.Error(err).Send()
	}
	fmt.Println(result)
}
