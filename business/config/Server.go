package config

type Server struct {
	Name string `mapstructure:"name" json:"name" yaml:"name" defalut:"blog"`
	Port int    `mapstructure:"port" json:"port" yaml:"port" defalut:"9999"`
}
