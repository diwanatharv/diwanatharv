package db

import (
	"fmt"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var once sync.Once
var dba *gorm.DB

func Makegormserver() *gorm.DB {
	once.Do(func() {

		env := viper.GetString("env") + "."
		user := viper.GetString(env + "db.user")
		password := viper.GetString(env + "db.password")

		host := viper.GetString(env + "db.host")
		port := viper.GetString(env + "db.port")
		dbname := viper.GetString(env + "db.name")
		schema := viper.GetString(env + "db.schema")
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require search_path=%s", host, user, password, dbname, port, schema)
		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		}), &gorm.Config{})
		dba = db
		if err != nil {
			log.Panic().Msgf("Error connecting to the database at %s:%s/%s", host, port, dbname)
		}
		sqlDB, err := dba.DB()
		if err != nil {
			log.Panic().Msgf("Error getting GORM DB definition")
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		log.Info().Msgf("Successfully established connection to %s:%s/%s", host, port, dbname)
	})
	return dba
}
