package database

import (
	"fmt"

	"order-service/public/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConn() *gorm.DB {
	host := config.Host
	dbUsername := config.User
	dbPassword := config.Password
	database := config.DBName
	port := config.PORTDB

	// DSN (Data Source Name) format untuk PostgreSQL berbeda dengan pgsql
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Shanghai", host, dbUsername, dbPassword, database, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("can't connect to database")
	}
	return db
}
