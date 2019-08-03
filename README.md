# ECTS

Elastic Crontab System

## Architecture

![dashboard](docs/overview/architecture.jpg)

### Master 
* RESTful API Server
* Web UI

### Worker
* Shell task actuator

## Installation

```bash
go get github.com/betterde/ects

# install frontend dependencies
cd web && yarn install

# build frontend asset
yarn build

# install go-bindata
cd ../ && go get -u github.com/shuLhan/go-bindata/...

# package resource file
go-bindata -pkg web -o web/bindata.go web/dist/...

go build -o ects main.go

```

## Initialization

### Run initialization web service
```bash
$ ects init -m web

Now listening on: http://localhost:9701
Application started. Press CMD+C to shut down.
```

#### Open your browser

Point your browser to http://localhost:9701

#### Config ETCD

![etcd](docs/overview/initialization/init-etcd.png)

#### Config JWT

![jwt](docs/overview/initialization/init-jwt.png)

#### Config DB

![db](docs/overview/initialization/init-db.png)

#### Config User

![user](docs/overview/initialization/init-user.png)

#### Complete

![finish](docs/overview/initialization/init-complete.png)

### Init with config file
```bash
ects init -m json -p CONFIG_FILE_PATH -n ADMIN_NAME -P ADMIN_PASS -e ADMIN_MAIL
ects init -m yaml -p CONFIG_FILE_PATH -n ADMIN_NAME -P ADMIN_PASS -e ADMIN_MAIL
```

## Web UI

### Pipeline list

![pipeline](docs/overview/pipeline.png)

### Task list

![pipeline](docs/overview/task.png)

### Node list

![pipeline](docs/overview/node.png)

### User list

![pipeline](docs/overview/user.png)

### User log

![pipeline](docs/overview/user-log.png)

## LICENSE
ECTS is open-sourced software licensed under the MIT license.