package redis_lua

import (
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/bookstairs/bookworm/filer"
	"github.com/bookstairs/bookworm/util"
)

func init() {
	filer.Stores = append(filer.Stores, &RedisLuaSentinelStore{})
}

type RedisLuaSentinelStore struct {
	UniversalRedisLuaStore
}

func (store *RedisLuaSentinelStore) GetName() string {
	return "redis_lua_sentinel"
}

func (store *RedisLuaSentinelStore) Initialize(configuration util.Configuration, prefix string) (err error) {
	return store.initialize(
		configuration.GetStringSlice(prefix+"addresses"),
		configuration.GetString(prefix+"masterName"),
		configuration.GetString(prefix+"username"),
		configuration.GetString(prefix+"password"),
		configuration.GetInt(prefix+"database"),
	)
}

func (store *RedisLuaSentinelStore) initialize(addresses []string, masterName string, username string, password string, database int) (err error) {
	store.Client = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:      masterName,
		SentinelAddrs:   addresses,
		Username:        username,
		Password:        password,
		DB:              database,
		MinRetryBackoff: time.Millisecond * 100,
		MaxRetryBackoff: time.Minute * 1,
		ReadTimeout:     time.Second * 30,
		WriteTimeout:    time.Second * 5,
	})
	return
}
