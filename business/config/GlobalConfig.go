package config

type GlobalConfig struct {
	Serve Server `json:"server" yaml:"server"  mapstructure:"server"`
	Redis Redis  `json:"redis"  yaml:"redis"   mapstructure:"redis" comment:"redis 全局配置"`
	Mysql Mysql  `json:"mysql"  yaml:"mysql"   mapstructure:"mysql" comment:"mysql多数据配置 默认第一个是 master"`
	Jwt   Jwt    `json:"jwt"    yaml:"jwt"     mapstructure:"jwt"`
}
