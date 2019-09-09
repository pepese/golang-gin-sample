# golang-gin-sample

[gin](https://github.com/gin-gonic/gin) を利用したサンプルアプリを作成する。

## 初期構築

```bash
$ mkdir -p $GOPATH/src/github.com/pepese
$ cd $GOPATH/src/github.com/pepese
$ cd golang-gin-sample
$ go mod init
$ touch app.go
$ go mod tidy
$ go run app.go
$ curl localhost:8080/ping
```

```bash
$ go build app.go
$ ./app
$ curl localhost:8080/ping
```

```bash
$ GOOS=linux GOARCH=386 go build app.go
$ docker build .
$ docker images
$ docker run -d -p 8080:8080 --name golang-gin xxxx
$ curl localhost:8080/ping
```