# 目录结构

## 主目录

```bash
.
├── LICENSE
├── README.md
├── cmd # 所有命令
├── config # 系统配置
├── controllers # 控制器
├── docs
├── ects.example.json # JSON 格式配置文件模板
├── ects.example.yaml # YAML 格式配置文件模板
├── go.mod
├── go.sum
├── internal # 内部公用模块
├── main.go # 入口文件
├── models # 数据模型
├── routes # RESTful 路由
├── services # 数据服务，后期考虑移除
├── vendor # 依赖的包
└── web # Web UI

10 directories, 7 files
```

## 后端框架

经过对多个框架的对比，最终决定选择 [iris](https://iris-go.com)。

### 命令行

命令行采用的是 [Cobra](https://github.com/spf13/cobra)，这个包提供了比较全面的命令行常用功能。可以看到包括像 Kubernetes、Docker、Etcd、Linkerd、Istio 和 Pouch 这样的项目都在使用它。

```bash
├── cmd
│   ├── initialize.go # 初始化命令
│   ├── master.go # master 节点启动命令
│   ├── root.go
│   ├── single.go # 单节点启动命令，可以不需要 ETCD
│   └── worker.go # worker 节点启动命令
```

### 路由

```bash
├── routes
│   ├── auth.go # 用户认证
│   ├── dashboard.go # 概览
│   ├── initialize.go # 初始化
│   ├── log.go # 日志
│   ├── node.go # 节点
│   ├── pipeline.go # 流水线
│   ├── profile.go # 个人信息
│   ├── register.go # 用户注册
│   ├── setting.go # 系统设置
│   ├── task.go # 任务
│   └── user.go # 用户管理
```

### 控制器

```bash
├── controllers
│   ├── account # 账户
│   │   └── profile.go
│   ├── auth # 认证
│   │   └── authentication.go
│   ├── dashboard # 概览
│   │   └── main.go
│   ├── initialize # 初始化
│   │   └── main.go
│   ├── log # 日志
│   │   └── main.go
│   ├── node # 节点
│   │   └── main.go
│   ├── organization # 组织
│   │   └── user.go # 用户
│   ├── pipeline # 流水线
│   │   └── main.go
│   ├── setting # 设置
│   │   └── main.go
│   └── task # 任务
│       └── main.go
```

### 模型

```bash
── models
│   ├── log.go # 日志
│   ├── main.go # MySQL 配置信息
│   ├── node.go # 节点
│   ├── password_reset.go # 密码重置
│   ├── pipeline.go # 流水线
│   ├── pipeline_node_pivot.go # 流水线和节点关联关系模型
│   ├── pipeline_record.go # 流水线执行记录模型
│   ├── pipeline_task_pivot.go # 流水线和任务关联模型
│   ├── task.go # 任务
│   ├── task_record.go # 任务执行记录
│   └── user.go # 用户
```

### 开发环境

* MySQL 8.0.11
* ETCD 3.3.12

```bash
# 启动 MySQL
$ docker run -d \
    --name mysql \
    -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=your-secret-pw mysql:8.0.11 \
    --character-set-server=utf8mb4 \
    --collation-server=utf8mb4_unicode_ci

# 启动 ETCD
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

### 初始化

这一步主要是将配置信息写入到 ETCD 中

```bash
$ go run main.go init -h
Run initialize elastic crontab system service

Usage:
  ects init [flags]

Examples:
ects init

Flags:
  -e, --email string   Set admin email # 如果是采用配置文件方式初始化，则需要提供管理员邮箱
  -h, --help           help for init
  -m, --mode string    Set initialize mode with web ui or json, yaml config file (default "web") # 支持的模式有 web、json、yaml
  -n, --name string    Set admin name # 如果是采用配置文件方式初始化，则需要提供管理员用户名
  -P, --pass string    Set admin pass # 如果是采用配置文件方式初始化，则需要提供管理员密码
  -p, --path string    Set config file path # 如果是采用配置文件方式初始化，则需要提供配置文件路径
```

### 打包前端资源到二进制

在打包之前，请参照前端框架的内容，先眼妆依赖，并打包前端资源。

```bash
# install go-bindata
$ go get -u github.com/shuLhan/go-bindata/...

# 打包静态资源
$ go-bindata -pkg web -o web/bindata.go web/dist/...
```

### 启动服务

```bash
$ go run main.go master
```

## 前端框架

前端框架我们采用的是 `Vue Cli 3.*` 搭建的 SPA，以及基于 Vue 而开发的组件库 [Element-UI](https://element.eleme.cn/)。

```bash
.
├── README.md
├── babel.config.js
├── bindata.go # 前端静态资源打包成二进制后的文件
├── dist # 执行 yarn build 后的输出目录
├── docs # 项目文档目录，采用 Vuepress 搭建
│   ├── README.md
│   ├── developer
│   │   ├── README.md
│   │   └── api.md
│   ├── introduction
│   │   ├── README.md
│   │   ├── architecture.md
│   │   ├── configuration.md
│   │   ├── dependencies.md
│   │   ├── managerment.md
│   │   ├── more.md
│   │   └── services.md
│   └── version
│       └── README.md
├── package.json
├── public
│   ├── favicon.ico
│   └── index.html
├── src
│   ├── App.vue # 根组件
│   ├── apis # 与后端交互的 API
│   │   ├── index.js
│   │   └── modules # API 模块
│   ├── assets # 前端静态资源
│   │   └── styles
│   ├── components # 自定义公共组件
│   │   ├── Collapse.vue
│   │   └── CronExpression.vue
│   ├── element-variables.scss
│   ├── entries # 入口文件
│   │   ├── index.js
│   │   └── initialize.js
│   ├── language # 本地化
│   │   ├── en.js
│   │   └── index.js
│   ├── layouts # 页面布局
│   │   ├── Backend.vue # 登录后的布局
│   │   └── Guest.vue # 未登录时的布局
│   ├── plugins # Vue cli 3.* 的插件目录
│   │   ├── axios.js # 发起 XHR 请求的插件
│   │   └── element.js # Element-UI 插件
│   ├── router
│   │   └── index.js # 前端路由
│   ├── store # vuex 目录
│   │   ├── actions.js
│   │   ├── index.js
│   │   ├── modules # 模块
│   │   ├── mutations.js
│   │   └── types.js
│   └── views # 模块页面
│       ├── About.vue # 关于
│       ├── Dashboard.vue # 概览
│       ├── Initialize.vue # 初始化配置
│       ├── Log.vue # 日志查询
│       ├── Node.vue # 节点管理
│       ├── NotFound.vue # 404 所有前端为定义的路由将跳转至该页面
│       ├── Pipeline.vue # 流水线管理
│       ├── Setting.vue # 系统设置
│       ├── SignIn.vue # 用户登录
│       ├── Task.vue # 任务管理
│       └── User.vue # 系统用户管理
├── vue.config.js # 项目配置，代理、HTTPS
├── yarn-error.log
└── yarn.lock

24 directories, 93 files
```
### 安装依赖

```bash
$ cd web
$ yarn install
$ yarn build
```

### 项目配置

前端项目的打包配置都在 `web/vue.config.js` 文件中，其中 `devServer` 配置根据自己的需要进行设置。
