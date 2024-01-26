package config

import "fmt"

type Postgres struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`                             // 服务器地址
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                             // 端口
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`                     // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                 // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                 // 数据库密码
	SSLMode      string `mapstructure:"sslmode" json:"sslmode" yaml:"sslmode"`                    // SSL模式
	TimeZone     string `mapstructure:"time-zone" json:"timeZone" yaml:"time-zone"`               // 时区
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
}

func (p *Postgres) Dsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		p.Host, p.Username, p.Password, p.Dbname, p.Port, p.SSLMode, p.TimeZone)
}

//host=localhost user=postgres password=root dbname=vector_db port=5432 sslmode=disable TimeZone=Asia/Shanghai
