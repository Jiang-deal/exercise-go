package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//定义模型
type student struct {
	gorm.Model
	Name string
	Age  int64
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	// db.AutoMigrate(&student{})

	// stu1 := student{Name: "廖泽敏", Age: 22}
	// db.Create(&stu1)
	// stu2 := student{Name: "国兴", Age: 12}
	// db.Create(&stu2)

	//查询
	//一般查询
	//var result student

	var results []student
	// db.First(&result) //查询的是第一条记录,只有主键是int类型的才可以
	// fmt.Println("First的值：", result)
	// db.Last(&result) //查询最后一条记录
	// fmt.Println("Last的值：", result)
	// db.Take(&result) //查询随机的一条记录
	// fmt.Println("Take的值：", result)
	// db.Find(&results) //查询所有的记录
	// fmt.Println("Find的值：", results)
	// db.Debug().First(&result, 2) //查询指定的一条记录，只要主键是id才可以用
	// fmt.Println("指定id的值：", result)

	//Where查询
	// Get first matched record
	// db.Where("Name=?", "江兴鸿").Debug().First(&result)
	// fmt.Println("First的值：", result)
	// // Get all matched records
	// db.Where("Name=?", "江兴鸿").Debug().First(&results)
	// fmt.Println("First的值：", result)

	// <> 输出了输入的匹配之外的值，相当于一个或操作
	// db.Where("name <> ?", "江兴宇").Find(&results)
	// fmt.Println("First的值：", results)
	// //// SELECT * FROM users WHERE name <> 'jinzhu';

	// IN
	db.Where("name IN (?)", []string{"江鸿", "江兴"}).Find(&results)
	fmt.Println("First的值：", results)
	//// SELECT * FROM users WHERE name in ('jinzhu','jinzhu 2');

	// LIKE 只要有兴字都能查到
	// db.Where("name LIKE ?", "%兴%").Find(&results)
	// fmt.Println("First的值：", results)
	//// SELECT * FROM users WHERE name LIKE '%jin%';

	// // AND
	// db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	// //// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

	// // Time
	// db.Where("updated_at > ?", lastWeek).Find(&users)
	// //// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

	// // BETWEEN
	// db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
	// //// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

}
