package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	goredis "github.com/go-redis/redis/v8"
)

type redisDriver struct {
	client *goredis.Client
}

type RedisOptions struct {
	Endpoint string
	Password string
}

type dataTyp uint8

const (
	redisString dataTyp = iota
)

func NewRedisDriver(options *RedisOptions) (*redisDriver, error) {
	client := goredis.NewClient(&goredis.Options{
		Addr:     options.Endpoint,
		Password: options.Password,
		DB:       0, // default DB,
		// TLSConfig: &tls.Config{},
	})
	return &redisDriver{
		client: client,
	}, nil
}

func (d *redisDriver) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pong, err := d.client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("redis连接失败: ", pong, err)
		return err
	} else {
		fmt.Println("redis连接成功: ", pong)
		return nil
	}
}

func (d *redisDriver) Close() error {
	return d.client.Close()
}

func (d *redisDriver) Get(key string, redisTyp dataTyp) (interface{}, error) {
	switch redisTyp {
	case redisString:
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		val, err := d.client.Get(ctx, key).Result()
		switch {
		case err == goredis.Nil:
			return nil, errors.New("key dose not exist")
		case err != nil:
			return nil, fmt.Errorf("Get failed %w", err)
		case val == "":
			return nil, errors.New("value is empty")
		}
		return val, nil
	}
	return nil, errors.New("未实现指定类型的数据查询")
}
