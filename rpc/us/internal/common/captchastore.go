package common

import (
	"github.com/mojocn/base64Captcha"
	"github.com/tal-tech/go-zero/core/stores/redis"
)

type RedisCaptchaStore struct {
	RedisConn *redis.Redis
}

func NewRedisCaptchaStore(rConn *redis.Redis) base64Captcha.Store {
	if rConn == nil {
		return nil
	}
	v := RedisCaptchaStore{
		RedisConn: rConn,
	}
	return &v
}

// Set sets the digits for the captcha id.
func (s *RedisCaptchaStore) Set(id string, value string) error {
	return s.RedisConn.Setex(id, value, 60*2)
}

// Get returns stored digits for the captcha id. Clear indicates
// whether the captcha must be deleted from the store.
func (s *RedisCaptchaStore) Get(id string, clear bool) string {
	value, err := s.RedisConn.Get(id)
	if err != nil {
		return ""
	}
	if clear == true {
		s.RedisConn.Del(id)
	}
	return value
}

//Verify captcha's answer directly
func (s *RedisCaptchaStore) Verify(id, answer string, clear bool) bool {
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
