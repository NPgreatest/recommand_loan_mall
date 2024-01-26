package initialize

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/global"
	"main.go/initialize/internal"
)

func GormPostgres() *gorm.DB {
	p := global.GVA_CONFIG.Postgres // 假设您有一个类似的配置结构体
	if p.Dbname == "" {
		return nil
	}
	postgresConfig := postgres.Config{
		DSN:                  p.Dsn(), // DSN data source name
		PreferSimpleProtocol: true,    // 禁用批量插入
	}
	if db, err := gorm.Open(postgres.New(postgresConfig), internal.Gorm.Config()); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		return db
	}
}
