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
```