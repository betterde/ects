# 架构设计

![Architecture](/ects/architecture.png)

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

以下是 PUT 到 ETCD 中的流水线数据结构
```json
{
    "id": "b6e54dc3-7e12-4029-b499-fb4603c72163",
    "name": "测试任务流水线",
    "description": "",
    "spec": "*/10 * * * * * *",
    "status": 0,
    "finished": "7a0a46fb-7f61-4b36-8005-a4da351f32fa",
    "failed": "7a0a46fb-7f61-4b36-8005-a4da351f32fa",
    "overlap": 1,
    "created_at": "2019-08-17 12:03:40",
    "updated_at": "2019-09-02 14:32:52",
    "nodes": [
        "1d9988aa-4630-4230-a646-4a2f377d4530"
    ],
    "steps": [
        {
            "id": "2bf41417-a2b6-43c8-8755-9a54c3fa654e",
            "pipeline_id": "b6e54dc3-7e12-4029-b499-fb4603c72163",
            "task_id": "4afe4ea7-67a3-4836-b8ef-3b4c356ff3fd",
            "step": 1,
            "timeout": 0,
            "interval": 0,
            "retries": 0,
            "directory": "",
            "user": "",
            "environment": "",
            "dependence": "strong",
            "created_at": "2019-08-17 12:05:26",
            "updated_at": "2019-08-17 12:05:26",
            "task": {
                "id": "4afe4ea7-67a3-4836-b8ef-3b4c356ff3fd",
                "name": "测试任务",
                "mode": "shell",
                "url": "",
                "method": "post",
                "content": "sleep 3 && echo \"done\"",
                "description": "测试任务",
                "created_at": "2019-08-17 12:02:53",
                "updated_at": "2019-08-28 14:13:06"
            }
        },
        {
            "id": "4379cc1a-ddbf-4940-8ce2-9add3e276abb",
            "pipeline_id": "b6e54dc3-7e12-4029-b499-fb4603c72163",
            "task_id": "d2e3f791-339b-41c4-8b37-10c3bbd9aa34",
            "step": 2,
            "timeout": 0,
            "interval": 0,
            "retries": 0,
            "directory": "",
            "user": "",
            "environment": "",
            "dependence": "strong",
            "created_at": "2019-08-31 15:09:59",
            "updated_at": "2019-08-31 15:09:59",
            "task": {
                "id": "d2e3f791-339b-41c4-8b37-10c3bbd9aa34",
                "name": "查看目录",
                "mode": "shell",
                "url": "",
                "method": "post",
                "content": "sleep 3 && ls -la",
                "description": "查看目录结构",
                "created_at": "2019-08-31 15:09:42",
                "updated_at": "2019-08-31 15:09:42"
            }
        }
    ],
    "finished_task": {
        "id": "7a0a46fb-7f61-4b36-8005-a4da351f32fa",
        "name": "绑定IP",
        "mode": "shell",
        "url": "",
        "method": "post",
        "content": "bind aliyun dns",
        "description": "绑定IP到DNS",
        "created_at": "2019-08-29 17:31:42",
        "updated_at": "2019-08-29 17:31:42"
    },
    "failed_task": {
        "id": "7a0a46fb-7f61-4b36-8005-a4da351f32fa",
        "name": "绑定IP",
        "mode": "shell",
        "url": "",
        "method": "post",
        "content": "bind aliyun dns",
        "description": "绑定IP到DNS",
        "created_at": "2019-08-29 17:31:42",
        "updated_at": "2019-08-29 17:31:42"
    }
}
```
* nodes 对应的是流水线需要在那些 Worker 节点上运行，节点在监听到 PUT 事件以后，根据 `nodes` 里的节点 ID 判断自己是否需要将流水线加入调度列表。
* steps 流水线执行步骤，按照顺序逐个执行，如果中途发生异常则终止整个流水线，并记录日志。
* finished_task 当流水线成功后触发的任务。
* failed_task 当流水线失败后触发的任务。

## 任务

单个任务无法直接在 Worker 上进行调度，Worker 调度的最小单位是流水线。

## 关联关系

![Relations](/ects/pipelines_relations.png)

如果想在某台 Worker 节点上运行一个定时任务，你必须将任务添加到流水线的执行步骤中，紧接着将流水线绑定到对应的节点上，这样实现的目的是尽可能的避免跨主机的任务重复定义的问题。

上图中，`task2` 可以在 `pipeline1` 中被调度执行，也可以在 `pipeline2` 中别调度执行，而 `pipeline1` 又绑定到了 `worker1`、`worker2`、`worker3` 节点上，也就是说三个节点上都会调度执行 `pipeline1` 下的任务。
