package config

import "strconv"

type Redis struct {
	Db       int    `json:"db" mapstructure:"db" yaml:"db"`
	Port     int16  `json:"port" mapstructure:"port" yaml:"port"`
	Host     string `json:"host" mapstructure:"host" yaml:"host"`
	Password string `json:"password" mapstructure:"password" yaml:"password"`
}

func (r *Redis) GetAddr() string {
	return r.Host + ":" + strconv.FormatInt(int64(r.Port), 10)
}
