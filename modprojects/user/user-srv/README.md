# User Service

This is the User service

Generated with

```
micro new user-srv --namespace=bambooRat.micro.sdk --alias=user --type=srv --gopath=false
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: bambooRat.micro.sdk.srv.user
- Type: srv
- Alias: user

## Dependencies
### consul
推荐用consul进行服务发现检测，如果没有安装consul可以执行[tools/framework_install.sh][framework_install.sh]

```
# mac install consul
brew install consul

# run consul
consul agent -dev
```
### mysql
推荐直接docker 拉取镜像安装自己需求的版本
```
# docker 中下载 mysql 现在默认是8.x 如果想下载5.x可以显式指定mysql5.6
docker pull mysql

#启动 启动的端口和密钥自行设置
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql

#进入容器
docker exec -it mysql bash

#登录mysql
mysql -u root -p
ALTER USER 'root'@'localhost' IDENTIFIED BY '123456';

#添加远程登录用户
CREATE USER 'guest_user'@'%' IDENTIFIED WITH mysql_native_password BY '123456';
GRANT ALL PRIVILEGES ON *.* TO 'guest_user'@'%';
```

## Usage

Makefile负责把 main包编译成二进制文件 

```
make build
```

启动服务
```
./user-srv
```

创建一个docker镜像
```
make docker
```

[framework_install.sh]: https://github.com/xuyiwenak/bambooRat/blob/master/tools/framework_install.sh