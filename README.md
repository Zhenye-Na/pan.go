# pan.golang

> Go 实战仿百度云盘 实现企业级分布式云存储系统

## 关于分支和标签

- 分支说明:
    - master分支: 通过原生net/http实现各接口功能
    - gin分支: 通过Gin框架来改造(微服务化章节之后主要基于Gin框架来进行演示)
    - pkg分支: 通过go-bindata等工具实现静态资源打包的一个示例

- 标签
    - tag是分支代码某个阶段的快照。如 `master` 主分支的 `v0.1` 是首个发布的版本, 该版本是接入了阿里 OSS


## 关于需要手动安装的库

如下:

```shell
go get github.com/garyburd/redigo/redis
go get github.com/go-sql-driver/mysql
go get github.com/garyburd/redigo/redis
go get github.com/json-iterator/go
go get github.com/aliyun/aliyun-oss-go-sdk/oss
go get gopkg.in/amz.v1/aws
go get gopkg.in/amz.v1/s3
go get github.com/streadway/amqp
```
其中如果有提示`golang.org/x`相关的包无法下载的话，可以参考这篇文章:

[国内下载golang.org/x/net](https://yq.aliyun.com/articles/292301?spm=5176.10695662.1996646101.searchclickresult.6155183eCmXHbQ)

## 关于应用启动

- 在加入rabbitMQ实现文件异步转移之前，启动方式：

    - 启动上传应用程序:

        ```bash
        # cd $GOPATH/<你的工程目录>
        > cd $GOPATH/filestore-server
        > go run main.go
        ```

- 在加入rabbitMQ实现文件异步转移阶段，启动方式(分裂成了两个独立程序)：

    - 启动上传应用程序:

        ```bash
        # cd $GOPATH/<你的工程目录>
        > cd $GOPATH/filestore-server
        > go run service/upload/main.go
        ```

    - 启动转移应用程序:

        ```bash
        # cd $GOPATH/<你的工程目录>
        > cd $GOPATH/filestore-server
        > go run service/transfer/main.go
        ```

## 进度说明：

* [x] 简单的文件上传服务
* [x] MySQL 存储文件元数据
* [x] 账号系统, 注册/登录/查询用户或文件数据
* [x] 基于帐号的文件操作接口
* [x] 文件秒传功能
* [x] 文件分块上传/断点续传功能
* [x] 搭建及使用 Ceph 对象存储集群
* [x] 使用阿里云 OSS 对象存储服务
* [x] 使用 RabbitMQ 实现异步任务队列
* [ ] 微服务化 (API网关, 服务注册, RPC通讯)
* [ ] CI/CD (持续集成)

## 参考资料

- Go入门: [语言之旅](https://tour.go-zh.org/welcome/1)
- MySQL: [偶然翻到的一位大牛翻译的使用手册](https://chhy2009.github.io/document/mysql-reference-manual.pdf)
- Redis: [命令手册](http://redisdoc.com/)
- Ceph: [中文社区](http://ceph.org.cn/) [中文文档](http://docs.ceph.org.cn/)
- RabbitMQ: [英文官方](http://www.rabbitmq.com/getstarted.html) [一个中文版文档](http://rabbitmq.mr-ping.com/)
- 阿里云OSS: [文档首页](https://help.aliyun.com/product/31815.html?spm=a2c4g.750001.3.1.47287b13LQI3Ah)
- gRPC: [官方文档中文版](http://doc.oschina.net/grpc?t=56831)
- k8s: [中文社区](https://www.kubernetes.org.cn/docs)
- 我的手记博客 [http://www.imooc.com/u/6198190](http://www.imooc.com/u/6198190)

