tools:install-go install-glide
		go get github.com/golang/mock/gomock
		go get github.com/golang/mock/mockgen
		go get github.com/clipperhouse/gen

install-go:
		sh sh/install_go.sh 1.5.3

install-glide:
		sh sh/install_glide.sh 0.8.3

deps: install-glide
		glide install

deps-update: install-glide
		glide update

dbs:
		mysql -uroot -h 127.0.0.1 -p -e "CREATE DATABASE IF NOT EXISTS bbs_go"
		mysql -uroot -h 127.0.0.1 -p -e "CREATE DATABASE IF NOT EXISTS bbs_go_test"

dbs-ci:
		mysql -uroot -h 127.0.0.1 -e "CREATE DATABASE IF NOT EXISTS bbs_go"
		mysql -uroot -h 127.0.0.1 -e "CREATE DATABASE IF NOT EXISTS bbs_go_test"

migrate:
		go get bitbucket.org/liamstask/goose/cmd/goose
		goose -env=test up
		goose -env=local up

test:
		go test $(shell go list ./... | grep -v vendor)
		go vet $(shell go list ./... | grep -v vendor)

build:
		go build -o interface-pattern/bbs interface-pattern/main.go