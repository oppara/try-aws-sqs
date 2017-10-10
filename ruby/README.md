# ruby

## 前準備

```
% bundle install --path vendor/bundle
```

アクセスキー、シークレットキー等の設定

```
% cp .env.sample .env
% vi env
```

## 実行

送信処理

```
% bundle exec ruby send.rb 
```

受信 & 削除処理

```
% bundle exec ruby recv.rb 
```
