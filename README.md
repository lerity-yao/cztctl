# cztctl

一个依赖于go-zero的服务代码生成工具

目前支持生成:

- rabbitmq 消费者服务端代码

## 使用方式

### 安装工具

```go
GOPROXY=https://goproxy.cn/,direct
go install github.com/lerity-yao/cztctl@latest
```

### 配置环境

将$GOPATH/bin中的 cztctl 添加到环境变量

## 命令

### cztctl

| 命令        | 简写 | 描述       |
|-----------|----|----------|
| --version | -v | 查看当前版本   |
| --help    | -h | 查看帮助提示信息 |

```shell
## 执行命令 cztctl -v
# cztctl -v
cztctl version 0.0.8 windows/amd64

```

```
## 执行命令 cztctl -h
# cztctl -h  
A cli tool to generate go-zero api service ect...

Usage:
  cztctl [command]

Available Commands:
  completion        Generate the autocompletion script for the specified shell
  go                Generate Go source files
  help              Help about any command

Flags:
  -h, --help      help for cztctl
  -v, --version   version for cztctl


Use "cztctl [command] --help" for more information about a command.
```


### cztctl go

生成 go 语言代码

| 命令       | 简写 | 描述                 |
|----------|--|--------------------|
| go       |  | 示例，并不会生成任何文件       |
| rabbitmq |  | 生成rabbitmq消费者服务端代码 |


```shell
## 执行 cztctl go
# cztctl go   
Done.
```

```
## 执行 cztctl go -h
# cztctl go -h         
Generate Go source files

Usage:
  cztctl go [flags]
  cztctl go [command]

Available Commands:
  rabbitmq    Generate Go source files for a RabbitMQ consumer service from a .rabbitmq definition file using the go-zero framework.

Flags:
  -h, --help          help for go
      --home string   The cztctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority


Use "cztctl go [command] --help" for more information about a command.
```

```shell
## 执行命令 cztctl go rabbitmq -h
# cztctl go rabbitmq -h
Generate Go source files for a RabbitMQ consumer service from a .rabbitmq definition file using the go-zero framework.

Usage:
  cztctl go rabbitmq [flags]

Flags:
      --branch string     The branch of the remote repo, it does work with --remote
      --dir string        The target dir, source files will be output here
  -h, --help              help for rabbitmq
      --home string       The cztctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
      --rabbitmq string   The .rabbitmq file
      --remote string     The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
                          The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure
      --style string      The file naming format, see [https://github.com/zeromicro/go-zero/blob/master/tools/goctl/config/readme.md] (default "goCzt")
      --test              Generate test files
      --type-group        Generate type group files
```

### cztctl go rabbitmq

生成 go 语言 rabbitmq 消费者服务代码

| 命令 | 简写 | 描述                                   |
|--|--|--------------------------------------|
| --home |  | 本地模板位置, 如果没有输入home，则使用默认模板           |
| --rabbitmq |  | 生成rabbitmq的api文件位置                   |
| --dir |  | 生成的代码输出的位置                           |
| --remote |  | 远程模板位置，是一个可以用git拉取的项目地址，优先级最高，高于home |
| --branch |  | 远程模板的分支                              |
| --style |  | 生成的代码命名风格，默认是cztCtl驼峰格式              |

```shell
# 使用默认模板命令
cztctl go rabbitmq --rabbitmq demoA.api -dir ./bf
```

```shell
# 使用本地模板命令
cztctl go rabbitmq --rabbitmq demoA.api -dir ./bf --home ./deploy

```

```shell
# 使用远程模板命令
cztctl go rabbitmq --rabbitmq ./demoA.api --dir ./zf --remote https://{username}:{pwd}@codeup.aliyun.com/5ff662d499fffffb1c5f96c0/ystz/goctl-template.git

# 使用远程模板指定分支命令
cztctl go rabbitmq --rabbitmq ./demoA.api --dir ./zf --remote https://{username}:{pwd}@codeup.aliyun.com/5ff662d499fffffb1c5f96c0/ystz/goctl-template.git --branch dev-yaox
```

生成的代码目录结构如下:

```
\---zf
    +---etc
    \---internal
        +---config
        +---handler
        |   +---demoA
        |   \---demoB
        +---logic
        |   +---demoA
        |   \---demoB
        +---svc
        \---types
```

### rabbitmq api文件语法

type 语法完全跟 go-zero 一致

@server 中只能写 group

service 中@listener消费者名称

@listener下一行代表队列命名，可以写多个，每一个用/开头，空格隔开多个

service名称必须全局唯一

示例

```api
type (
    // 用户结构体
   User {
        Name string `json:"name"` // 用户名称
   }

   Name {
        Code int `json:"code"`
   }
)

@server(
    // 分组 demoA
    group: demoA
)
service demoA {
    // 消费者 GDemoA
    @listener GDemoA
    /queue.test.5 /queue6
}

@server(
    // 分组demoA
    group: demoA
)
service demoA {
    // 消费者 GDemoAZZ
    @listener GDemoAZZ
    /queue3 /queue4
}

@server(
    // 分组demoB
    group: demoB
)
service demoA {
    // 消费者 GDemoBZZ
    @listener GDemoBZZ
    /queue.test.5 /queue6
}
```