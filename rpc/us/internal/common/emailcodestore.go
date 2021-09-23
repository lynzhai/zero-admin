package common

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
)

type RedisEmailCodeStore struct {
	RedisConn *redis.Redis
}

func NewRedisEmailCodeStore(rConn *redis.Redis) *RedisEmailCodeStore {
	if rConn == nil {
		return nil
	}
	v := RedisEmailCodeStore{
		RedisConn: rConn,
	}
	return &v
}

// Set sets the digits for the captcha id.
func (s *RedisEmailCodeStore) Set(id string, value string) error {
	return s.RedisConn.Setex(id, value, 60*60 *3)
}


func (s *RedisEmailCodeStore) Get(id string, clear bool) string {
	value, err := s.RedisConn.Get(id)
	if err != nil {
		return ""
	}
	if clear == true {
		s.RedisConn.Del(id)
	}
	return value
}


func (s *RedisEmailCodeStore) Verify(id, answer string, clear bool) bool {
	value, err := s.RedisConn.Get(id)
	if err != nil {
		return false
	}
	if clear == true {
		s.RedisConn.Del(id)
	}
	if value == answer {
		return true
	}

	return false
}

