# entity / service

## レイヤー
Enterprise Business Rules

## 概念
- Entities
  - Domain Service

## 内容
- ドメインに持たせるべきではない処理
- ドメインモデルの集合に対する処理
- ドメイン集約のトランザクション処理

の実装

`repository` + α のことをしたいときに利用する

参考
[【Go】厳密なClean Architectureとその妥協案 - Qiita](https://qiita.com/ariku/items/659a11767912c2ec266d)

## 実装
- ドメインサービス
