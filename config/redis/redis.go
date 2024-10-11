package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go-gin-template/config"
	"go-gin-template/config/log"
)

var ctx = context.Background()

// ######################### 单机版本 #################################################

// Rdb 单机版本
var Rdb *redis.Client

func InitRedis() {

	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.ServiceConfig.Redis.Ip + ":" + config.ServiceConfig.Redis.Port,
		Password: config.ServiceConfig.Redis.Password,
		DB:       config.ServiceConfig.Redis.Db,
	})

	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		log.Errorf("Redis缓存数据库连接异常: %v", err)
	}
	log.Infof("===========Redis缓存数据库连接成功==============")
}

// ######################### 读写分离版本 #################################################

// RdbReadWrite 读写分离
var RdbReadWrite *RedisReadWrite

// RedisReadWrite 读写分离
type RedisReadWrite struct {
	Master *redis.Client
	Slaves []*redis.Client
}

// InitRedisReadWrite 创建 Redis 读写分离客户端
func InitRedisReadWrite() {
	// 主节点地址
	masterAddr := "192.168.2.22:6379"
	// 从节点地址列表
	slaveAddrs := []string{"192.168.2.22:6380"}

	masterOpt := &redis.Options{
		Addr:     masterAddr,
		Password: "", // 多节点不推荐使用密码
		DB:       0,  // use default DB
	}

	masterClient := redis.NewClient(masterOpt)

	var slaves []*redis.Client
	if len(slaveAddrs) > 0 {
		for _, addr := range slaveAddrs {
			slaveOpt := &redis.Options{
				Addr:     addr,
				Password: "", // 多节点不推荐使用密码
				DB:       0,  // use default DB
			}
			slaves = append(slaves, redis.NewClient(slaveOpt))
		}
	} else {
		slaves = append(slaves, masterClient)
	}

	RdbReadWrite = &RedisReadWrite{
		Master: masterClient,
		Slaves: slaves,
	}
}

// ######################### 切片集群版本 #################################################

// RdbCluster 切片集群
var RdbCluster *redis.ClusterClient

// InitRedisCluster 创建 Redis 切片集群客户端
func InitRedisCluster() {
	// Redis Cluster 的节点列表
	nodes := []string{
		"192.168.2.22:6379",
		"192.168.2.22:6380",
		"192.168.2.22:6381",
		"192.168.2.22:6382",
		"192.168.2.22:6383",
		"192.168.2.22:6384",
	}

	// 创建集群客户端
	RdbCluster = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    nodes,
		ReadOnly: true, // 读写分离
	})
}
