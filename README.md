# golang-gin-sample

[gin](https://github.com/gin-gonic/gin) を利用したサンプルアプリを作成する。

## 実行

```zsh
% go mod tidy
% go run main.go
% curl localhost:8080/ping
```

## Docker

```zsh
% docker build -t app .
% docker run -d -p 8080:8080 --name app app
% curl localhost:8080/ping
% docker container stop app
```

## Docker Compose

アプリをローカル実行する際の

## 構成

- Web Framework : [gin](https://github.com/gin-gonic/gin)
- ORM : [gorm](https://github.com/jinzhu/gorm)
- Logger : [zap](https://github.com/uber-go/zap)
  - [gin と連携](https://github.com/gin-contrib/zap)
- Env : [envconfig]](https://github.com/kelseyhightower/envconfig) 、 [dotenv](https://github.com/joho/godotenv)
- Security : [secure](https://github.com/unrolled/secure)
- Validator : [validator](https://github.com/go-playground/validator)
- etc : [ここみる？](https://github.com/go-playground)