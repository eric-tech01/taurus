MainVersion = v0.0.1
SUB_VERSION = $(shell git describe --tags --always)
VERSION = ${MainVersion}_${SUB_VERSION}
GO_VERSION=$(shell go version)
# BUILD_TIME=$(shell date +%F-%Z/%T)
BUILD_TIME = $(shell powershell -Command "Get-Date -Format 'yyyy-MM-dd HH:mm:ss'")

CommitId=$(shell git rev-parse HEAD) # 项目是 Git 版本控制，则获取当前 commit id

LDFLAGS = " \
-X 'github.com/eric-tech01/simple-version.Version=${VERSION}' \
-X 'github.com/eric-tech01/simple-version.GoVersion=${GO_VERSION}' \
-X 'github.com/eric-tech01/simple-version.BuildTime=${BUILD_TIME}' \
-X 'github.com/eric-tech01/simple-version.CommitId=${CommitId}'"

all:
	go build -ldflags ${LDFLAGS} -o main_${VERSION}.exe .