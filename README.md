![dashboard](web/docs/.vuepress/public/logo.png)

[![Build Status](https://github.com/betterde/ects/actions/workflows/ci.yml/badge.svg)](https://github.com/betterde/ects/actions/workflows/ci.yml/badge.svg)
[![GitHub issues](https://img.shields.io/github/issues/betterde/ects)](https://github.com/betterde/ects/issues)
[![GitHub forks](https://img.shields.io/github/forks/betterde/ects)](https://github.com/betterde/ects/network)
[![GitHub stars](https://img.shields.io/github/stars/betterde/ects)](https://github.com/betterde/ects/stargazers)
[![GitHub license](https://img.shields.io/github/license/betterde/ects)](https://github.com/betterde/ects/blob/master/LICENSE)

Elastic Crontab System 是基于 ETCD 实现的分布式定时任务管理系统，让跨主机的定时任务管理变得更加简单高效。

## 架构设计

![dashboard](web/docs/.vuepress/public/architecture.png)

## 源码安装

```bash
# 克隆到本地
$ git clone git@github.com:betterde/ects.git

# 安装前端依赖
$ cd web && yarn install

# 打包前端资源
$ yarn build

# 编译
$ go build -o ects main.go
```

## 下载可执行程序

[下载地址](https://github.com/betterde/ects/releases)

## 了解更多细节

[文档地址](https://betterde.github.io/ects/)

## TODO

- [x] Web UI；
- [x] Master 节点 API；
- [x] 基于 ETCD 的服务注册于发现；
- [x] 基于 ETCD 的流水线发布于订阅；
- [x] 基于 ETCD 实现的分布式锁，用于更新 Worker 节点的状态；
- [x] 实现 HTTP 任务执行器；
- [x] 实现 Hook 任务执行器；
- [x] 集成 Docker 部署；
- [x] 集成 CI；
- [x] 项目文档；
- [] 角色权限管理模块。

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=betterde/ects&type=Date)](https://star-history.com/#betterde/ects&Date)

## 赞助商

[![jetbrains](web/docs/.vuepress/public/jetbrains.svg)](https://www.jetbrains.com/?from=ects)

