# Supported Commands

- Keys
    - del
    - expire
    - expireat
    - pexpire
    - pexpireat
    - ttl
    - pttl
    - persist
    - exists
    - type
    - rename
    - renamenx
- Server
    - flushdb
    - flushall
    - keys
    - bgrewriteaof
    - copy
- String
    - set
    - setnx
    - setex
    - psetex
    - mset
    - mget
    - msetnx
    - get
    - getex
    - getset
    - getdel
    - incr
    - incrby
    - incrbyfloat
    - decr
    - decrby
    - randomkey
- List
    - lpush
    - lpushx
    - rpush
    - rpushx
    - lpop
    - rpop
    - rpoplpush
    - lrem
    - llen
    - lindex
    - lset
    - lrange
    - ltrim
    - linsert
- Hash
    - hset
    - hsetnx
    - hget
    - hexists
    - hdel
    - hlen
    - hstrlen
    - hmget
    - hmset
    - hkeys
    - hvals
    - hgetall
    - hincrby
    - hincrbyfloat
    - hrandfield
- Set
    - sadd
    - sismember
    - srem
    - spop
    - scard
    - smembers
    - sinter
    - sinterstore
    - sunion
    - sunionstore
    - sdiff
    - sdiffstore
    - srandmember
- SortedSet
    - zadd
    - zscore
    - zincrby
    - zrank
    - zcount
    - zrevrank
    - zcard
    - zrange
    - zrevrange
    - zrangebyscore
    - zrevrangebyscore
    - zrem
    - zremrangebyscore
    - zremrangebyrank
    - zlexcount
    - zrangebylex
    - zremrangebylex
    - zrevrangebylex
- Pub / Sub
    - publish
    - subscribe
    - unsubscribe
- Geo
    - GeoAdd
    - GeoPos
    - GeoDist
    - GeoHash
    - GeoRadius
    - GeoRadiusByMember