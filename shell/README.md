# shell


## 前準備

jqをインストールしとく

```
% brew install jq
```

アクセスキー、シークレットキー等の設定

```
% cp .env.sample .env
% vi env
```

## 実行

送信処理

```
% ./send.php
```

受信 & 削除処理

```
% ./recv.php
```
