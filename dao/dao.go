package dao

import (
	"collector/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dao struct {
	db *gorm.DB
}

func New() *Dao {
	// 配置数据库连接字符串
	dsn := "root:password@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"

	// 打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 自动迁移：自动创建/更新数据库表结构
	err = db.AutoMigrate(&model.TourismDB{})
	if err != nil {
		log.Fatal("Error migrating database:", err)
	}

	return &Dao{
		db: db,
	}

}
