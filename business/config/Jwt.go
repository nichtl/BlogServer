package config

type Jwt struct {
	SigningKey  string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`
	ExpiresTime string `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"`
	BufferTime  string `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
}
