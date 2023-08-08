package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type DataBase struct {
	GormDB *gorm.DB
}

func NewDataBase(dsn string) *DataBase {

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	log.Println("Successfully connected to database")
	gormDB = gormDB.Debug()
	sqlDB, err := gormDB.DB()
	// 设置连接池的最大空闲连接数
	sqlDB.SetMaxIdleConns(10)
	// 设置连接池的最大打开连接数
	sqlDB.SetMaxOpenConns(100)
	if err != nil {
		panic("Failed to close database connection")
	}
	//sqlDB.Close()
	return &DataBase{GormDB: gormDB}
}
