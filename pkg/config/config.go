package config

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config captures the application configuration parameters
type Config struct {
	Port int
}

// FromFile returns a Config struct from the input configFile
func FromFile(configFile string) (Config,error) {
	if configFile == "" {
		return Config{}, fmt.Errorf("missing configFile")
	}
	cfgDdir, cfgFile := filepath.Split(configFile)
	v := viper.New()
	v.SetConfigName(cfgFile)
	v.SetConfigType(filepath.Ext(cfgFile)[1:])
	v.AddConfigPath(cfgDdir)
	if err := v.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("couldn't load config: %w", err)
	}
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		return Config{}, err
	}
	if cfg.Port <= 1024 || cfg.Port > 65535 {
		return Config{}, fmt.Errorf("port %d is not valid", cfg.Port)
	}
	return cfg, nil
}
