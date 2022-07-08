package main

import (
	"fmt"
	"rpr/app/proxy"
	"rpr/entities"
	"rpr/pkg/logger"
	"rpr/pkg/redis"
)

func main() {
	logger.InitLogger(
		"storage/logs/logs.log",
		64,
		5,
		30,
		false,
		"single",
		"debug",
	)

	// 建立 Redis 连接
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", "175.178.212.108", 16379),
		"",
		"",
		0,
	)

	for name := range redis.Redis.HGetAll("heartbeats") {
		proxy.AddProxy(proxy.Proxy{ID: name})
	}

	proxies := proxy.GetProxies()
	for _, proxy := range proxies {
		uuids := redis.Redis.SMembers(fmt.Sprintf("proxy:%s:usersOnline", proxy.ID))
		if len(uuids) <= 0 {
			continue
		}

		players := redis.Redis.HMGets("uuid-cache", uuids)
		for _, player := range players {
			player, err := entities.NewCachedPlayerFromJSON(player)
			logger.LogIf(err)
			logger.InfoString("Proxy", proxy.ID, player.Name)
		}
	}
}
