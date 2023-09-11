package repository

//var DBInstance *gorm.DB
//
//func InitDB() {
//	dsn := config.Global.System.Db
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic("Failed to connect to database")
//	}
//	log.Println("Successfully connected to database")
//	db = db.Debug()
//	sqlDB, err := db.DB()
//	// 设置连接池的最大空闲连接数
//	sqlDB.SetMaxIdleConns(10)
//	// 设置连接池的最大打开连接数
//	sqlDB.SetMaxOpenConns(100)
//	DBInstance = db
//	if err != nil {
//		panic("Failed to close database connection")
//	}
//	//sqlDB.Close()
//}
