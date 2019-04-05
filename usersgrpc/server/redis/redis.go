package redis

import (
	"github.com/gomodule/redigo/redis"
	"os"
)

var redisPool = &redis.Pool{
	MaxIdle:   10,
	MaxActive: 100,
	Dial: func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", getEnv("REDIS_ADDR", "localhost:8011"))
		if err != nil {
			panic(err.Error())
		}
		return c, err
	},
}

func perform(fn func(c redis.Conn) (interface{}, error)) (interface{}, error) {
	pool := redisPool
	conn := pool.Get()
	defer conn.Close()
	reply, err := fn(conn)
	return reply, err
}

// SET sets value for key. If ttl == 0, pair will be preserved indefinitely.
// Otherwise, ttl is a number of seconds after which the entry will expire and key will be freed.
func SET(k, v string, ttl int64) (interface{}, error) {
	fn := func(c redis.Conn) (interface{}, error) {
		if ttl == 0 {
			reply, err := c.Do("SET", k, v)
			return reply, err
		}
		reply, err := c.Do("SET", k, v, "EX", ttl, "NX")
		return reply, err
	}
	return perform(fn)
}

func GET(k string) (interface{}, error) {
	fn := func(c redis.Conn) (interface{}, error) {
		reply, err := c.Do("GET", k)
		return reply, err
	}
	return perform(fn)
}

func DEL(k string) (interface{}, error) {
	fn := func(c redis.Conn) (interface{}, error) {
		reply, err := c.Do("DEL", k)
		return reply, err
	}
	return perform(fn)
}

// HGET returns the value associated with field in the hash stored at key.
//
// Return value
// Bulk string reply: the value associated with field, or nil when field is not present in the hash or key does not exist.
func HGET(h, f string) (interface{}, error) {
	fn := func(c redis.Conn) (interface{}, error) {
		reply, err := c.Do("HGET", h, f)
		return redis.String(reply, err)
	}
	return perform(fn)
}

// HGETALL returns all fields and values of the hash stored at key. In the returned value, every field name is followed by
// its value, so the length of the reply is twice the size of the hash.
//
// Return value
// Array reply: list of fields and their values stored in the hash, or an empty list when key does not exist.
func HGETALL(h string) (interface{}, error) {
	fn := func(c redis.Conn) (interface{}, error) {
		reply, err := c.Do("HGETALL", h)
		return redis.StringMap(reply, err)
	}
	return perform(fn)
}

// getEnv returns environmental variable called name, or fallback if empty
func getEnv(name string, fallback string) string {
	v, ok := os.LookupEnv(name)
	if ok {
		return v
	}
	return fallback
}