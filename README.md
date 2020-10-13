# Golang Portfolio
## 概要
ポートフォリオ&勉強用に作ったリポジトリです。
特にこれを運用していきたいという目的はありません。
以下の技術を勉強するために、無駄に盛り込んでいます。
- Golang
- gRPC
- GraphQL

## 起動方法
- 本番
  1. terraform applyする
    - 起動順序の問題で失敗することがある。失敗したら再実行でok。
    - /terraform/variable.tfは適宜修正する。

  2. mysqlのtable作成を行う。
    - /infrastructure/mysql/README.mdを参照。
- 開発環境
  - kubernetes
    1. minikubeがinstall&`minikube start`を実行
    2. /kubernetes/devにて`sh apply.sh`実行
  - docker-compose
    1. `docker-compose up --build`を実行

## 注意点
現状、本番は動かしてません。
そのため、死を招くレベルではないcredential情報は余裕でgit載せてます。

## 使用技術

### application
- Golang
- Protocol Buffers
- GraphQL
  - gqlgen
- Mysql
- JWT

### infrastructure
- dev
  - Docker
  - docker-compose
  - Kubernetes
    - minikube
  - docker hub

- prod
  - GCP
    - Cloud Run
    - Container Registory
    - Cloud SQL
    - IAM
  - Terraform

### Tools
  - Github
  - Github Actions