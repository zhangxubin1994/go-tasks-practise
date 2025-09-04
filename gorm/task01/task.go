package task01

import (
	"fmt"
	"gorm.io/gorm"
)

/**
  sql语句练习 - 基本的crud操作
*/

type student struct {
	gorm.Model
	Name  string
	Age   int
	Grade string
}

func Add(db *gorm.DB) {
	//自动创建数据库表结构  等
	db.AutoMigrate(student{})
	//student1 := student{Name: "李四", Age: 12, Grade: "三年级"}
	//1--插入一条数据
	//db.Create(&student1)
	//2--查询大于18岁的学生信息
	student2 := student{}
	db.Debug().First(&student2, "age > ?", "18")
	fmt.Println(student2)
	//3--更新学生的年级 张三
	//student3 := student{Name: "张三"}
	//db.Model(&student{}).Where("name = ?", "张三").Update("grade", "四年级")
	//4--删除小于十五岁的学生信息
	//硬删除  （unscoped）
	db.Debug().Unscoped().Where("age < ?", "15").Delete(&student{})

}
