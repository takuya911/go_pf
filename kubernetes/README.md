# Kubernetes

## 便利コマンドメモ

- service type loadbalancer or node portのエンドポイント取得

```
$ minikube service {service name} --url
```


- terminalからkube上のmysql接続

```
$ mysql -uroot -p -h {service ip} --port {port}
```