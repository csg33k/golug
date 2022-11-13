package config

import (
	"log"
	"path/filepath"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

const (
	ENV_PREFIX = "GOLUG"
)

var (
	instance *Config
	doOnce   sync.Once
)

type Config struct {
	viper     *viper.Viper
	AppConfig *ServerConfig
}
type ServerConfig struct {
	Server ServerSettings `mapstructure:"server"`
}

type ServerSettings struct {
	Address     string `mapstructure:"address"`
	DatabaseUri string `mapstructure:"dburi"`
}

func loadViperConfig(configPath string) *ServerConfig {
	searchDir := filepath.Dir(configPath)
	appName := filepath.Base(configPath)
	extention := filepath.Ext(appName)
	appName = appName[0 : len(appName)-len(extention)]

	v := viper.New()
	v.SetEnvPrefix(ENV_PREFIX)
	v.SetConfigName(appName)
	//Adds a few search paths for various use cases
	v.AddConfigPath("..")
	v.AddConfigPath("conf")
	v.AddConfigPath("../conf")
	v.AddConfigPath(searchDir)

	v.SetDefault("server.address", "localhost:3000")
	replacer := strings.NewReplacer(".", "__")
	v.SetEnvKeyReplacer(replacer)

	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	var cfg ServerConfig
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("Error decoding config file, %s", err)
	}

	if NewConfig().viper == nil {
		NewConfig().viper = v
		NewConfig().AppConfig = &cfg
	}

	return &cfg

}

func LoadConfig(configPath string) *ServerConfig {
	return loadViperConfig(configPath)
}

func NewConfig() *Config {
	doOnce.Do(func() {
		instance = &Config{}
	})

	return instance
}
