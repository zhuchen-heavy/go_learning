package main

/**
 * <p>
 * 使用自定义包：https://blog.csdn.net/hjxisking/article/details/105223099
 * e.g : go.mod 里面的 module 名下面的文件夹名
 * </p>
 * @author: zhu.chen
 * @date: 2020/9/18
 * @version: v1.0.0
 */
import (
	"fmt"
	"go_learning/a"
	"go_learning/crypt"
	"go_learning/nosql"
)

func main() {
	a.Api()
	value := crypt.Md5("zhangsan")
	fmt.Println("value is : " + value)
	//controller.HelloController()
	nosql.RedisOperate()
}
