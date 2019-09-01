# 架构设计

## 主节点

主节点又称 Master 节点，在整个系统中起到了管理 API 和从节点的状态更新功能。

当有新的节点注册到 ETCD 时，为了避免单个节点注册被多个 Master 节点重复更新数据库的情况。在实现上我们采用了基于 ETCD 的分布式锁，只有抢到锁的 Master 节点才能创建或更新节点信息。

## 从节点

从节点也就是 Worker 节点，此类型的节点启动后便向 ETCD 注册自己的节点信息，然后从 ETCD 中获取需要在本机上运行的定时任务，并开始调度。

当流水线执行完成后，无论成功或失败，都将执行结果保存到 MySQL 数据库中，便于用户排查问题。

## 服务注册

所有节点会根据配置的节点 Key 前缀，向 ETCD 中 PUT 自己的节点信息，默认前缀是 `/ects/nodes` 你可以使用 `etcdctl` 命令，来查看已经注册的节点信息。

```bash
# 使用 ETCD 3 API
$ export ETCDCTL_API=3

# 获取节点 Key 列表
$ etcdctl get /ects/nodes --prefix --keys-only
/ects/nodes/7935d870-99b0-4d32-b799-dd1ec0c1795d

# 获取节点详情
$ etcdctl get /ects/nodes/7935d870-99b0-4d32-b799-dd1ec0c1795d
/ects/nodes/7935d870-99b0-4d32-b799-dd1ec0c1795d
{"id":"7935d870-99b0-4d32-b799-dd1ec0c1795d","name":"George.local","host":"192.168.1.9","port":9701,"mode":"master","status":"online","version":"0.3.0","description":"master node"}
```

解析结果如下：

```json
{
    "id": "7935d870-99b0-4d32-b799-dd1ec0c1795d",
    "name": "George.local",
    "host": "192.168.1.9",
    "port": 9701,
    "mode": "master",
    "status": "online",
    "version": "0.3.0",
    "description": "master node"
}
```

> 节点 ID 为 UUID

## 服务发现

目前用到的服务发现，仅仅是 Master 检测 Worker 的变化。并更新数据库或发送通知给用户。

## 流水线

因为考虑到有些任务需要关联起来，并且相互依赖，所以我们引入了流水线的模式。流水线可以关联多个任务，并排序任务执行顺序。

## 任务

单个任务无法直接在 Worker 上进行调度，Worker 调度的最小单位是流水线。

## 关联关系

如果想在某台 Worker 节点上运行一个定时任务，你必须将任务添加到流水线的执行步骤中，紧接着将流水线绑定到对应的节点上，这样实现的目的是尽可能的避免跨主机的任务重复定义的问题。
