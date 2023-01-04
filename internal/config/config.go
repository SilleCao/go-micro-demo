package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/SilleCao/golang/go-micro-demo/internal/mutex"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/dao"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
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

type Config struct {
	DataBaseConfig DataBaseConfig `yaml:"database"`
	// once           sync.Once
	db *gorm.DB
}

func (c *Config) Init() error {
	start := time.Now()
	if err := c.connectDb(); err != nil {
		return err
	}

	log.Fatalf("config: successfully initialized [%s]", time.Since(start))
	return nil
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

func (c *Config) DatabaseDsn() string {
	switch c.DataBaseConfig.Diver {
	case MySQL, MariaDB:
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4,utf8&collation=utf8mb4_unicode_ci&parseTime=true",
			c.DataBaseConfig.User,
			c.DataBaseConfig.Password,
			c.DataBaseConfig.Host,
			c.DataBaseConfig.Port,
			c.DataBaseConfig.DbName,
		)
	default:
		log.Fatal("config: empty database dsn")
		return ""
	}
}

func (c *Config) connectDb() error {
	// Make sure this is not running twice.
	mutex.Db.Lock()
	defer mutex.Db.Unlock()

	// Get database driver and data source name.
	dbDriver := c.DataBaseConfig.Diver
	dbDsn := c.DatabaseDsn()

	if dbDriver == "" {
		return errors.New("config: database driver not specified")
	}

	if dbDsn == "" {
		return errors.New("config: database DSN not specified")
	}

	// Open database connection.
	db, err := gorm.Open(mysql.Open(c.DatabaseDsn()), &gorm.Config{})
	if err != nil || db == nil {
		for i := 1; i <= 12; i++ {
			db, err = gorm.Open(mysql.Open(c.DatabaseDsn()), &gorm.Config{})

			if db != nil && err == nil {
				break
			}

			time.Sleep(5 * time.Second)
		}

		if err != nil || db == nil {
			return err
		}
	}

	// Configure database logging.
	// db.LogMode(false)
	// db.SetLogger(log)

	// // Set database connection parameters.
	// db.DB().SetMaxOpenConns(c.DatabaseConns())
	// db.DB().SetMaxIdleConns(c.DatabaseConnsIdle())
	// db.DB().SetConnMaxLifetime(time.Hour)

	// // Check database server version.
	// if err = c.checkDb(db); err != nil {
	// 	if c.Unsafe() {
	// 		log.Error(err)
	// 	} else {
	// 		return err
	// 	}
	// }

	// Ok.
	c.db = db

	return nil
}

func (c *Config) InitDB() {
	c.connectDb()
	c.RegisterDb()
}

// Db returns the db connection.
func (c *Config) Db() *gorm.DB {
	if c.db == nil {
		log.Fatal("config: database not connected")
	}
	return c.db
}

func (c *Config) RegisterDb() {
	dao.SetDbProvider(c)
}
