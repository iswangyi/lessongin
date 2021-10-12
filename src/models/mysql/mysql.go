package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type (
	// UserInfoMode 原始数据库表字段
	UserInfoMode struct {
		gorm.Model
		Name  string `json:"name"`
		Sex   string `json:"sex"`
		Phone int    `json:"phone"`
		City  string `json:"city"`
	}
	// 处理返回的字段
	transformedUserInfo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

type UserInfo struct {
	Name  string `json:"name"`
	Sex   string `json:"sex"`
	Phone int    `json:"phone"`
	City  string `json:"city"`
}

// ConnectMySQL 连接数据库
func ConnectMySQL() (*gorm.DB, error) {
	alerts := "root:fsdaf123bfdsfABCK@tcp(110.42.129.171:3306)/ops_gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(alerts), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database %v:", err)
		return nil, err
	}
	return db, nil
}

func CreateTable(table interface{}, db *gorm.DB) error {
	err := db.AutoMigrate(table)
	if err != nil {
		return err
	}
	return nil
}
