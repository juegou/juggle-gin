// +----------------------------------------------------------------------
// | Juggle [ 让我们能更好的杂耍 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2020 http://www.XXXXXX.com All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: dingo <djhui1987@gmail.com>
// +----------------------------------------------------------------------

// Xorm适配器
package adapter

import (
	_ "github.com/go-sql-driver/mysql"
	"juggle-gin/src/juggle"
	"time"
	"xorm.io/xorm"
)

type XormAdapter struct {
	*xorm.Engine
}

// 初始化Gorm适配器
func NewXormAdapter() *XormAdapter {
	// 连接数据库
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(192.168.137.128:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		juggle.Error(err)
	}


	// 配置数据库连接池
	engine.DB().SetMaxIdleConns(5) //最大空闲数
	engine.DB().SetMaxOpenConns(10)//最大打开连接数
	engine.DB().SetConnMaxLifetime(time.Second*30)  //空闲连接生命周期

	return &XormAdapter{Engine: engine}
}