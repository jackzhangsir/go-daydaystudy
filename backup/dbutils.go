package backup


import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init()  {
	db, err := sqlx.Open("mysql", "user:passwd@tcp(127.0.0.1:3307)/base?charset=utf8")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = db
}

func QueryDoctor() int {

	fmt.Println("start ...")

	rows, err := Db.Query("SELECT f_name,f_title FROM t_content")
	if err != nil{
		fmt.Println("query failed,error： ", err)
		return 0
	}
	for rows.Next() {  //循环结果
		var f_name,f_title string
		err = rows.Scan(&f_name, &f_title)
		println(f_name, f_title)
	}


	fmt.Println("end...")
	return 0
}
