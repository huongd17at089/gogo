package db

import (
	"first_demo/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//var DB *gorm.DB

func InitDB(config config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Pass, config.Host, config.Port, config.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
		//logs.Error(err)
		fmt.Println("CONNECT MYSQL FAILED!")
		fmt.Println(err.Error())
	}

	fmt.Println("CONNECT MYSQL SUCCESS!")
	return db, err
}

func GetCon() *gorm.DB {
	config, err := config.ReadConfig("config/config.yml")
	var con *gorm.DB
	if err != nil {
		panic("invalid ConfigPath")
	} else {
		con, err = InitDB(config.DBconf)
		if err != nil {
			panic("can't connect db")
		}
	}
	return con
}
