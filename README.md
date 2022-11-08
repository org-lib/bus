<p align="center">
        <img src="img/logo.jpg">
</p>

# bus
极简

让开发 golang 更轻松快速、不需要到处寻找资源库

    代码目录结构：
    aliyun          oss-options
    assets          #
    aws             oss-options
    config          #封装的 viper 包
    cors            #跨域请求
    db              #数据库连接基础驱动包
    disk            磁盘空间-options
    example         #数据库、makeapp、api、并发控制 etc. 的使用例子
    img             #
    logger          #封装的 zap 包
    msi             #Windows msi 编译包 ， 暂时未开放，需要使用请联系项目开发者shangguannihao@gmail.com
    nanoid          生成唯一ID
    notify          系统弹窗提示
    parser          数据库SQL 语法解析
    pool            #
    xshell          #命令行窗口
    xxl-job         定时任务
    README.md       #说明书
目前分为如下几个部分：

    第一部分 标准化日志格式输出
            1)自定义配置文件路径(日常需要指定自己的配置文件路径
            2)日志统一（不需要使用API 让别人调用，只是普通的命令行程序打印日志
            3)api 日志统一（包括超时处理，跨域允许;使用API 的日志打印
    第二部分 数据库CRUD
            1)MySQL
              便捷DML、查询命令结果SHOW STATUS 、show slave status 、etc.
            2)Clickhouse
            3)努力更新中...
    第三部分 努力更新中...

# 第一部分 标准化日志格式输出
    #配置文件内容
    cat conf.yaml

    server:
      listen_ip: "0.0.0.0"
      listen_port: "80"
    log:
      path: "./agg.log"
      max_size: 100
      max_backup: 30
      max_age: 30
      level: "debug"

1)自定义配置文件路径(日常需要指定自己的配置文件路径，则引用格式如下)：

    # yaml 格式
    xxx.exe -conf=/xxx/yyy/zzz.yaml

    # json 格式
    xxx.exe -conf=/xxx/yyy/zzz.json

    # toml 格式
    xxx.exe -conf=/xxx/yyy/zzz.toml

    # ini 格式
    xxx.exe -conf=/xxx/yyy/zzz.ini

    # hcl 格式
    xxx.exe -conf=/xxx/yyy/zzz.hcl

2)日志统一（不需要使用API 让别人调用，只是普通的命令行程序，则引用格式如下：）

    #日志打印方式一（首先得有 yaml 配置文件）

    请参考 example/purelog

api 日志统一（包括超时处理，跨域允许;使用API，则引用格式如下：）
    
    请参考 example/api

# 第二部分 数据库CRUD
1)MySQL
请参考 example/mysql

2)Clickhouse
请参考 example/clickhouse
