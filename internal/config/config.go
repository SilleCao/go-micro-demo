package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

const (
	MySQL    = "mysql"
	MariaDB  = "mariadb"
	Postgres = "postgres"
	SQLite3  = "sqlite3"
)

type DataBaseConfig struct {
	Diver    string `yaml:"diver"`
	Host     string `yaml:"host"`
	DbName   string `yaml:"dbName"`
	User     string `yaml:"username"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
}

type Server struct {
	ContextPath string `yaml:"context-path"`
	Port        int    `yaml:"port"`
}

type Config struct {
	DataBaseConfig DataBaseConfig `yaml:"database"`
	Server         Server         `yaml:"server"`
	// once           sync.Once
	db *gorm.DB
}

func (c *Config) Parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

func NewConfig() (*Config, error) {
	config := &Config{}
	// s := filepath.Dir("config.yaml")
	data, err := os.ReadFile("configs/config.yaml")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if err := config.Parse(data); err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Printf("DBNAME：%s", config.DataBaseConfig.DbName)

	return config, nil
}
