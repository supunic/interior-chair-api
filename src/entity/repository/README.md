# entity / repository

## レイヤー
Enterprise Business Rules

## 概念
- Entities
  - Repository Interface

## 内容
`entity / service`で実装した、

- ドメインに持たせるべきではない処理
- ドメインモデルの集合に対する処理
- ドメイン集約のトランザクション処理

のインターフェース

ここで定義する場合、基本となる`adapter -> usecase -> entity`の依存関係が、
`usecase`を飛ばした`adapter -> entity`になることを妥協する

参考
[【Go】厳密なClean Architectureとその妥協案 - Qiita](https://qiita.com/ariku/items/659a11767912c2ec266d)

## 実装
- リポジトリインターフェース
