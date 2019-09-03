# 管理任务

## 概念介绍

### 任务

在 ECST 中，任务是可以公用的，只需要定义任务模式和内容就可以了。

### 流水线

流水线在系统中是最小的调度单位，一个流水线可以在多台 Worker 节点上调度。要想运行流水线，你需要为流水线添加需要执行的任务。流水线中至少包含一个任务，否者无法同步到 ETCD 也就无法被调度。如果你添加了多个任务，则可以自定义执行顺序。你还可以为流水线设置`完成时`和`失败时`执行的任务。可以用来发送`邮件通知`、`Hook`或`HTTP请求`。

### 节点

节点的添加方式有两种，一种是在远端直接运行未提供 `--node` 参数的命令，另一种方式是在后台创建，然后获取节点 ID，然后在远端运行时，提供 `--node` 参数。当为流水线添加了任务后，你还需要将流水线绑定到需要运行的节点上，这样才能被节点调度。

## 创建任务

![Task List](/task/create_task.png)

![Task](/task/list.png)

## 创建流水线

![Task List](/pipeline/create_pipeline.png)

![Task](/pipeline/list.png)

## 绑定节点

![Task List](/pipeline/bind_node.png)

![Task](/pipeline/node.png)

::: tip 注意
只有 Worker 节点才能绑定流水线
:::

## 用户管理

![User](/user.png)

## 日志查询

![Log](/log.png)
