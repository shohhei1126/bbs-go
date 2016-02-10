package model

import (
	"database/sql"
	"github.com/Sirupsen/logrus"
	"gopkg.in/gorp.v1"
)

func Init(db *sql.DB, logger *logrus.Logger) *gorp.DbMap {
	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	dbMap.TraceOn("", internalLogger{logurs: logger})
	dbMap.AddTableWithName(User{}, "users").SetKeys(true, "Id")
	dbMap.AddTableWithName(Thread{}, "threads").SetKeys(true, "Id")
	dbMap.AddTableWithName(Comment{}, "comments").SetKeys(true, "Id")
	return dbMap
}

type internalLogger struct {
	logurs *logrus.Logger
}

func (l internalLogger) Printf(format string, v ...interface{}) {
	l.logurs.Debugf(format, v...)
}
