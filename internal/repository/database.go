package repository

import (
	"database/sql"
	oracle "github.com/godoes/gorm-oracle"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"time"
)

type DbOption struct {
	Host               string
	Port               int
	Service            string
	User               string
	Pwd                string
	MaxIdleConnections int
	MaxOpenConnections int
	MaxLifeTime        time.Duration
}

func NewGormDB(cfg *DbOption) *gorm.DB {
	url := oracle.BuildUrl(cfg.Host, cfg.Port, cfg.Service, cfg.User, cfg.Pwd, nil)
	dialector := oracle.New(oracle.Config{
		DSN: url,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		logrus.Fatal("Failed to connect to database: " + err.Error())
	}
	return gormDB
}
func NewSQLDB(gorm *gorm.DB, cfg *DbOption) (*sql.DB, error) {
	sqlDB, err := gorm.DB()
	if err != nil {
		logrus.Error(err)
		return nil, err
	} else {
		err = sqlDB.Ping()
		if err != nil {
			logrus.Fatal("Failed to connect to database: " + err.Error())
		} else {
			logrus.Info("Connected to database")
		}
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConnections)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConnections)

	sqlDB.SetConnMaxLifetime(cfg.MaxLifeTime)

	return sqlDB, nil
}

func ProvideConfig() *DbOption {
	dbOption := &DbOption{
		Host:               viper.GetString("DB_HOST"),
		Port:               viper.GetInt("DB_PORT"),
		Service:            viper.GetString("DB_SERVICE"),
		User:               viper.GetString("DB_USER"),
		Pwd:                viper.GetString("DB_PWD"),
		MaxOpenConnections: viper.GetInt("MAX_OPEN_CONNECTIONS"),
		MaxIdleConnections: viper.GetInt("MAX_IDLE_CONNECTIONS"),
		MaxLifeTime:        time.Duration(viper.GetInt("MAX_LIFE_TIME")),
	}
	return dbOption
}
