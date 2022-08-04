package redix

import (
	"fmt"
	"github.com/go-redis/redis/v9"
	"os"
)

var Ctl *redis.Client

func init() {
	host := os.Getenv("redis_host")
	port := os.Getenv("redis_port")
	Ctl = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", host, port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
