package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//1,定义模型
type User struct {
	ID   int64
	Name *string `gorm:"default:'小王子'"`
	Age  int64
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//2.把模型和数据库中的表对应出来，自动迁移
	db.AutoMigrate(&User{}) //这里会自动创建一个表结构
	//3.创建记录
	// u := User{Name: "小江", Age: 22}
	// u2 := User{Age: 18}

	//一个问题：如果结构体中使用tag指定了默认值，然后呢你的结构体初始化时指定值为空那么会发现加载到数据库的时候还是会使用默认值？？？？
	/*理想形式：如果在初试化时没有初始化这个字段，那就使用默认值，有的话就用空值
	解决方法：指针
	*/
	t := new(string)
	*t = "nihao"
	fmt.Println("chon", *t)

	data := User{Name: new(string), Age: 85} //初始化一个空的指针,如果想赋值，直接把t赋值进去

	//fmt.Println(db.NewRecord(&u)) //判断主键是否为空，如果返回一个false就说明数据库中存在这样一个主键
	//db.Create(&u)
	result := db.Debug().Create(&data) //创建一个记录,加了Debug就会输出sql语句
	fmt.Println(result.RowsAffected)   //返回插入记录的条数
	fmt.Println(result.Error)          //返回插入记录的条数
	fmt.Println("chon", *t)

}
