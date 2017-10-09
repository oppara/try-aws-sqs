# golang

```
$ glide install
```

アクセスキー、シークレットキーの設定

```
$ cp .env.sample .env
$ vi env
```

送信処理

```
$ go run main.go send
```

受信 & 削除処理
```
$ go run main.go recv
```
