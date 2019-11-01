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