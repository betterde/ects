# 环境依赖

## ETCD

考虑到性能，这里采用的是 `ETCD 3.3.12` 版本，如果只有单机环境可以直接运行 `etcd single` 命令来部署单节点服务。

### 安装

```bash
$ docker run -d \
    --name etcd \
    -p 2379:2379 \
    -p 2380:2380 \
    --volume=/private/var/local/etcd/data:/data \
    gcr.io/etcd-development/etcd:v3.3.12 \
    /usr/local/bin/etcd \
    --name etcd \
    --data-dir /data \
    --listen-client-urls http://0.0.0.0:2379 \
    --advertise-client-urls http://0.0.0.0:2379 \
    --listen-peer-urls http://0.0.0.0:2380 \
    --initial-advertise-peer-urls http://0.0.0.0:2380 \
    --initial-cluster etcd=http://0.0.0.0:2380 \
    --initial-cluster-token betterde \
    --initial-cluster-state new
```

## MySQL

开发环境中 MySQL 采用的是 8.0.11，所以至少需要 7.* 及以上的 MySQL 版本。 

### 安装

```bash
$ docker run -d \
    --name mysql \
    -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=your-secret-pw mysql:8.0.11 \
    --character-set-server=utf8mb4 \
    --collation-server=utf8mb4_unicode_ci
```
