## Services
- アプリケーションのbackend service


### Graphql
- 他serviceへのproxy用api gateway

### User
- user情報を扱うservice

#### 細かい仕様
- 削除されたユーザーの扱い
  - 一度削除したユーザーの復活は基本しない
    - Emailをuniqueにしているため、処理が複雑になる
  - 削除されたユーザーの取得・更新は基本しない