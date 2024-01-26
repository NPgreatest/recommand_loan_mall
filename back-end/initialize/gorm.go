package initialize

import (
	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm(target string) *gorm.DB {
	switch target {
	case "mysql":
		return GormMysql()
	case "postgres":
		return GormPostgres()
	default:
		return nil
	}
}
