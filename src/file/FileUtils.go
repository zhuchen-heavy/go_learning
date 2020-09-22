package file

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

/**
 * <p>
 * (m *mysql)格式参考：https://juejin.im/post/6844903817952100365
 * </p>
 * @author: zhu.chen
 * @date: 2020/9/21
 * @version: v1.0.0
 */
type mysql struct {
	Host   string `yaml: "host"`
	User   string `yaml: "user"`
	Pwd    string `yaml: "pwd"`
	Dbname string `yaml: "dbname"`
}

// 读取yaml文件转换为mysql对象。
// (m *mysql) ： 变量的方法，面向对象的写法。
// Go接收器：方法上增加类型的描述。一个类型加上它的方法等价于面向对象Java中的一个类
func (m *mysql) getConf() *mysql {
	yamlFile, err := ioutil.ReadFile("./mysql.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, m)
	if err != nil {
		fmt.Println(err.Error())
	}
	return m
}

func ReadYmlConfig() {
	var m mysql
	mysql := m.getConf()
	// "*"：直接获取值 {localhost:3306 root zhuchen1994 cloud_note}
	fmt.Println(*mysql)
	// "&"：获取变量的地址 0xc0000ca030
	fmt.Println(&mysql)
}
