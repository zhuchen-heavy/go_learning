package main

import "go_learning/web"

func main() {
	// 1：go的可见性
	//a.Api()

	// 2：md5工具类
	//value := crypt.Md5("zhangsan")
	//fmt.Println("value is : " + value)

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
	//file.ReadYmlConfig()

	web.StartTest()
}
