# 运行服务

## 命令行介绍

### master

```bash
$ ects master -h
Run a master node service on this server

Usage:
  ects master [flags]

Flags:
      --config string   Set the key used to get configuration information (default "/ects/config")
      --desc string     Set master node description (default "master node")
      --etcd strings    Set Etcd endpoints (default [127.0.0.1:2379])
  -h, --help            help for master
      --host string     Set listen on IP (default "0.0.0.0")
      --name string     Set master node name
  -n, --node string     Set master node id (default "b2567790-0f8d-458e-8560-6dff90e8c4ad")
      --port int        Set listen on port (default 9701)
```

* config：系统配置信息在 ETCD 中的 Key
* desc：节点的描述
* etcd：etcd 的 endpoints，用英文逗号隔开
* host：服务监听的地址
* name：节点名称
* node：节点ID，如果未提供节点ID，则自动生成ID，并注册到 MySQL
* port 服务监听端口

### worker

```bash
$ ects worker -h
Run a worker node service on this server

Usage:
  ects worker [flags]

Flags:
      --config string   Set the key used to get configuration information (default "/ects/config")
      --desc string     Set worker node description (default "worker node")
      --etcd strings    Set Etcd endpoints (default [127.0.0.1:2379])
  -h, --help            help for worker
      --name string     Set worker node name
  -n, --node string     Set node id
```

* config：系统配置信息在 ETCD 中的 Key
* desc：节点的描述
* etcd：etcd 的 endpoints，用英文逗号隔开
* name：节点名称
* node：节点ID，如果未提供节点ID，则自动生成ID，并注册到 MySQL

## 启动 master 节点

```bash
$ ects master \
--name=Master \
--desc="master node" \
--config=/ects/config \
--etcd=127.0.0.1:2379 \
--host=192.168.1.253 \
--port=9701 \
--node=24b29238-86bb-4cf7-a52a-be009d768c84
```

## 启动 worker 节点

```bash
$ ects worker \
--name=Master \
--desc="worker node" \
--config=/ects/config \
--etcd=127.0.0.1:2379 \
--node=24b29238-86bb-4cf7-a52a-be009d768c84
```

## 运行单机模式

```bash
$ ects single
```
