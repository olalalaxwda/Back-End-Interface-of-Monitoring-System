package orm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"monitoring/orm/tables"
)

var (
	Db *gorm.DB
)

func Init() {
	dsn := "monitoring:monitoring@tcp(127.0.0.1:3306)/monitoring?charset=utf8mb4&parseTime=True&loc=Local"
	Db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// 自动迁移
	Db.AutoMigrate(&tables.ApiFailed{})
	Db.AutoMigrate(&tables.ApiSucceed{})
	Db.AutoMigrate(&tables.Cpstate{})
	Db.AutoMigrate(&tables.Jspe{})
	Db.AutoMigrate(&tables.Reject{})
	Db.AutoMigrate(&tables.Rse{})
	Db.AutoMigrate(&tables.Ws{})
	Db.AutoMigrate(&tables.KeyPerformanceData{})
	Db.AutoMigrate(&tables.Ps{})
	Db.AutoMigrate(&tables.Pv{})
	Db.AutoMigrate(&tables.Uv{})
}
