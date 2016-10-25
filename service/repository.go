package service

import (
    "github.com/garyburd/redigo/redis"
)

type repository interface {
    redisGetValue(key string) (string, error)
    redisSetValue(key, value string) error
}

type repoHandler struct{}

func (r *repoHandler) redisSetValue(key, value string) error {
    c := REDIS.Get()
    defer c.Close()
    _, err := c.Do("SET", key, value)
    return err
}

func (r *repoHandler) redisGetValue(key string) (string, error ){
    c := REDIS.Get()
    defer c.Close()
    var value string
    reply, err := redis.Values(c.Do("MGET", key))
    redis.Scan(reply, &value)
    return value, err
}
