[![TravisBuildStatus](https://api.travis-ci.org/bridgewwater/golang-project-temple-docker-db.svg?branch=master)](https://travis-ci.org/bridgewwater/golang-project-temple-docker-db)
[![GoDoc](https://godoc.org/github.com/bridgewwater/golang-project-temple-docker-db?status.png)](https://godoc.org/github.com/bridgewwater/golang-project-temple-docker-db/)
[![GoReportCard](https://goreportcard.com/badge/github.com/bridgewwater/golang-project-temple-docker-db)](https://goreportcard.com/report/github.com/bridgewwater/golang-project-temple-docker-db)
[![codecov](https://codecov.io/gh/bridgewwater/golang-project-temple-docker-db/branch/master/graph/badge.svg)](https://codecov.io/gh/bridgewwater/golang-project-temple-docker-db)

## for what

- this project used to github golang

replace

- `bridgewwater/golang-project-temple-docker-db` to you code pkg
- `containt out port` replace `35001`
- `mysql port` replace `35011`
- `redist port` replace `35021`

- and run

## depends

in go mod project

```bash
# warning use privte git host must set
# global set for once
# add private git host like github.com to evn GOPRIVATE
$ go env -w GOPRIVATE='github.com'
# use ssh proxy
# set ssh-key to use ssh as http
$ git config --global url."git@github.com:".insteadOf "http://github.com/"
# or use PRIVATE-TOKEN
# set PRIVATE-TOKEN as gitlab or gitea
$ git config --global http.extraheader "PRIVATE-TOKEN: {PRIVATE-TOKEN}"
# set this rep to download ssh as https use PRIVATE-TOKEN
$ git config --global url."ssh://github.com/".insteadOf "https://github.com/"

# before above global settings
# test version info
$ git ls-remote -q http://github.com/bridgewwater/golang-project-temple-docker-db.git

# test depends see full version
$ go list -v -m -versions github.com/bridgewwater/golang-project-temple-docker-db
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -m -versions github.com/bridgewwater/golang-project-temple-docker-db | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## use

```bash
make init
```

- test code

```bash
make test
```

add main.go file and run

```bash
make run
```

# dev

## evn

- golang sdk 1.13+

## docker

base docker file can replace

# cli tools to init project fast

```
$ curl -L --fail https://raw.githubusercontent.com/bridgewwater/golang-project-temple-docker-db/master/temp-golang-base
# let temp-golang-base file folder under $PATH
$ chmod +x temp-golang-base
# see how to use
$ temp-golang-base -h
```
