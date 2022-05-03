package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"xiangerfer.com/db_flagship/model"
	_ "github.com/go-sql-driver/mysql"

)


var DB *gorm.DB;

func InitDB() *gorm.DB{


	driverName := viper.GetString("datasource.dirverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	user := viper.GetString("datasource.username")
	passwored := viper.GetString("datasource.password")
	database := viper.GetString("datasource.database")
	charset :=  viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		user, passwored, host, port, database, charset)
	fmt.Printf(args)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect mysql, err: " + err.Error())
	}

	db.AutoMigrate(&model.User{}, &model.Host{})
	DB = db
	return db
}


func GetDB() *gorm.DB{
	return DB
}