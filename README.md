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

アプリをローカル実行する際の MySQL を起動する際に使用する。

```zsh
% docker-compose up -d
% go run main.go
```

## 構成

- Web Framework : [gin](https://github.com/gin-gonic/gin)
- ORM : [gorm](https://github.com/jinzhu/gorm)
    - mysql : `go get -u github.com/go-sql-driver/mysql`
- Logger : [zap](https://github.com/uber-go/zap)
  - [gin と連携](https://github.com/gin-contrib/zap)
- Env : [envconfig](https://github.com/kelseyhightower/envconfig) 、 [dotenv](https://github.com/joho/godotenv)
- Security : [secure](https://github.com/unrolled/secure)
  - [cors](https://github.com/gin-contrib/cors) はどうしよう、、、
- Validator : [validator](https://github.com/go-playground/validator)
- etc : [ここみる？](https://github.com/go-playground)

## テーブルの確認

```zsh
% docker exec -it <container_id> /bin/bash
/# mysql
mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
| testdb             |
+--------------------+
5 rows in set (0.00 sec)

mysql> use testdb;
mysql> show tables;
+------------------+
| Tables_in_testdb |
+------------------+
| users            |
+------------------+
1 row in set (0.00 sec)

mysql> explain users;
+------------+------------------+------+-----+---------+----------------+
| Field      | Type             | Null | Key | Default | Extra          |
+------------+------------------+------+-----+---------+----------------+
| id         | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| created_at | timestamp        | YES  |     | NULL    |                |
| updated_at | timestamp        | YES  |     | NULL    |                |
| deleted_at | timestamp        | YES  | MUL | NULL    |                |
| first_name | varchar(255)     | YES  |     | NULL    |                |
| last_name  | varchar(255)     | YES  |     | NULL    |                |
+------------+------------------+------+-----+---------+----------------+
6 rows in set (0.01 sec)

mysql> quit;
/# exit
```

## 単体テスト

```zsh
% go test ./...

% go test -v -cover ./...
% go test -coverprofile=./coverage/cover.out ./...
% go tool cover -html=./coverage/cover.out -o ./coverage/cover.html
% open ./coverage/cover.html
```

## 単結合テスト

```zsh
% curl -X POST -H 'Content-Type:application/json' -d '{"first_name":"first","last_name":"last"}' localhost:8080/api/v1/users | jq .
{
  "id": 1,
  "first_name": "first",
  "last_name": "last",
  "CreatedAt": "2019-11-11T18:55:04.015057+09:00",
  "UpdatedAt": "2019-11-11T18:55:04.015057+09:00",
  "DeletedAt": null
}

% curl localhost:8080/api/v1/users | jq .
[
  {
    "id": 1,
    "first_name": "first",
    "last_name": "last",
    "CreatedAt": "2019-11-11T18:55:04+09:00",
    "UpdatedAt": "2019-11-11T18:55:04+09:00",
    "DeletedAt": null
  }
]

% curl -X PUT -H 'Content-Type:application/json' -d '{"first_name":"First","last_name":"Last"}' localhost:8080/api/v1/users/1 | jq .
{
  "id": 1,
  "first_name": "First",
  "last_name": "Last",
  "CreatedAt": "0001-01-01T00:00:00Z",
  "UpdatedAt": "2019-11-11T18:55:53.262974+09:00",
  "DeletedAt": null
}

% curl localhost:8080/api/v1/users/1 | jq .
{
  "id": 1,
  "first_name": "First",
  "last_name": "Last",
  "CreatedAt": "2019-11-11T18:55:04+09:00",
  "UpdatedAt": "2019-11-11T18:55:53+09:00",
  "DeletedAt": null
}

% curl -X DELETE localhost:8080/api/v1/users/1 | jq .
{
  "id": 1,
  "first_name": "",
  "last_name": "",
  "CreatedAt": "0001-01-01T00:00:00Z",
  "UpdatedAt": "0001-01-01T00:00:00Z",
  "DeletedAt": null
}

% curl localhost:8080/api/v1/users | jq .
[]
```