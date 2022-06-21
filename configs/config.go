package configs

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type AppConfig struct {
	Database struct {
		User                  string        `yaml:"user" env:"SUMMARISE_DB_USER"`
		Password              string        `yaml:"password" env:"SUMMARISE_DB_PASSWORD"`
		Host                  string        `yaml:"host" env:"SUMMARISE_DB_HOST"`
		Name                  string        `yaml:"name" env:"SUMMARISE_DB_NAME"`
		MaxIdleConnections    int           `yaml:"maxIdleConnections" env:"SUMMARISE_DB_MAX_IDLE_CONNECTIONS"`
		MaxOpenConnections    int           `yaml:"maxOpenConnections" env:"SUMMARISE_DB_MAX_OPEN_CONNECTIONS"`
		MaxConnectionLifeTime time.Duration `yaml:"maxConnectionLifetime" env:"SUMMARISE_DB_MAX_CONNECTION_LIFETIME"`
		MaxConnectionIdleTime time.Duration `yaml:"maxConnectionIdletime" env:"SUMMARISE_DB_MAX_CONNECTION_IDLETIME"`
		DisableTLS            bool          `yaml:"disableTLS" env:"SUMMARISE_DB_DISABLE_TLS"`
		Debug                 bool          `yaml:"debug" env:"SUMMARISE_DB_DEBUG"`
	} `yaml:"database"`
}

func LoadAppConfig() AppConfig {
	var conf AppConfig
	err := cleanenv.ReadConfig("application.yaml", &conf)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return conf
}
