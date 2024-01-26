package config

type Server struct {
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// oss
	Postgres Postgres `mapstructure:"postgres" json:"postgres" yaml:"postgres"`
	Local    Local    `mapstructure:"local" json:"local" yaml:"local"`
	Keys     Keys     `mapstructure:"keys" json:"keys" yaml:"keys"`
}
