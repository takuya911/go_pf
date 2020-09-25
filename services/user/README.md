# User Service

## 仕様メモ
- 削除されたユーザーの扱い
  - 一度削除したユーザーの復活は基本しない
    - Emailをuniqueにしようとしている為、処理が複雑になる
      - DBで設定はしていない
  - 削除されたユーザーの取得・更新は基本しない
    - 登録・更新前は基本的にDBから情報をとって、確認してからinsert/updateしているのでinsert/updateはgormで行っている
- 削除されたユーザーに厳しい理由
  - protoで生成したコードをGraphQLで読み込んで処理しており、time型のnull値が混ざるとめんどくさいから。
  - 参照：/interface/controller/transform.go


#### 内部実装の決まり

- Repositoryのsqlについて
  - selectと削除系のupdateは基本生でsql書く。
    - 全てはdeleted_atのせい 