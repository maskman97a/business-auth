package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	go_ora "github.com/sijms/go-ora/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "gorm.io/driver/postgres"
	"time"
)

type DbOption struct {
	Host    string
	Port    int
	Service string
	User    string
	Pwd     string
}

func NewSQLDB(cfg *DbOption) (*sql.DB, error) {
	connStr := go_ora.BuildUrl(cfg.Host, cfg.Port, cfg.Service, cfg.User, cfg.Pwd, nil)
	sqlDB, err := sql.Open("oracle", connStr)
	if err != nil {
		logrus.Error(err)
		return nil, err
	} else {
		logrus.Info("Connected to database")
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
		Host:    viper.GetString("DB_HOST"),
		Port:    viper.GetInt("DB_PORT"),
		Service: viper.GetString("DB_SERVICE"),
		User:    viper.GetString("DB_USERNAME"),
		Pwd:     viper.GetString("DB_PASSWORD"),
	}
	return dbOption
}
