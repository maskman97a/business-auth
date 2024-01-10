package repository_impl

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type DbOption struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func NewGormDB(cfg *DbOption) *gorm.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Database, cfg.Password)
	logrus.Info("connection string: ", connStr)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func NewSQLDB(gormDB *gorm.DB) (*sql.DB, error) {
	sqlDB, err := gormDB.DB()
	if err != nil {
		logrus.Error(err)
		return nil, err
	} else {
		logrus.Info("Connection to database successful")
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(viper.GetInt("MAX_IDLE_CONNS"))

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(viper.GetInt("MAX_OPEN_CONNS"))

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	var maxLifeTime, err2 = time.ParseDuration(viper.GetString("MAX_LIFETIME"))
	if err2 != nil {
		logrus.Error(err2)
		return nil, err2
	}
	sqlDB.SetConnMaxLifetime(time.Minute * maxLifeTime)

	return sqlDB, nil
}

func ProvideConfig() *DbOption {
	dbOption := &DbOption{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		Database: viper.GetString("DB_DATABASE"),
		Username: viper.GetString("DB_USERNAME"),
		Password: viper.GetString("DB_PASSWORD"),
	}
	return dbOption
}
