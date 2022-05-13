package shrd_service

import (
	"context"
	"testing"
	"time"

	shrd_utils "github.com/StevanoZ/dv-shared/utils"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	STRUCT_KEY = "cache-struct-key"
	SLICE_KEY  = "cache-slice-key"
	CACHE      = "cache"
)

type testData struct {
	ID        uuid.UUID
	CreatedAt time.Time
}

func initCacheSvc() (CacheSvc, *redis.Client) {
	config := shrd_utils.LoadBaseConfig("../app", "test")
	redisClient := NewRedisClient(config)

	cacheSvc := NewCacheSvc(config, redisClient)

	return cacheSvc, redisClient
}

func TestCacheSvc(t *testing.T) {
	ctx := context.Background()
	cacheSvc, redisClient := initCacheSvc()
	assert.NotNil(t, redisClient)
	defer redisClient.Close()

	assert.NotNil(t, cacheSvc)

	dataStruct := testData{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
	}

	dataSlice := []testData{
		{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
		},
	}

	t.Run("Set, get and delete data (struct)", func(t *testing.T) {
		output := testData{}

		err := cacheSvc.Get(ctx, STRUCT_KEY, &output)
		assert.Equal(t, redis.Nil, err)

		err = cacheSvc.Set(ctx, STRUCT_KEY, dataStruct)
		assert.NoError(t, err)

		cacheSvc.Get(ctx, STRUCT_KEY, &output)

		assert.Equal(t, dataStruct.ID, output.ID)
		assert.Equal(t, dataStruct.CreatedAt.Year(), output.CreatedAt.Year())
		assert.Equal(t, dataStruct.CreatedAt.Month(), output.CreatedAt.Month())
		assert.Equal(t, dataStruct.CreatedAt.Day(), output.CreatedAt.Day())
		assert.Equal(t, dataStruct.CreatedAt.Hour(), output.CreatedAt.Hour())
		assert.Equal(t, dataStruct.CreatedAt.Minute(), output.CreatedAt.Minute())
		assert.Equal(t, dataStruct.CreatedAt.Second(), output.CreatedAt.Second())

		cacheSvc.DelByPrefix(ctx, CACHE)
	})

	t.Run("Set, get and delete data (slice)", func(t *testing.T) {
		err := cacheSvc.Set(ctx, SLICE_KEY, dataSlice)
		assert.NoError(t, err)

		output := []testData{}
		cacheSvc.Get(ctx, SLICE_KEY, &output)

		assert.NotEmpty(t, output)
		assert.Equal(t, 1, len(output))
		cacheSvc.DelByPrefix(ctx, CACHE)
	})

	t.Run("Not save key to cache when data is empty", func(t *testing.T) {
		err := cacheSvc.Set(ctx, SLICE_KEY, []testData{})
		assert.NoError(t, err)

		output := []testData{}
		cacheSvc.Get(ctx, SLICE_KEY, &output)

		assert.Empty(t, output)
		assert.Equal(t, 0, len(output))
	})

	t.Run("Not save key to cache when data is nil", func(t *testing.T) {
		err := cacheSvc.Set(ctx, SLICE_KEY, nil)
		assert.NoError(t, err)

		output := []testData{}
		cacheSvc.Get(ctx, SLICE_KEY, &output)

		assert.Empty(t, output)
		assert.Equal(t, 0, len(output))
	})

	t.Run("Failed get cache when given invalid payload type", func(t *testing.T) {
		cacheSvc.Set(ctx, STRUCT_KEY, dataStruct)
		err := cacheSvc.Get(ctx, STRUCT_KEY, &[]testData{})
		assert.Error(t, err)

		cacheSvc.DelByPrefix(ctx, CACHE)
	})

	t.Run("Failed set cache when given invalid data type", func(t *testing.T) {
		err := cacheSvc.Set(ctx, STRUCT_KEY, func() {})
		assert.Error(t, err)

	})

	t.Run("Set data to cache, when data is not available in cache", func(t *testing.T) {
		output := []testData{}

		err := cacheSvc.Get(ctx, SLICE_KEY, &output)
		assert.Error(t, err)
		assert.Equal(t, 0, len(output))

		data, err := cacheSvc.GetOrSet(ctx, SLICE_KEY, func() any {
			return dataSlice
		})

		assert.NoError(t, err)
		assert.Equal(t, 1, len(data.([]testData)))

		err = cacheSvc.Get(ctx, SLICE_KEY, &output)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(output))

		cacheSvc.DelByPrefix(ctx, CACHE)
	})

	t.Run("Get data from cache when key exist", func(t *testing.T) {
		cacheSvc.Set(ctx, SLICE_KEY, dataSlice)

		data, err := cacheSvc.GetOrSet(ctx, SLICE_KEY, func() any {
			return nil
		})

		assert.NotNil(t, data)
		assert.NoError(t, err)

		output := []testData{}

		err = shrd_utils.ConvertInterface(data, &output)

		assert.NoError(t, err)
		assert.Equal(t, 1, len(output))
		assert.IsType(t, []testData{}, output)

		cacheSvc.DelByPrefix(ctx, CACHE)
	})

	t.Run("Delete key from cache", func(t *testing.T) {
		cacheSvc.Set(ctx, SLICE_KEY, dataSlice)

		output := []testData{}

		err := cacheSvc.Get(ctx, SLICE_KEY, &output)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(output))

		err = cacheSvc.Del(ctx, SLICE_KEY)
		assert.NoError(t, err)

		err = cacheSvc.Del(ctx, SLICE_KEY)
		assert.Error(t, redis.Nil, err)
	})

	t.Run("Close client", func(t *testing.T) {
		err := cacheSvc.CloseClient()
		assert.NoError(t, err)
	})
}
