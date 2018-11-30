- https://github.com/gin-gonic/gin

```bash
$ mkdir -p $GOPATH/src/github.com/pepese
$ cd $GOPATH/src/github.com/pepese
$ git clone https://github.com/pepese/golang-gin-sample.git
$ cd golang-gin-sample
$ echo "vendor/" > .gitignore
$ dep init
$ touch app.go
$ dep ensure -v
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