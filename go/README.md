# golang

## 前準備

```
% glide install
```

アクセスキー、シークレットキー等の設定

```
% cp .env.sample .env
% vi env
```

## 実行

送信処理

```
% go run main.go send
```

受信 & 削除処理

```
% go run main.go recv
```
