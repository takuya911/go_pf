# Prod環境の初期処理

- instanceに接続
```
$ gcloud sql connect go-pf-instance --user=user
```

- DBを選択
```
$ USE {database name}
```

- table作成
  - /init配下のsqlを実行していく
