package redis

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

/*
NewRedis provides client for redis
*/
func NewRedis() redis.Conn {
	conn, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		log.Fatal(err)
	}

	return conn
}
