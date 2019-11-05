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
    - mysql : `go get -u github.com/go-sql-driver/mysql`
- Logger : [zap](https://github.com/uber-go/zap)
  - [gin と連携](https://github.com/gin-contrib/zap)
- Env : [envconfig]](https://github.com/kelseyhightower/envconfig) 、 [dotenv](https://github.com/joho/godotenv)
- Security : [secure](https://github.com/unrolled/secure)
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

## テスト

```zsh
% go test -v github.com/pepese/golang-gin-sample/...
```