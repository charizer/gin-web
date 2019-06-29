package cache

import (
	"github.com/gomodule/redigo/redis"
	"math/rand"
	"time"
)

type Cache struct {
	pools    []*redis.Pool
	poolSize int
}

func NewCache(cfg []string) *Cache {
	poolSize := len(cfg)
	c := &Cache{
		poolSize: poolSize,
		pools:    make([]*redis.Pool, poolSize),
	}
	for i := 0; i < c.poolSize; i++ {
		addr := cfg[i]
		c.pools[i] = &redis.Pool{
			MaxIdle:     10,
			MaxActive:   200,
			Wait:        true,
			IdleTimeout: 240 * time.Second,
			Dial:        func() (redis.Conn, error) { return redis.DialURL(addr) },
		}
	}
	return c
}

func (c *Cache) pool() *redis.Pool {
	slot := rand.Int() % c.poolSize
	return c.pools[slot]
}

func (c *Cache) SafeDo(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := c.pool().Get()
	defer conn.Close()
	reply, err = conn.Do(commandName, args...)
	return reply, err
}
