/*******************************************************************************
* 2019 - present Contributed by Apulis Technology (Shenzhen) Co. LTD
*
* This program and the accompanying materials are made available under the
* terms of the MIT License, which is available at
* https://www.opensource.org/licenses/MIT
*
* See the NOTICE file distributed with this work for additional
* information regarding copyright ownership.
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
* WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
* License for the specific language governing permissions and limitations
* under the License.
*
* SPDX-License-Identifier: MIT
******************************************************************************/

package dao

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-labs/internal/configs"
	"github.com/go-labs/internal/logging"
	"github.com/go-labs/internal/models"
	"reflect"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ctxTransactionKey struct{}

var database *gorm.DB
var statsDB *gorm.DB //统计推理db

func InitDb() error {
	dbConf := configs.GetAppConfig().Db

	var err error

	switch dbConf.ServerType {
	case "postgres", "postgresql":
		// create database if not exists
		preDsn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=%s",
			dbConf.Host, dbConf.Port, dbConf.Username, dbConf.Password, dbConf.Sslmode)
		logging.Debug().Str("dsn.info", preDsn)
		database, err = gorm.Open(postgres.Open(preDsn), &gorm.Config{})
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

		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", dbConf.Host,
			dbConf.Port, dbConf.Username, dbConf.Password, dbConf.Database, dbConf.Sslmode)
		database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
	//@add: set connection pool configuration
	SetConnPool(dbConf, database)
	return initTables()
}

func GetDB(ctx context.Context) *gorm.DB {
	if ctx == nil {
		return database
	}

	tx, ok := ctx.Value(ctxTransactionKey{}).(*gorm.DB)
	if ok {
		return tx
	}

	return database
}

func GetDBBase() *gorm.DB {
	return database
}

func initStatsDB(dbConf configs.DbConfig) {
	var err error

	if dbConf.StatsDataBase == "" {
		dbConf.StatsDataBase = "pg_data_source_connector_db"
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", dbConf.Host,
		dbConf.Port, dbConf.Username, dbConf.Password, dbConf.StatsDataBase, dbConf.Sslmode)
	statsDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logging.Error(err).Msgf("create db connection error")
		panic(err)
	}
	statsDB = statsDB.Debug()
	statsConn, err := statsDB.DB()
	if err != nil {
		logging.Error(err).Msgf("get db instance error")
		panic(err)
	}
	if configs.GetAppConfig().Db.MaxIdleConns > 0 {
		statsConn.SetMaxIdleConns(configs.GetAppConfig().Db.MaxIdleConns)
	}
	if configs.GetAppConfig().Db.MaxOpenConns > 0 {
		statsConn.SetMaxOpenConns(configs.GetAppConfig().Db.MaxOpenConns)
	}
	if configs.GetAppConfig().Db.MaxConnLifeTime > 0 {
		statsConn.SetConnMaxLifetime(time.Duration(configs.GetAppConfig().Db.MaxConnLifeTime) * time.Second)
	} else {
		statsConn.SetConnMaxLifetime(time.Duration(1800) * time.Second)
	}

	go pingDB(statsDB)
}

func pingDB(db *gorm.DB) {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	//nolint gosimple
	for {
		select {
		case <-ticker.C:
			sqlDb, _ := db.DB()
			if err := sqlDb.Ping(); err != nil {
				timeStr := time.Now().Format("2006-01-02 15:04:05")
				logging.Error(err).Msgf("ping db error: %s", timeStr)
			}
		}
	}
}

func initTables() error {

	modelTypes := []interface{}{&models.Product{}}

	for _, modelType := range modelTypes {
		err := autoMigrateTable(modelType)
		if err != nil {
			logging.Error(err).Send()
			return err
		}
	}

	return nil
}

func autoMigrateTable(modelType interface{}) error {
	val := reflect.Indirect(reflect.ValueOf(modelType))
	modelName := val.Type().Name()

	logging.Info().Str("modelName", modelName).Msg("Migrating Table")

	err := database.AutoMigrate(modelType)
	if err != nil {
		logging.Error(err).Send()
		return err
	}
	return nil
}

// ExecDBTx 执行DB 事务操作，在fn中执行多个DB操作，需要带 ctx 给dao层
func ExecDBTx(fn func(ctx context.Context) error) error {
	mCtx := context.Background()
	return GetDB(mCtx).Transaction(func(tx *gorm.DB) error {
		newCtx := context.WithValue(mCtx, ctxTransactionKey{}, tx)
		return fn(newCtx)
	})
}
func SetConnPool(dbConf configs.DbConfig, db *gorm.DB) {
	//dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", dbConf.Host,
	//	dbConf.Port, dbConf.Username, dbConf.Password, dbConf.Database, dbConf.Sslmode)
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	if dbConf.MaxIdleConns > 0 {
		sqlDb.SetMaxIdleConns(dbConf.MaxIdleConns)
	}
	if dbConf.MaxOpenConns > 0 {
		sqlDb.SetMaxOpenConns(dbConf.MaxOpenConns)
	}
	if dbConf.MaxConnLifeTime > 0 {
		sqlDb.SetConnMaxLifetime(time.Duration(dbConf.MaxConnLifeTime) * time.Second)
	} else {
		sqlDb.SetConnMaxLifetime(time.Duration(3600) * time.Second)
	}

	data, _ := json.Marshal(sqlDb.Stats())
	logging.Info().Str("db.stats", string(data)).Send()

	go pingDB(db)

}
