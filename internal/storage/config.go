package storage

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

const (
	envPrefix          string = "DB"
	dataSourceTemplate string = "host=%s port=%d user=%s password=%s dbname=%s"
)

type Config struct {
	Host     string `envconfig:"DB_HOST"`
	Port     int    `envconfig:"DB_PORT"`
	User     string `envconfig:"DB_USER"`
	Password string `envconfig:"DB_PASS"`
	DbName   string `envconfig:"DB_NAME"`
	DbDriver string `default:"pgx" envconfig:"DB_DRIVER"`
}

func LoadConfig() *Config {
	var config Config
	if err := envconfig.Process(envPrefix, &config); err != nil {
		log.Fatal(err)
	}
	return &config
}

func (config Config) GetDataSourceName() string {
	return fmt.Sprintf(dataSourceTemplate, config.Host, config.Port, config.User,
		config.Password, config.DbName)
}