package main

import (
	"github.com/shohhei1126/bbs-go/model"
	bbsmysql "github.com/shohhei1126/bbs-go/mysql"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shohhei1126/bbs-go/conf"
	"github.com/shohhei1126/bbs-go/dao"
	"github.com/shohhei1126/bbs-go/handler"
	"github.com/shohhei1126/bbs-go/http/response"
	"github.com/shohhei1126/bbs-go/log"
	"github.com/shohhei1126/bbs-go/service"
	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
	"net/http"
)

func main() {
	conf := parseConf()
	parseLogger(conf.LogFile, conf.LogLevel)
	dbm := parseDb(conf.DbMaster)
	dbs := parseDb(conf.DbSlave)
	dbMMap := model.Init(dbm, log.Logger)
	dbSMap := model.Init(dbs, log.Logger)

	mux := goji.NewMux()
	userDao := dao.NewUser(dbMMap, dbSMap)
	threadDao := dao.NewThread(dbMMap, dbSMap)
	threadService := service.NewThread(userDao, threadDao)
	threadHandler := handler.NewThread(threadService)
	mux.HandleFuncC(pat.Get("/v1/threads"), wrap(threadHandler.List))
	log.Logger.Info("starting server...")
	http.ListenAndServe("localhost:8080", mux)
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
	db, err := bbsmysql.NewMySqlDb(dbString)
	if err != nil {
		panic(err)
	}
	return db
}

func parseLogger(logfile, logLevel string) {
	err := log.Init(log.Conf{LogFile: logfile, LogLevel: logLevel})
	if err != nil {
		panic(err)
	}
}
