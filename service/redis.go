package service

import (
    "github.com/garyburd/redigo/redis"
)

//REDIS client to be shared throughout service
var REDIS *redis.Pool


//InitRedisClient returns a redis client
func InitRedisClient(port, password string) (*redis.Pool, error) {
    redisPool := redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", port)

		if err != nil {
			return nil, err
		}
		return c, err
	}, 10)
    return redisPool, nil
}
