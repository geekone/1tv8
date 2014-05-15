package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
	"github.com/robfig/config"
)

var (
	Engine *xorm.Engine
)

func init(){
	revel.OnAppStart(Init)
}

func Init(){
	c, err := config.ReadDefault(revel.BasePath + "/conf/my.conf")
	driver, _ := c.String("database", "db.driver")
	dbname, _ := c.String("database", "db.dbname")
	user, _ := c.String("database", "db.user")
	password, _ := c.String("database", "db.password")
	host, _ := c.String("database", "db.host")

	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", user, password, host, dbname)
	Engine, err = xorm.NewEngine(driver, params)
	if err !=nil{
		panic(err)
	}
	Engine.ShowSQL = revel.DevMode
	err = Engine.Sync(
		new(User),
		new(Category),
		new(Movie),
		new(Article),
		// new(Reply),
		// new(Permissions),
	)

	if err != nil {
		panic(err)
	}
}