package main

import (
	"fmt"
	"strings"

	"github.com/SilleCao/golang/go-micro-demo/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var db *gorm.DB

func connectDB(dsn string) (db *gorm.DB) {
	var err error
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	return db
}

func init() {
	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	db = connectDB(conf.DatabaseDsn())
}

func main() {

	g := gen.NewGenerator(gen.Config{
		OutPath:      "./internal/pkg/dao",
		ModelPkgPath: "./internal/modules/sys/model",
	})

	g.UseDB(db)
	g.WithJSONTagNameStrategy(nameStrategy)
	// g.ApplyBasic(g.GenerateAllTable()...)

	// g.GenerateModel("sys_dept")
	// g.GenerateModel("sys_dict_data")
	// g.GenerateModel("sys_dict_type")
	// g.GenerateModel("sys_log_error")
	// g.GenerateModel("sys_log_login")
	// g.GenerateModel("sys_log_operation")
	// g.GenerateModel("sys_menu")
	// g.GenerateModel("sys_oss")
	// g.GenerateModel("sys_params")
	g.GenerateModel("sys_role")
	// g.GenerateModel("sys_role_data_scope")
	// g.GenerateModel("sys_role_menu")
	// g.GenerateModel("sys_role_user")
	// g.ApplyBasic(g.GenerateModel("sys_user"))
	// g.GenerateModel("sys_user_token")
	g.Execute()

}

func nameStrategy(c string) string {
	if c == "<empty>" {
		return c
	}
	subNames := strings.Split(c, "_")
	var newName string
	for index, subName := range subNames {
		if index == 0 {
			newName += subName
			continue
		}
		subName = strings.ToUpper(subName[:1]) + subName[1:]
		newName += subName

	}
	return newName
}
