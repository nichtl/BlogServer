package config

import (
	"fmt"
	"strconv"
	"strings"
)

type Mysql struct {
	Name                  string  `json:"name" mapstructure:"name" yaml:"name" comment:"数据源配置名称"`
	Host                  string  `json:"host" mapstructure:"host" yaml:"host"`
	Port                  int     `json:"port" mapstructure:"port" yaml:"port"`
	Enable                bool    `json:"enable"  mapstructure:"enable" yaml:"enable"`
	Database              string  `json:"database" mapstructure:"database" yaml:"database" comment:"数据库名称"`
	Username              string  `json:"username" mapstructure:"username" yaml:"username"`
	Password              string  `json:"password" mapstructure:"password" yaml:"password"`
	MaxOpenConn           int     `json:"maxOpenConn" mapstructure:"maxOpenConn" yaml:"maxOpenConn" comment:"数据库连接最大连接数"`
	MinIdelConn           int     `json:"minIdelConn" mapstructure:"maxOpenConn" yaml:"minIdelConn" comment:"空闲连接池最大连接数"`
	ConnMaxLifeTimeMillis int     `json:"connMaxLifeTimeMillis" mapstructure:"connMaxLifeTimeMillis" yaml:"connMaxLifeTimeMillis" comment:"连接池连接可复用的最大时间" `
	Config                string  `json:"config" mapstructure:"config" yaml:"config"`
	Slave                 []Mysql `yaml:"slave" json:"slave" `
}

func (m *Mysql) GetDSN() (dsn string, err error) {
	if m.Host == "" || m.Database == "" || m.Port == 0 || m.Username == "" || m.Password == "" {
		err = fmt.Errorf("mysql initialize error  invalid config")
	}
	port := strconv.Itoa(m.Port)

	if index := strings.Index(m.Config, "?"); index != -1 {
		m.Config = m.Config[index+1:]
	}

	dsn = m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + port + ")/" + m.Database + "?" + m.Config
	return dsn, nil
}
