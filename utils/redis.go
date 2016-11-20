package utils

import (
	"fmt"
	"gopkg.in/redis.v5"
	"time"
)

var client *redis.Client

func GetRedisClient() *redis.Client {
	if client == nil {
		panic("redis client should be init before use!")
	}
	return client
}

func RedisNewClient(addr, pwd string, dbindex int) {
	clienttmp := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       dbindex,
	})
	for {
		_, err := client.Ping().Result()
		if err != nil {
			fmt.Println("redis new client err:", err.Error())
			time.Sleep(1 * time.Second)
			continue
		} else {
			break
		}
	}
	client = clienttmp
}
