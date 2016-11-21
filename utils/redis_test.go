package utils

import (
	"testing"
	"fmt"
	"gopkg.in/redis.v5"
)

func Init() (string, string, int) {
	if _, err := ParseConfig(yamlpath); err != nil {
		panic(err)
	}
	testmap := AwesomeConfig["redis_test"].(CONF)
	return testmap["addr"].(string), testmap["pwd"].(string), testmap["db"].(int)
}

func TestRedis(t *testing.T) {
	addr, pwd, dbindex := Init()
	if _,err:=RedisNewClient(addr, pwd, dbindex);err!=nil{
		t.Error(err)
		return
	}
	GetRedisClient()
	client.FlushDb()
	if strkey:=client.Get("testkey").Val();strkey!=""{
		t.Error("get(testkey) not nil")
		return
	}
	fmt.Println("---------test redis strings---------")
	client.Set("testkey","test redis",0)
	if client.Get("testkey").Val()!="test redis" {
		t.Error("get(testkey) != test redis")
		return
	}
	client.Set("view:blogid",1,0)
	client.Incr("view:blogid")
	if view,_:=client.Get("view:blogid").Int64();view!=int64(2){
		t.Error("view !=2 ,but %v",view)
		return
	}
	client.Incr("testkey")
	client.IncrBy("testkey",15)
	fmt.Println(client.Get("testkey").Val())
	client.Append("view:blogid","append")
	fmt.Println(client.Get("view:blogid").Val())
	client.Append("helloworld","hello world!")
	fmt.Println(client.GetRange("helloworld",3,7).String())
	client.SetRange("helloworld",0,"H")
	fmt.Println(client.GetRange("helloworld",0,-1).String())
	fmt.Println("---------test redis list---------")
	client.RPush("list-key","last")
	client.LPush("list-key","first")
	client.RPush("list-key","'new last'")
	fmt.Println(client.LRange("list-key",0,-1).String())
	client.LPop("list-key")
	client.LPop("list-key")
	fmt.Println(client.LRange("list-key",0,-1).String())
	client.RPush("list-key","a","b","c","d")
	client.LTrim("list-key",2,-1)
	fmt.Println(client.LRange("list-key",0,-1).String())
	client.RPopLPush("list-key","list-key-copy")
	fmt.Println(client.LRange("list-key",0,-1).String())
	fmt.Println(client.LRange("list-key-copy",0,-1).String())

	fmt.Println("---------test redis set---------")
	client.SAdd("set-key","a","b","c")
	fmt.Println(client.SCard("set-key").String())
	fmt.Println(client.SMembers("set-key").String())
	fmt.Println(client.SAdd("set-key","a","d","e").String())
	fmt.Println(client.SIsMember("set-key","a"))
	fmt.Println(client.SIsMember("set-key","z"))
	fmt.Println(client.SRem("set-key","d","z").String())
	fmt.Println(client.SMove("set-key","set-key2","d").String())
	fmt.Println(client.SMove("set-key","set-key2","a").String())
	fmt.Println( client.SMembers("set-key2").String())
	client.SAdd("set-key2","b")
	fmt.Println(client.SDiff("set-key","set-key2").String())
	client.SDiffStore("set-diff","set-key","set-key2")
	fmt.Println(client.SMembers("set-diff"))

	fmt.Println("---------test redis hash---------")
	client.HMSet("hash-key",map[string]string{
		"k1":"v1",
		"k2":"v2",
		"k3":"v3",
	})
	fmt.Println(client.HMGet("hash-key","k1","k2").String())
	fmt.Println(client.HMGet("hash-key","k1","k4").String())
	fmt.Println(client.HLen("hash-key").String())
	fmt.Println(client.HDel("hash-key","k1","k4").String())
	fmt.Println(client.HGetAll("hash-key").String())
	fmt.Println(client.HKeys("hash-key").String())
	fmt.Println(client.HVals("hash-key").String())
	fmt.Println(client.HMGet("hash-key","k2","k3").String())
	fmt.Println(client.HExists("hash-key","k4"))

	fmt.Println("---------test redis zset---------")
	client.ZAdd("zadd-key",redis.Z{Score:3,Member:"a"},redis.Z{Score:2,Member:"b"},redis.Z{Score:1,Member:"a"})
	fmt.Println(client.ZCard("zadd-key"))
	fmt.Println(client.ZRangeWithScores("zadd-key",0,-1))
	client.ZAdd("zadd-key",redis.Z{Score:3,Member:"c"})
	client.ZIncrBy("zadd-key",3,"c")
	client.ZIncrBy("zadd-key",1,"a")
	fmt.Println(client.ZRangeWithScores("zadd-key",0,-1))
	fmt.Println(client.ZScore("zadd-key","c"))
	fmt.Println(client.ZRank("zadd-key","a"))
	fmt.Println(client.ZRank("zadd-key","b"))
	fmt.Println(client.ZRank("zadd-key","c"))
	fmt.Println(client.ZRank("zadd-key","d"))
	fmt.Println(client.ZCount("zadd-key","2","6"))
	client.ZRem("zadd-key","a")
	fmt.Println(client.ZRange("zadd-key",0,-1))
	client.ZIncrBy("zadd-key",1,"a")

	fmt.Println("---------test redis pub/sub---------")

}
