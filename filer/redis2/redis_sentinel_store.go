package redis2

import (
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/bookstairs/bookworm/filer"
	"github.com/bookstairs/bookworm/util"
)

func init() {
	filer.Stores = append(filer.Stores, &Redis2SentinelStore{})
}

type Redis2SentinelStore struct {
	UniversalRedis2Store
}

func (store *Redis2SentinelStore) GetName() string {
	return "redis2_sentinel"
}

func (store *Redis2SentinelStore) Initialize(configuration util.Configuration, prefix string) (err error) {
	return store.initialize(
		configuration.GetStringSlice(prefix+"addresses"),
		configuration.GetString(prefix+"masterName"),
		configuration.GetString(prefix+"username"),
		configuration.GetString(prefix+"password"),
		configuration.GetInt(prefix+"database"),
	)
}

func (store *Redis2SentinelStore) initialize(addresses []string, masterName string, username string, password string, database int) (err error) {
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
