package model

import (
	"fmt"

	"github.com/go-redis/redis"
)

var redisPool *RedisPool

type RedisPool struct {
	client *redis.Client
}

func InitPool(address string, password string) (err error) {
	fmt.Println("开始初始化redis连接池...")
	redisPool = &RedisPool{
		client: redis.NewClient(&redis.Options{
			Addr:     address,
			DB:       0,
			Password: password,
			PoolSize: 10,
		}),
	}
	_, err = redisPool.Get().Ping().Result()
	if err != nil {
		fmt.Println("初始化redis连接池失败...")
		return
	}
	fmt.Println("初始化redis连接池成功...")
	fmt.Println("-------------------------------------")
	return
}

func PrintPoolStats() {
	stats := redisPool.Get().PoolStats()
	fmt.Println("-------------------------------------")
	fmt.Printf("池中的总连接数            %d\n", stats.TotalConns)
	fmt.Printf("池中的空闲连接数          %d\n", stats.IdleConns)
	fmt.Printf("池中删除的过时连接数      %d\n", stats.StaleConns)
	fmt.Printf("池中找到空闲连接的次数    %d\n", stats.Hits)
	fmt.Printf("池中找不到空闲连接的次数  %d\n", stats.Misses)
	fmt.Printf("等待超时发生的次数        %d\n", stats.Timeouts)
	fmt.Println("-------------------------------------")
}

func (this *RedisPool) Get() (client *redis.Client) {
	return this.client
}
