package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type DataBaseConfig struct {
	Host     string `yaml:"host"`
	DbName   string `yaml:"dbName"`
	User     string `yaml:"username"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
}

type Config struct {
	DataBaseConfig DataBaseConfig `yaml:"database"`
}

func (c *Config) Parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

func NewConfig() (*Config, error) {
	config := &Config{}
	// s := filepath.Dir("config.yaml")
	data, err := ioutil.ReadFile("configs/config.yaml")
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
