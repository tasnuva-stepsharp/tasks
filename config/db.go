

package config

import (
 "gorm.io/driver/postgres"
 "gorm.io/gorm"
 "log"
)

func InitDB() (*gorm.DB, error) {
 dsn := "host=localhost user=postgres password=postgres dbname=tasks port=5432 sslmode=disable"
 db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
 if err != nil {
 log.Fatal(err)
 }
 return db, nil
}

