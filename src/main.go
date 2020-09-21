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
	"go_learning/file"
)

func main() {
	// 1：go的可见性
	a.Api()

	// 2：md5工具类
	value := crypt.Md5("zhangsan")
	fmt.Println("value is : " + value)

	// 3：restful api controller
	//controller.HelloController()

	// 4：操作Redis
	//nosql.RedisOperate()

	// 5：操作mysql
	//sql.InsertData()
	//sql.QueryOne()
	//sql.QueryMulti()
	//sql.UpdateData()
	//sql.DeleteData()

	// 6：读取yaml文件
	file.ReadYmlConfig()
}
