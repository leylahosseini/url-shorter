package db

import (
	"fmt"
	"leylahosseini/my-url-shorter/db/models"

	//	"github.com/leylahosseini/my-url-shorter/db/models"

	//"github.com/labstack/gommon/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//var dsn string

func ConnectToMysql() *gorm.DB {

	// user := os.Getenv("DB_USER")
	// pass := os.Getenv("DB_PASS")
	// url := os.Getenv("DB_HOST")
	// dbname := os.Getenv("DB_NAME")
	//port := os.Getenv("DB_PORT")

	//dsn := "user:pass@tcp(url:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:123456@tcp(mysql:3306)/shorter?charset=utf8mb4&parseTime=True&loc=Local"

	//if datasource == "mysql" {
	//	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, url, dbname)
	//dsn := "user:pass@tcp(url:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error Connection Database")
		panic(err)
	}
	fmt.Println("Connected to Database ...")

	db.AutoMigrate(&models.Url{})

	return db
}
