
# 02 Recipe Api

チーム02のバックエンドリポジトリ

## Using

- Golang
  - go version go1.20.2
  - 確認の際は `go version`

- Docker
  - docker version 20.10.22
  - 確認の際は `docker version`
  - Docker内で使用
    - MySQL
    - phpMyAdmin

- docker compose
  - Docker Compose version v2.15.1
  - 確認の際は `docker compose version`

## Git Clone

`git clone https://github.com/qin-team-recipe/02-recipe-api.git [任意のファイルパス]`

## Environment Building

- .envをルートディレクトリに作成
- .env.exampleを参照、またはコピペし入力してください。
- ./app/config/config.goで.envを読み込まれ環境変数が参照されます。

```.env
SERVER_PORT=使用の環境に合わせてください
CONTAINER_SERVER_PORT=8080
ENV=development

DB_ROOT_PASS=任意のルートユーザーのパスワードを入力
DB_NAME=任意のデータベース名を入力
DB_USER=任意のユーザー名を入力
DB_PASS=任意のパスワードを入力
DB_HOST=mysql  #<-docker-composeでコンテナ起動させるので固定
DB_PORT=3306
```

## Running

`make up`

- `docker-compose up`コマンドが実行されます。

設定ファイルは`./docker-compose.yml`に記載しています。

Webサーバー、phpMyAdmin、MySQLがコンテナ内で起動します。

基本的には上記コマンドで一括してコンテナを起動している状態を想定しています。

## Hot Reload

[Air](https://github.com/cosmtrek/air)をDocker内で使用しています。

Goファイルに更新があった場合は自動でビルドし再度立ち上がります。

## Directory Structure

### 主なディレクトリ構成

| ディレクトリ名  | 内容 |
| ------------- | ------------- |
| app  | Golangで記述しているファイルを配置  |
| mysql  | データベース関連のものを配置  |
| mysql/config  | データベースの設定ファイル  |
| mysql/migrations  | Docker起動時に初期化させたいファイルを配置  |
| mysql/sql  | SQLファイルを配置し、mysqlにアクセスすると実行できるファイル  |

### /app ディレクトリ内の構成

| ディレクトリ名  | 内容 |
| ------------- | ------------- |
| cmd  | main.goのみを配置  |
| config  | 環境変数ファイル  |
| constants  | 定数ファイル（あまり使わないかも）  |
| docs  | ドキュメント類  |
| internal  | アプリケーションで使用するコードを配置(ここから外部へは参照しない)  |
| pkg  | 自作パッケージを配置(utilitiesのようなイメージ)  |

### /internal ディレクトリ内の構成

| ディレクトリ名  | 内容 |
| ------------- | ------------- |
| domain  | データを永続化する層  |
| infrastructure  | インフラ層  |
| interface  | インターファイス層  |
| interface/controllers  | パラメータを受け取りusecaseへ接続する  |
| interface/gateways  | usecaseとインフラ層を繋ぐ  |
| interface/gateways/repository  | Gormのメソッドの記述などをここに配置  |
| interface/presenters  | レスポンスをフロントの扱いやすいように整生する  |
| usecase  | ユースケース配置  |
| usecase/repository  | interfaceの役割（DI）  |
| usecase/interactor  | ビジネスロジックを配置  |
