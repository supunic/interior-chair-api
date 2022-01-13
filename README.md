# Interior Chair Api

## Description
- インテリアで有名なイスの情報を提供するAPI
- オシャレ椅子に詳しくなりたい人向け

## Example
wip

## Dependencies
- go v1.17
- echo v4.6
- gorm v2.0
- mysql v8.0
- docker v4.1

## Setup
```shell
make start
```

## Usage
```shell
# goコンテナに入る
make shell

# mysqlクライアント実行（password：root）
make mysql
```
その他はMakefile参照

## Directory Structure
```text
src
├── adapter
│   ├── controller
│   ├── gateway
│   └── presenter
├── driver
│   ├── db
│   └── web
├── entity
│   ├── model
│   ├── repository
│   └── service
├── usecase
│   ├── interactor
│   └── port
└── main.go
```
詳細は各ディレクトリのREADME参照

### Reference
- [Go言語とClean ArchitectureでAPIサーバを構築する -Qiita](https://qiita.com/ariku/items/874656b33d2e5acdf281)
- [【Go】厳密なClean Architectureとその妥協案 - Qiita](https://qiita.com/ariku/items/659a11767912c2ec266d)

