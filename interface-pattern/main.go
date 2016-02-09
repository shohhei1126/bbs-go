package main

import (
	"github.com/shohhei1126/bbs-go/common/db"
	"github.com/shohhei1126/bbs-go/interface-pattern/model"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shohhei1126/bbs-go/common/conf"
	"github.com/shohhei1126/bbs-go/common/http/response"
	"github.com/shohhei1126/bbs-go/common/log"
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
	parseLogger(conf.LogFile, conf.LogLevel)
	dbm := parseDb(conf.DbMaster)
	dbs := parseDb(conf.DbSlave)
	dbMMap := model.Init(dbm, log.Logger)
	dbSMap := model.Init(dbs, log.Logger)

	mux := goji.NewMux()
	userDao := dao.NewUser(dbMMap, dbSMap)
	userService := service.NewUser(userDao)
	userHandler := handler.NewUser(userService)

	mux.HandleFuncC(pat.Get("/v1/users/:id"), wrap(userHandler.Show))
	mux.Handle(pat.Get("/*"), http.FileServer(http.Dir(conf.Assets)))

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
	db, err := db.NewMySqlDb(dbString)
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
