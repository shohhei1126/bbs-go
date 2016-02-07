package model

import (
	"database/sql"
	"github.com/Sirupsen/logrus"
	"github.com/shohhei1126/bbs-go/interface-pattern/model/user"
	"gopkg.in/gorp.v1"
)

func Init(db *sql.DB, logger *logrus.Logger) *gorp.DbMap {
	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	dbMap.TraceOn("", internalLogger{logurs: logger})
	dbMap.AddTableWithName(user.User{}, "users").SetKeys(true, "Id")
	return dbMap
}

type internalLogger struct {
	logurs *logrus.Logger
}

func (l internalLogger) Printf(format string, v ...interface{}) {
	l.logurs.Debugf(format, v...)
}
