package backup

import (
	//导入MYSQL数据库驱动，这里使用的是GORM库封装的MYSQL驱动导入包，实际上大家看源码就知道，这里等价于导入github.com/go-sql-driver/mysql
	//这里导入包使用了 _ 前缀代表仅仅是导入这个包，但是我们在代码里面不会直接使用。
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func init()  {
	//
}

type CityArea struct {
	Id int
	AreaCode int
	AreaName string
	CityCode int
	CreateTime string
	UpdateTime string
}

//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (u *CityArea) TableName() string {
	//绑定MYSQL表名为users
	return "t_city_area"
}

func main()  {

	//配置MySQL连接参数
	username := ""  //账号
	password := "" //密码
	host := "" //数据库地址，可以是Ip或者域名
	port := 3306 //数据库端口
	Dbname := "" //数据库名

	//通过前面的数据库参数，拼接MYSQL DSN， 其实就是数据库连接串（数据源名称）
	//MYSQL dsn格式： {username}:{password}@tcp({host}:{port})/{Dbname}?charset=utf8&parseTime=True&loc=Local
	//类似{username}使用花括号包着的名字都是需要替换的参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	//连接MYSQL
	_db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	//设置数据库连接池参数
	_db.DB().SetMaxOpenConns(10)   //设置数据库连接池最大连接数
	_db.DB().SetMaxIdleConns(5)   //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。

	//定义一个用户，并初始化数据
	u := CityArea{
		Id:         0,
		AreaCode:   0,
		AreaName:   "",
		CityCode:   0,
		CreateTime: "",
		UpdateTime: "",
	}

	fmt.Println("当前时间= ", time.Now().String())

	//插入一条用户数据
	//下面代码会自动生成SQL语句：INSERT INTO `users` (`username`,`password`,`createtime`) VALUES ('tizi365','123456','1540824823')
	if err := _db.Create(u).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}

	//查询并返回第一条数据
	//定义需要保存数据的struct变量
	u = CityArea{}
	//自动生成sql： SELECT * FROM `users`  WHERE (username = 'tizi365') LIMIT 1
	isNotFound := _db.Where("f_city_code = ?", "1000").First(&u).RecordNotFound()
	if isNotFound {
		fmt.Println("找不到记录")
		return
	}
	//打印查询到的数据


	//更新
	//自动生成Sql: UPDATE `users` SET `password` = '654321'  WHERE (username = 'tizi365')
	_db.Model(CityArea{}).Where("f_city_code = ?", "1000").Update("f_area_code", "1000")

	//删除
	//自动生成Sql： DELETE FROM `users`  WHERE (username = 'tizi365')
	//db.Where("username = ?", "tizi365").Delete(CityArea{})
}







