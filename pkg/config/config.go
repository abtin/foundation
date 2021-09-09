package config

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

// WebServerSection is the webserver part of the config file
type WebServerSection struct {
	Address string
	Port    int
	Name    string
}

type AuthSection struct {
	TokenDuration  time.Duration
	CookieDuration time.Duration
}

// Config captures the application configuration parameters
type Config struct {
	WebServer          WebServerSection `mapstructure:"webserver"`
	Auth               AuthSection      `mapstructure:"auth"`
	JwtSecretKey       string           `mapstructure:"JWT_SECRET_KEY"`
	GithubClientID     string           `mapstructure:"GITHUB_CLIENT_ID"`
	GithubClientSecret string           `mapstructure:"GITHUB_CLIENT_SECRET"`
}

// FromFile returns a Config struct from the input configFile
func FromFile(configFile string) (Config, error) {
	if configFile == "" {
		return Config{}, fmt.Errorf("missing configFile")
	}

	cfgDdir, cfgFile := filepath.Split(configFile)
	if ext := filepath.Ext(cfgFile); ext != ".yaml" {
		return Config{}, fmt.Errorf("%s config files are not supported. Please use a .yaml file", ext)
	}
	v := viper.New()
	v.SetConfigName(cfgFile)
	v.SetConfigType(filepath.Ext(cfgFile)[1:])
	v.AddConfigPath(cfgDdir)
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("couldn't load config: %w", err)
	}

	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func (c Config) WebServerURL() string {
	return fmt.Sprintf("%s:%d", c.WebServer.Address, c.WebServer.Port)
}
