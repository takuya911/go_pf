# Prod環境の初期処理

- instanceに接続
  - Web コンソール or Set Upが済んでいればlocal PCからも実行可能
  - passwordはterraformに書いてあるやつ。
```
$ gcloud sql connect {db instance} --user=user
```

- DBを選択
  - database nameはterraformに書いてあるやつ。
```
$ USE {database name}
```

- table作成
  - /init配下のsqlを実行していく

