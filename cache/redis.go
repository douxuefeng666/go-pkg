/*
 * @Author: i@douxuefeng.cn
 * @Date: 2024-05-21 16:43:08
 * @LastEditTime: 2024-05-21 17:22:09
 * @LastEditors: i@douxuefeng.cn
 * @Description:
 */
package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func InitRedis(ctx context.Context, opt *redis.Options) error {
	rdb = redis.NewClient(opt)
	return rdb.Ping(ctx).Err()
}

func GetRdb() *redis.Client {
	return rdb
}
