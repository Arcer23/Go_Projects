package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error
 
func DatabaseInit(){
	host := "localhost"
	user := "arcer"
	password := "pass123"
	dbName := "new-db"
	port:=3306

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&tls=false", user, password, host, port, dbName)
	database,e = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e);
	}
}
func DB() *gorm.DB {
	return database
}