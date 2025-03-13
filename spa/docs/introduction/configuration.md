# 服务配置

## 文件配置方式

支持从 JSON 或 YAML 配置文件倒入配置到 ETCD。

### JSON 配置文件

将如下配置保存到 `ects.json`：

```json
{
  "database": {
    "host": "localhost",
    "port": 3306,
    "name": "ects",
    "user": "root",
    "pass": "PASSWORD",
    "char": "utf8mb4"
  },
  "auth": {
    "secret": "SECRET",
    "ttl": 86400
  },
  "etcd": {
    "killer": "/ects/killer",
    "locker": "/ects/locker",
    "service": "/ects/nodes",
    "pipeline": "/ects/pipelines",
    "config": "/ects/config",
    "endpoints": [
      "localhost:2379"
    ],
    "timeout": 5
  }
}
```
运行如下命令，将配置倒入的 ETCD：

```bash
$ ects init --mode=json \
--path=./ects.json \
--name=User \
--pass=Password \
--email=user@mail.com
```

### YAML 配置文件

将如下配置保存到 `ects.yaml`：
```yaml
database:
  host: localhost
  port: 3306
  name: ects
  user: root
  pass: PASSWORD
  char: utf8mb4
auth:
  secret: SECRET
  ttl: 86400
etcd:
  killer: /ects/killer
  locker: /ects/locker
  service: /ects/service
  pipeline: /ects/pipeline
  config: /ects/config
  endpoints:
    - localhost:2379
  timeout: 5
```
运行如下命令，将配置倒入的 ETCD：
```bash
$ ects init --mode=yaml \
--path=./ects.json \
--name=User \
--pass=Password \
--email=user@mail.com
```

## Web UI 配置方式

### 启动初始化服务

```bash
$ ects init --mode=web
Now listening on: http://localhost:9701
Application started. Press CMD+C to shut down.
```

### 生成配置

#### ETCD 配置
![ETCD](/ects/configuration/etcd.png)

#### JWT 配置

![JWT](/ects/configuration/jwt.png)

#### MySQL 配置

![JWT](/ects/configuration/db.png)

如果数据库已存在则会提示是否覆盖现有数据：
![JWT](/ects/configuration/db_alert.png)

#### 创建管理员

![JWT](/ects/configuration/user.png)

#### 初始化完成

![JWT](/ects/configuration/finish.png)
