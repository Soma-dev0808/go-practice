# 　コマンド一覧

## 立ち上げ

- docker-compose up -d

## mysql アクセス

- mysql -h 127.0.0.1 -P 3307 -u docker sampledb -p

## main.go 実行

- go run main.go

## main.go 環境変数を用いて実行

- DB_USER=docker DB_PASSWORD=docker DB_NAME=sampledb go run main.go

## repository test 実行

- cd repositories
- go test -v
