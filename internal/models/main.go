package models

import (
	"fmt"
	"github.com/betterde/ects/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

type Model struct {
	Page int64 `xorm:"-"`
	PageSize int64 `xorm:"-"`
}

var Engine *xorm.Engine

const DefaultTimeFormat = "2006-01-02 15:04:05"

func Connection() (*xorm.Engine, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		config.Conf.Database.User,
		config.Conf.Database.Pass,
		config.Conf.Database.Host,
		config.Conf.Database.Port,
		config.Conf.Database.Name,
		config.Conf.Database.Char,
	)
	engine, err := xorm.NewEngine("mysql", dsn)
	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(30)

	go keepAlived()

	return engine, err
}

// 定时Ping，保证连接不被服务器断开
func keepAlived() {
	t := time.Tick(180 * time.Second)
	for {
		<-t
		if err := Engine.Ping(); err != nil {
			// TODO
		}
	}
}

// 迁移数据库
func Migrate() error {
	tables := []interface{}{
		&User{},
	}

	for _, table := range tables {
		exist, err := Engine.IsTableExist(table)
		if exist {
			continue
		}

		if err != nil {
			return err
		}

		err = Engine.Sync2(table)
		if err != nil {
			return err
		}
	}

	return nil
}