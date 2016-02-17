# bbs-go

This is a Golang web application using an interface.

[![Circle CI](https://circleci.com/gh/shohhei1126/bbs-go.svg?style=svg)](https://circleci.com/gh/shohhei1126/bbs-go)

## preparing

```
$ GO15VENDOREXPERIMENT=1
$ go get github.com/shohhei1126/bbs-go
$ cd $GOPATH/src/github.com/shohhei1126/bbs-go
$ make install-glide
$ make deps
$ make dbs
$ make migrate // if you need a password of root, modify db/dbconf.yml
```

## testing

```
$ export "BBSGO_DB_TEST=root:{YOUR_PASSWORD}@tcp(localhost:3306)/bbs_go_test?parseTime=true&loc=Local" // if you need
$ make test
go test github.com/shohhei1126/bbs-go github.com/shohhei1126/bbs-go/bbstime github.com/shohhei1126/bbs-go/conf github.com/shohhei1126/bbs-go/dao github.com/shohhei1126/bbs-go/handler github.com/shohhei1126/bbs-go/http/response github.com/shohhei1126/bbs-go/log github.com/shohhei1126/bbs-go/model github.com/shohhei1126/bbs-go/mysql github.com/shohhei1126/bbs-go/service
?   	github.com/shohhei1126/bbs-go	[no test files]
?   	github.com/shohhei1126/bbs-go/bbstime	[no test files]
?   	github.com/shohhei1126/bbs-go/conf	[no test files]
ok  	github.com/shohhei1126/bbs-go/dao	0.248s
ok  	github.com/shohhei1126/bbs-go/handler	0.008s
?   	github.com/shohhei1126/bbs-go/http/response	[no test files]
?   	github.com/shohhei1126/bbs-go/log	[no test files]
ok  	github.com/shohhei1126/bbs-go/model	0.005s
?   	github.com/shohhei1126/bbs-go/mysql	[no test files]
ok  	github.com/shohhei1126/bbs-go/service	0.006s
go vet github.com/shohhei1126/bbs-go github.com/shohhei1126/bbs-go/bbstime github.com/shohhei1126/bbs-go/conf github.com/shohhei1126/bbs-go/dao github.com/shohhei1126/bbs-go/handler github.com/shohhei1126/bbs-go/http/response github.com/shohhei1126/bbs-go/log github.com/shohhei1126/bbs-go/model github.com/shohhei1126/bbs-go/mysql github.com/shohhei1126/bbs-go/service
```

## running

```
$ export "BBSGO_DB_MASTER=root:{YOUR_PASSWORD}@tcp(localhost:3306)/bbs_go?parseTime=true&loc=Local"
$ export "BBSGO_DB_SLAVE=root:{YOUR_PASSWORD}@tcp(localhost:3306)/bbs_go?parseTime=true&loc=Local"
$ go run main.go
INFO[0000] starting server...
```

## calling api

```
$ curl -XGET "http://localhost:8080/v1/threads?limit=5&offset=0"
[{"id":9,"title":"i","body":"i","createdAt":"2016-02-17T10:36:53+09:00","updatedAt":"2016-04-17T10:36:55+09:00","user":{"id":1,"username":"hoge","password":"hoge","displayName":"hoge","status":"member","createAt":"2016-02-16T21:13:57+09:00","updatedAt":"2016-02-16T21:13:58+09:00"}},{"id":7,"title":"g","body":"g","createdAt":"2016-02-17T10:36:30+09:00","updatedAt":"2016-04-17T10:36:32+09:00","user":{"id":3,"username":"bar","password":"bar","displayName":"bar","status":"member","createAt":"2016-02-16T21:14:50+09:00","updatedAt":"2016-02-16T21:14:52+09:00"}},{"id":12,"title":"l","body":"l","createdAt":"2016-02-17T10:37:44+09:00","updatedAt":"2016-02-21T10:37:45+09:00","user":{"id":4,"username":"buzz","password":"buzz","displayName":"buzz","status":"member","createAt":"2016-02-16T21:15:22+09:00","updatedAt":"2016-02-16T21:15:24+09:00"}}]
```
