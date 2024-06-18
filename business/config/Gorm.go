package config

import "gorm.io/gorm/logger"

type GormConfig struct {
	LogLevel string `yaml:"logLevel" json:"logLevel" default:"info"`
}

func (g *GormConfig) GetLogLevel() (log logger.Interface) {
	switch g.LogLevel {
	case "info":
		log = logger.Default.LogMode(logger.Info)
	case "warn":
		log = logger.Default.LogMode(logger.Warn)
	case "error":
		log = logger.Default.LogMode(logger.Error)
	default:
		log = logger.Default.LogMode(logger.Silent)
	}
	return
}
