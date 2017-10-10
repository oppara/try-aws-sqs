# php

## 前準備

```
% composer install
```

アクセスキー、シークレットキー等の設定

```
% cp .env.sample .env
% vi env
```

## 実行

送信処理

```
% php send.php
```

受信 & 削除処理

```
% php recv.php
```
