package shrd_service

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	shrd_utils "github.com/StevanoZ/dv-shared/utils"
	"github.com/go-redis/redis/v8"
)

type RedisClient interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd
}

type CacheSvc interface {
	Set(ctx context.Context, key string, data any) error
	Get(ctx context.Context, key string, output any) error
	Del(ctx context.Context, key string) error
	DelByPrefix(ctx context.Context, prefixName string)
	GetOrSet(ctx context.Context, key string, function func() any) (any, error)
}

type CacheSvcImpl struct {
	config  *shrd_utils.BaseConfig
	cacheDb RedisClient
}

func NewCacheSvc(config *shrd_utils.BaseConfig, cacheDb RedisClient) CacheSvc {
	return &CacheSvcImpl{
		config:  config,
		cacheDb: cacheDb,
	}
}

func NewRedisClient(config *shrd_utils.BaseConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_HOST,
		Password: config.REDIS_PASSWORD, // no password set
		DB:       0,                     // use default DB
		Username: config.REDIS_USERNAME,
	})
	return rdb
}
func (s *CacheSvcImpl) Set(ctx context.Context, key string, data any) error {

	if data != nil {
		if reflect.TypeOf(data).Kind() == reflect.Slice {
			if reflect.ValueOf(data).Len() == 0 {
				fmt.Println("no data to save, array is empty")
				return nil
			}
		}

		cacheData, err := json.Marshal(data)

		if err != nil {
			return err
		}

		return s.cacheDb.Set(ctx, key, cacheData, s.config.CACHE_DURATION).Err()
	}
	return nil
}

func (s *CacheSvcImpl) Get(ctx context.Context, key string, output any) error {

	val, err := s.cacheDb.Get(ctx, key).Result()

	if err != nil {
		fmt.Println(fmt.Sprintf("failed when getting key -> %s ", key), err)
		return err
	}

	err = json.Unmarshal([]byte(val), &output)

	if err != nil {
		fmt.Println("failed when unmarshal data")
		return err
	}

	fmt.Println("get data from cache with key -->", key)

	return nil
}

func (s *CacheSvcImpl) DelByPrefix(ctx context.Context, prefixName string) {
	var foundedRecordCount int = 0
	iter := s.cacheDb.Scan(ctx, 0, fmt.Sprintf("%s*", prefixName), 0).Iterator()
	fmt.Printf("your search pattern: %s\n", prefixName)

	for iter.Next(ctx) {
		fmt.Printf("deleted= %s\n", iter.Val())
		s.cacheDb.Del(ctx, iter.Val())
		foundedRecordCount++
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	fmt.Printf("deleted Count %d\n", foundedRecordCount)
}

func (s *CacheSvcImpl) GetOrSet(ctx context.Context, key string, function func() any) (any, error) {
	var data any
	err := s.Get(ctx, key, &data)

	if err != nil && err == redis.Nil {
		data = function()
		err := s.Set(ctx, key, data)
		fmt.Println("set data to cache with key -->", key)

		return data, err
	}

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *CacheSvcImpl) Del(ctx context.Context, key string) error {
	err := s.cacheDb.Del(ctx, key).Err()

	if err != nil {
		return err
	}

	return nil
}
