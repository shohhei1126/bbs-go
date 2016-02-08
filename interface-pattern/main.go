package main

import (
	"github.com/shohhei1126/bbs-go/common/db"
	"github.com/shohhei1126/bbs-go/interface-pattern/model"

	"database/sql"
	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shohhei1126/bbs-go/common/conf"
	"github.com/shohhei1126/bbs-go/common/http/response"
	"github.com/shohhei1126/bbs-go/common/logger"
	"github.com/shohhei1126/bbs-go/interface-pattern/dao"
	"github.com/shohhei1126/bbs-go/interface-pattern/handler"
	"github.com/shohhei1126/bbs-go/interface-pattern/service"
	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
	"net/http"
)

func main() {
	conf := parseConf()
	dbm := parseDb(conf.DbMaster)
	dbs := parseDb(conf.DbSlave)
	logger := parseLogger(conf.LogLevel)
	dbMMap := model.Init(dbm, logger)
	dbSMap := model.Init(dbs, logger)

	mux := goji.NewMux()

	userDao := dao.NewUser(dbMMap, dbSMap)
	userService := service.NewUser(userDao)
	userHandler := handler.NewUser(userService)

	mux.HandleFuncC(pat.Get("/user/:id"), wrap(userHandler.Show))

	http.ListenAndServe("localhost:8000", mux)
}

func wrap(action func(ctx context.Context, req *http.Request) response.Response) func(context.Context, http.ResponseWriter, *http.Request) {
	return func(ctx context.Context, out http.ResponseWriter, req *http.Request) {
		res := action(ctx, req)
		if res == nil {
			res = response.ServerError
		}
		res.WriteTo(out)
	}
}

func parseConf() *conf.Conf {
	conf, err := conf.Parse()
	if err != nil {
		panic(err)
	}
	return conf
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
