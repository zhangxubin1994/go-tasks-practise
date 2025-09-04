package main

import (
	"go-tasks-practise/gorm/task02"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "time/tzdata" // 导入时区数据包
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/test_gorm?charset=utf8&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	task02.Add(db)
}
