// 针对的是模块名
module go_learning

go 1.14

// 引入第三方包的时候，需要调用 go get 命令。e.g： go get github.com/go-redis/redis

// Redis的依赖包
require github.com/go-redis/redis v6.15.9+incompatible
