package nosql

/**
 * <p>
 * redis连接类
 * </p>
 * @author: zhu.chen
 * @date: 2020/9/19
 * @version: v1.0.0
 */

// 引入第三方包
import (
	"fmt"
	"github.com/go-redis/redis"
)

func RedisOperate() {
	addr := "127.0.0.1:6379"
	password := "zhuchen1994"
	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       15,
	})

	p, err := c.Ping().Result()

	if err != nil {
		fmt.Println("redis kill")
	}

	fmt.Println(p)

	c.Do("SET", "key", "zhuchen")

	rs := c.Do("GET", "key").Val()
	fmt.Println(rs)
	c.Close()
}
