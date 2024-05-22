/*
 * @Author: i@douxuefeng.cn
 * @Date: 2024-05-21 17:20:34
 * @LastEditTime: 2024-05-21 17:20:41
 * @LastEditors: i@douxuefeng.cn
 * @Description:
 */
package pkg

import (
	"github.com/bwmarrin/snowflake"
)

var SnowFlake *snowflake.Node

func InitSnowflake(epoch int64, i int64) {
	snowflake.Epoch = epoch
	SnowFlake, _ = snowflake.NewNode(i)
}
