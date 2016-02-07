tools:install-go install-glide
		go get github.com/golang/mock/gomock
		go get github.com/golang/mock/mockgen
		go get github.com/clipperhouse/gen
		go get bitbucket.org/liamstask/goose/cmd/goose

install-go:
		sh sh/install_go.sh 1.5.3

install-glide:
		sh sh/install_glide.sh 0.8.3

deps: install-glide
		glide install

test:
		go test $(shell go list github.com/shohhei1126/bbs-go/... | grep -v vendor)