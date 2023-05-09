package dao

import (
	"errors"
	"fmt"
	"github.com/go-labs/internal/configs"
	"github.com/go-labs/internal/logging"
	"github.com/go-labs/internal/plugins"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() error {
	dbConf := configs.GetAppConfig().Db

	var err error

	switch dbConf.ServerType {
	case "postgres", "postgresql":
		// create database if not exists
		//driverName := "dm"
		//dataSourceName := "dm://SYSDBA:SYSDBA@localhost:5236"
		dataSourceName := fmt.Sprintf("SYSDBA:SYSDBA@tcp(localhost:5236)/%s", "DMHR")
		preDsn := fmt.Sprintf(dataSourceName)
		logging.Debug().Str("dsn.info", preDsn)
		database, err = gorm.Open(mysql.Open(preDsn), &gorm.Config{})
		if err != nil {
			logging.Error(err).Send()
			return err
		}

		exit := 0
		err = database.Table("pg_database").Select("count(1)").
			Where("datname = ?", dbConf.Database).Scan(&exit).Error
		if err != nil {
			logging.Error(err).Send()
			return err
		}

		if exit == 0 {
			logging.Info().Msg(fmt.Sprintf("Trying to create database: %s", dbConf.Database))
			err = database.Exec(fmt.Sprintf("CREATE DATABASE %s", dbConf.Database)).Error
			if err != nil {
				logging.Error(err).Send()
				return err
			}
		}

		dsn := preDsn
		database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			logging.Error(err).Send()
			return err
		}
	default:
		return errors.New("unsupported database type")
	}

	if configs.GetAppConfig().Db.Debug {
		database = database.Debug()
	}
	err = database.Use(plugins.New())
	if err != nil {
		logging.Error(err).Send()
	}
	//@add: set connection pool configuration
	SetConnPool(dbConf, database)
	return initTables()
}
