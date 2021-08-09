include .env
export

APP_NAME=boxius

GOPATH=$(shell go env GOPATH)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

CURRENT_PATH=$(shell pwd)
GO111MODULE=on

deploy-domainesia:
	npx pm2 stop --silent ${APP_NAME} || :
	npx pm2 delete --silent ${APP_NAME} || :
	chmod +x ${APP_NAME}.app
	npx pm2 start ${APP_NAME}.json
