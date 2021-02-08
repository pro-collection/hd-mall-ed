package cache

import (
	"github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"
	"github.com/go-redis/redis/v7"
	"hd-mall-ed/packages/client/config"
)

var Manager *cache.Cache

func SetUp() {
	redisStore := store.NewRedis(redis.NewClient(&redis.Options{
		Addr: config.DatabaseConfig.RedisHost,
	}), &store.Options{
		Expiration: config.DatabaseConfig.RedisExpire,
	})

	Manager = cache.New(redisStore)

	// 使用示例
	/*
	err := Manager.Set("my-key", "my-value", &store.Options{Expiration: 15 * time.Second})
	if err != nil {
		panic(err)
	}

	value, _ := Manager.Get("my-key")
	fmt.Println(value)
	*/
}
