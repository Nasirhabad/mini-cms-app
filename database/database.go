package database

import (
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

VAR(
	DB_USERNAME string = "root"
	DB_PASSWORD string = "" 
	DB_NAME     string = "minicmsdb"
	DB_HOST     string = "localhost"
	DB_PORT     string = "3306"
)

func InitDB() {
	var err error
	var dsn string = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)%s?charset=utf8mb48parseTime=True&loc=Local",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)
	
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.config{})
	if err != nill {
		log.fatalf("eeror when connecting to the database: %s\n",err)
	}

	log.Println("connected to the database")

}