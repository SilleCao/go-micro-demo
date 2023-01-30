package dao

import (
	"fmt"
	"log"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Supported test databases.
// const (
// 	MySQL = "mysql"
// SQLite3         = "sqlite3"
// SQLiteTestDB    = ".test.db"
// SQLiteMemoryDSN = ":memory:"
// )

// dbConn is the global gorm.DB connection provider.
var dbConn Gorm

// Gorm is a gorm.DB connection provider interface.
type Gorm interface {
	Db() *gorm.DB
}

// DbConn is a gorm.DB connection provider.
type DbConn struct {
	Driver string
	Dsn    string

	once sync.Once
	db   *gorm.DB
}

// Db returns the gorm db connection.
func (g *DbConn) Db() *gorm.DB {
	g.once.Do(g.Open)

	if g.db == nil {
		log.Fatal("migrate: database not connected")
	}

	return g.db
}

// Open creates a new gorm db connection.
func (g *DbConn) Open() {
	// if HasDbProvider() {
	// 	g.db = dbConn.Db()
	// 	return
	// }
	db, err := gorm.Open(mysql.Open(g.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil || db == nil {
		for i := 1; i <= 12; i++ {
			fmt.Printf("gorm.Open(%s, %s) %d\n", g.Driver, g.Dsn, i)
			gorm.Open(mysql.Open(g.Dsn), &gorm.Config{})
			if db != nil && err == nil {
				break
			} else {
				time.Sleep(5 * time.Second)
			}
		}

		if err != nil || db == nil {
			fmt.Println(err)
			log.Fatal(err)
		}
	}
	// db.SetLogger(log)
	// db.DB().SetMaxIdleConns(4)
	// db.DB().SetMaxOpenConns(256)

	g.db = db
}

// Close closes the gorm db connection.
// func (g *DbConn) Close() {
// 	if g.db != nil {
// 		if err := g.db.Close(); err != nil {
// 			log.Fatal(err)
// 		}

// 		g.db = nil
// 	}
// }

// // IsDialect returns true if the given sql dialect is used.
// func IsDialect(name string) bool {
// 	return name == Db().Dialect().GetName()
// }

// // DbDialect returns the sql dialect name.
// func DbDialect() string {
// 	return Db().Dialect().GetName()
// }

// SetDbProvider sets the Gorm database connection provider.
func SetDbProvider(conn Gorm) {
	dbConn = conn
}

// HasDbProvider returns true if a db provider exists.
func HasDbProvider() bool {
	return dbConn != nil
}
