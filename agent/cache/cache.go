package cache

import (
	"github.com/go-redis/redis"
	"time"
)


//var redisservers = gin.H{
//	"Addr" : "127.0.0.1:6379",
//	"Password" : "",
//	"DB": 0,
//}



type Cache struct {
	cacheObj *redis.Client
}



func New() (*Cache, error){
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "", // redis密码，没有则留空
		DB:       0,  // 默认数据库，默认是0
	})
	_,err:=client.Ping(client.Context()).Result()

	if err != nil {
		return nil, err
	}
	cache := &Cache{
		cacheObj: client,
	}
	return cache, nil
}


func (c *Cache) Get(key string) (interface{}, error){
	val, err := c.cacheObj.Get(c.cacheObj.Context(), key).Result()
	return val, err
}


func (c *Cache) Set(key string, value interface{}, timeout int32) error{
	err := c.cacheObj.Set(c.cacheObj.Context(), key, value, time.Duration(timeout)).Err()

	return err
}


var CommonCache, _ = New()