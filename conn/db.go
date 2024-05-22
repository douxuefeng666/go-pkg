/*
 * @Author: i@douxuefeng.cn
 * @Date: 2024-05-21 16:17:00
 * @LastEditTime: 2024-05-21 16:23:57
 * @LastEditors: i@douxuefeng.cn
 * @Description:
 */
package conn

import (
	"context"
	"database/sql"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	db    *gorm.DB
	sqlDb *sql.DB
)

type DBOpt struct {
	Dsn                                      string //master
	MaxIdle                                  int
	MaxOpen                                  int
	LogMode                                  logger.LogLevel
	Models                                   []any
	EnableMigrate                            bool
	DisableForeignKeyConstraintWhenMigrating bool
}

func InitDb(ctx context.Context, opt *DBOpt) error {
	var err error
	db, err = gorm.Open(mysql.Open(opt.Dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: opt.DisableForeignKeyConstraintWhenMigrating,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(opt.LogMode),
		NowFunc: func() time.Time {
			return time.Now().Truncate(time.Millisecond)
		},
	})
	if err != nil {
		return err
	}
	if opt.EnableMigrate {
		if err := db.AutoMigrate(opt.Models...); err != nil {
			return err
		}
	}
	//连接池设置
	if sqlDb, err = db.DB(); err != nil {
		return err
	} else {
		sqlDb.SetConnMaxIdleTime(time.Hour)
		sqlDb.SetMaxIdleConns(opt.MaxIdle)
		sqlDb.SetMaxOpenConns(opt.MaxOpen)
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	if sqlDb.Ping() != nil {
		return
	}
	_ = sqlDb.Close()
}
