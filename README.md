# mongodb-with-golang
MongoDB  Quick Start with Golang

```sh
  # 起動コマンド
$ docker-compose up -d
  # 停止させるコマンド
$ docker-compose down
```


https://www.mongodb.com/docs/drivers/go/current/quick-start/

https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo



* bson.M（BSON Map）:

MongoDBのドキュメントを表現するためのマップ（連想配列）の一種
キーと値のペアを含むデータを表現
キーは文字列で、値はさまざまなデータ型になる
```go
doc := bson.M{
    "name":  "hogehoge",
    "age":   30,
    "email": "hogehoge@example.com",
}
```

* bson.D（BSON Document）:

MongoDBのドキュメントを表現するためのドキュメント（ディクショナリ）の一種
キーと値のペアを持ち、挿入順に格納
キーは文字列で、値はさまざまなデータ型になる
bson.M と異なり、挿入順が保持されるため、順序が重要な場合に使用
```go
doc := bson.D{
    {"name", "hogehoge"},
    {"age", 30},
    {"email", "hogehoge@example.com"},
}
```