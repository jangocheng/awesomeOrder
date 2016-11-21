package utils

import (
	"fmt"
	"gopkg.in/redis.v5"
	"time"
)

var client *redis.Client = nil

func GetRedisClient() *redis.Client {
	if client == nil {
		panic("redis client should be init before use!")
	}
	return client
}

func RedisNewClient(addr, pwd string, dbindex int) (*redis.Client, error) {
	clienttmp := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       dbindex,
	})
	i := 0
	for {
		if i > 30 {
			return nil, fmt.Errorf("can't ping redis server addr:%v,pwd:%v,dbindex:%v", addr, pwd, dbindex)
		}
		_, err := clienttmp.Ping().Result()
		if err != nil {
			fmt.Println("redis new client err:", err.Error())
			time.Sleep(1 * time.Second)
			i++
			continue
		} else {
			break
		}
	}
	client = clienttmp
	return clienttmp, nil
}
