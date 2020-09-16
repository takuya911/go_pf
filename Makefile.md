## proto file

- コード生成コマンド(repository root)
```
$ make pb-gen service=${service}
```
- 叩けない時はpathがおかしくなっていることがある
```
$ export PATH=$PATH:$GOPATH/bin
```