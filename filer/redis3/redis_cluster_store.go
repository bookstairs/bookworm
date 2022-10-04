package redis3

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"

	"github.com/bookstairs/bookworm/filer"
	"github.com/bookstairs/bookworm/util"
)

func init() {
	filer.Stores = append(filer.Stores, &RedisCluster3Store{})
}

type RedisCluster3Store struct {
	UniversalRedis3Store
}

func (store *RedisCluster3Store) GetName() string {
	return "redis_cluster3"
}

func (store *RedisCluster3Store) Initialize(configuration util.Configuration, prefix string) (err error) {

	configuration.SetDefault(prefix+"useReadOnly", false)
	configuration.SetDefault(prefix+"routeByLatency", false)

	return store.initialize(
		configuration.GetStringSlice(prefix+"addresses"),
		configuration.GetString(prefix+"password"),
		configuration.GetBool(prefix+"useReadOnly"),
		configuration.GetBool(prefix+"routeByLatency"),
	)
}

func (store *RedisCluster3Store) initialize(addresses []string, password string, readOnly, routeByLatency bool) (err error) {
	store.Client = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:          addresses,
		Password:       password,
		ReadOnly:       readOnly,
		RouteByLatency: routeByLatency,
	})
	store.redsync = redsync.New(goredis.NewPool(store.Client))
	return
}