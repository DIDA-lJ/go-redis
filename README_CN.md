# Go-Redis（Godis）

Go-Redis 是一个用 Go 语言实现的 Redis 服务器（Godis）。

关键功能:
- 支持 string, list, hash, set, sorted set, bitmap 数据结构
- 自动过期功能(TTL)
- 发布订阅
- 地理位置
- AOF 持久化及 AOF 重写
- 加载和导出 RDB 文件
- 主从复制 (测试中)
- Multi 命令开启的事务具有`原子性`和`隔离性`. 若在执行过程中遇到错误, godis 会回滚已执行的命令
- 内置集群模式. 集群对客户端是透明的, 您可以像使用单机版 redis 一样使用 godis 集群
  - 使用 raft 算法维护集群元数据(测试中)
  - `MSET`, `MSETNX`, `DEL`, `Rename`, `RenameNX`  命令在集群模式下原子性执行, 允许 key 在集群的不同节点上
  - 在集群模式下支持在同一个 slot 内执行事务
- 并行引擎, 无需担心您的操作会阻塞整个服务器.


# 运行 Go-Redis

使用 windows 命令行启动 Godis 服务器

```bash
 .\main.exe
```

![](https://github.com/DIDA-lJ/go-redis/blob/dev/img/img.png)

godis 默认监听 0.0.0.0:6399，可以使用 redis-cli 或者其它 redis 客户端连接 Godis 服务器。

![image](https://github.com/DIDA-lJ/go-redis/assets/97254796/a6915fbc-8532-4df6-b964-78393e2ccbb5)


godis 首先会从CONFIG环境变量中读取配置文件路径。若环境变量中未设置配置文件路径，则会尝试读取工作目录中的 redis.conf 文件。 若 redis.conf 文件不存在则会使用自带的默认配置。

## 集群模式

godis 支持以集群模式运行，请在 redis.conf 文件中添加下列配置:

```ini
peers localhost:7379,localhost:7389 // 集群中其它节点的地址
self  localhost:6399 // 自身地址
```

可以使用 node1.conf 和 node2.conf 配置文件，在本地启动一个双节点集群,然后由于是windows，所以需要移动到不同的文件夹，将 node1.conf 或者 node 2.conf 修改成 redis.conf 才能启动:

```bash
main.exe
```

集群模式对客户端是透明的，只要连接上集群中任意一个节点就可以访问集群中所有数据：

```bash
redis-cli -p 6399
```

## 支持的命令

请参考 [commands.md](https://github.com/DIDA-lJ/go-redis/blob/dev/commands.md)

## 性能测试

环境:

Go version：1.17

System: macOS Catalina 10.15.7

CPU: 2.6GHz 6-Core Intel Core i7

Memory: 16 GB 2667 MHz DDR4

redis-benchmark 测试结果:

```
PING_INLINE: 87260.03 requests per second
PING_BULK: 89206.06 requests per second
SET: 85034.02 requests per second
GET: 87565.68 requests per second
INCR: 91157.70 requests per second
LPUSH: 90334.23 requests per second
RPUSH: 90334.23 requests per second
LPOP: 90334.23 requests per second
RPOP: 90415.91 requests per second
SADD: 90909.09 requests per second
HSET: 84104.29 requests per second
SPOP: 82918.74 requests per second
LPUSH (needed to benchmark LRANGE): 78247.26 requests per second
LRANGE_100 (first 100 elements): 26406.13 requests per second
LRANGE_300 (first 300 elements): 11307.10 requests per second
LRANGE_500 (first 450 elements): 7968.13 requests per second
LRANGE_600 (first 600 elements): 6092.73 requests per second
MSET (10 keys): 65487.89 requests per second
```

