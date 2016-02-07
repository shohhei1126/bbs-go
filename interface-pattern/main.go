package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/shohhei1126/bbs-go/common/db"
	"github.com/shohhei1126/bbs-go/interface-pattern/model"

	"database/sql"
	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shohhei1126/bbs-go/common/logger"
)

type appConf struct {
	DbMaster string `envconfig:"db_master"`
	DbSlave  string `envconfig:"db_slave"`
	LogLevel string `envconfig:"log_level"`
}

func main() {
	conf := parseAppConf()
	dbm := parseDb(conf.DbMaster)
	dbs := parseDb(conf.DbSlave)
	logger := parseLogger(conf.LogLevel)
	dbMMap := model.Init(dbm, logger)
	dbSMap := model.Init(dbs, logger)
	fmt.Printf("%+v", dbMMap)
	fmt.Printf("%+v", dbSMap)
}

func parseAppConf() appConf {
	appConf := appConf{}
	err := envconfig.Process("bbsgo", &appConf)
	if err != nil {
		panic(err)
	}
	return appConf
}

func parseDb(dbString string) *sql.DB {
	db, err := db.NewMySqlDb(dbString)
	if err != nil {
		panic(err)
	}
	return db
}

func parseLogger(logLevel string) *logrus.Logger {
	l, err := logger.NewLogger(logger.Conf{LogLevel: logLevel})
	if err != nil {
		panic(err)
	}
	return l
}
