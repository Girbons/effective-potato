package config

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var conf *PotatoConfig

type PotatoConfig struct {
	JWTSigningKey string        `mapstructure:"jwt_signing_key"`
	JWTExpiration time.Duration `mapstructure:"jwt_expiration"`
}

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	viper.SetDefault("jwt_expiration", 24)
}

func (p *PotatoConfig) loadConfig() (*PotatoConfig, error) {
	if err := viper.ReadInConfig(); err != nil {
		log.Error(err)
		return p, err
	}

	return p, viper.Unmarshal(&p)
}

// GetConf returns the service Configuration
func GetConf() (*PotatoConfig, error) {
	return conf.loadConfig()
}
