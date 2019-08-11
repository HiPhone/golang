package operation

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"rest-demo/entity"
	"time"
)

func updateOperation() {
	db, err := gorm.Open("mysql", "root:123456@/gorm?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	user := entity.User{Name: "HiPhone", Age: 18, Birthday: time.Now()}
	db.Create(&user)

	//先取出数据库中的user
	var result entity.User
	db.First(&result)

	//set name = hello where active = true
	db.Model(&result).Where("active = ?", true).Update("name", "hello")

	//使用map更新多个属性，指挥更新这些更改的字段
	db.Model(&result).Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})

	//struct更新多个字段,指挥更新这些更改的和非空白字段 **false为空白字段**
	//select and omit选择更新的字段,下面只更新name字段
	db.Model(&result).Select("name").Update(entity.User{Name: "hello", Age: 18})

	//batch updates
	db.Table("users").Where("id IN (?)", []int{10, 11}).Updates(map[string]interface{}{"name": "hello", "age": 18})

	//使用RowsAffected获取更新记录计数
	affected := db.Model(entity.User{}).Updates(entity.User{Name: "hello", Age: 18}).RowsAffected
	fmt.Println(affected)



}