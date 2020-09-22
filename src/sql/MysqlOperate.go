package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 一定需要加该行，否则会报错找不到驱动 unknown driver "mysql" (forgotten import?)
	"time"
)

/**
 * <p>
 * go 操作mysql示例参考：https://www.cnblogs.com/rickiyang/p/11074169.html
 * </p>
 * @author: zhu.chen
 * @date: 2020/9/21
 * @version: v1.0.0
 */
const (
	USERNAME = "root"
	PASSWORD = "zhuchen1994"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "cloud_note"
)

type User struct {
	Id       string `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Nick     string `json:"nick" form:"nick"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"token"`
}

// 创建连接 "*"：指针运算符，指针传递，返回变量的值。"&"：取地址符号，地址传递，返回变量的地址。
func createConnect() (DB *sql.DB) {
	// conn格式：root:zhuchen1994@tcp(127.0.0.1:3306)/cloud_note
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	fmt.Println(conn)
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}
	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超时的连接就close
	DB.SetMaxOpenConns(100)
	return DB
}

//插入数据
func InsertData() {
	DB := createConnect()
	result, err := DB.Exec("insert INTO user(name,nick,password,token) values(?,?,?,?)", "zhangsan", "xiaoya", "test", "123456")
	if err != nil {
		fmt.Printf("Insert data failed,err:%v", err)
		return
	}
	lastInsertID, err := result.LastInsertId() //获取插入数据的自增ID
	if err != nil {
		fmt.Printf("Get insert id failed,err:%v", err)
		return
	}
	fmt.Println("Insert data id:", lastInsertID)

	rowsaffected, err := result.RowsAffected() //通过RowsAffected获取受影响的行数
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("Affected rows:", rowsaffected)
}

//查询单行
func QueryOne() {
	DB := createConnect()
	user := new(User) //用new()函数初始化一个结构体对象
	row := DB.QueryRow("select id,name,nick,password,token from user where id=?", 1)
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	if err := row.Scan(&user.Id, &user.Name, &user.Nick, &user.Password, &user.Token); err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Println("Single row data:", *user)
}

//查询多行
func QueryMulti() {
	DB := createConnect()
	user := new(User)
	rows, err := DB.Query("select id,name,nick,password,token from user where id=?", 2)

	defer func() {
		if rows != nil {
			rows.Close() //关闭掉未scan的sql连接
		}
	}()
	if err != nil {
		fmt.Printf("Query failed,err:%v\n", err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Nick, &user.Password, &user.Token) //不scan会导致连接不释放
		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return
		}
		fmt.Println("scan successd:", *user)
	}
}

//更新数据
func UpdateData() {
	DB := createConnect()
	result, err := DB.Exec("UPDATE user set password=? where id=?", "111111", 1)
	if err != nil {
		fmt.Printf("Insert failed,err:%v\n", err)
		return
	}
	fmt.Println("update data successd:", result)

	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v\n", err)
		return
	}
	fmt.Println("Affected rows:", rowsaffected)
}

//删除数据
func DeleteData() {
	DB := createConnect()
	result, err := DB.Exec("delete from user where id=?", 1)
	if err != nil {
		fmt.Printf("Insert failed,err:%v\n", err)
		return
	}
	fmt.Println("delete data successd:", result)

	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v\n", err)
		return
	}
	fmt.Println("Affected rows:", rowsaffected)
}
