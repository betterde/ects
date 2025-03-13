# RESTful API

## 用户认证

### 登录

```http request
POST https://{{host}}/api/auth/signin
```
参数
```json
{
	"username": "george@betterde.com",
	"password": "chopin"
}
```
响应
```json
{
    "code": 200,
    "message": "登录成功",
    "data": {
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Njc0OTgyNzcsImlhdCI6MTU2NzQxMTg3NywiaXNzIjoiZWN0cyIsIm5iZiI6MTU2NzQxMTg3Nywic3ViIjoiMzM1NWExZDAtMTJjOS00MmE0LWE3ZDYtYmRjZjdkNDY5NjQzIn0.vMwFfIyH_CJcbbiO2-uRYiTDoaqWZ8zgiN5gGjdUhRY",
        "token_type": "Bearer",
        "expires_in": 86400
    }
}
```

## 概览

## 节点管理

## 任务管理

## 流水线管理

## 系统设置

## 日志查询

## 用户管理
