package operation

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"rest-demo/entity"
	"time"
)

func delete() {
	db, err := gorm.Open("mysql", "root:123456@/gorm?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	user := entity.User{Name: "HiPhone", Age: 18, Birthday: time.Now()}
	db.Create(&user)

	var result entity.User
	db.First(&result)
	db.Delete(&result)

	//delete from users where name like %Hi%
	db.Where("name like ?", "%Hi%").Delete(entity.User{})
	db.Delete(entity.User{}, "email like ?", "%Hi%")

	/*****使用gorm.Model创建库有delete_at字段, 自动获取软删除能力,即删除的话会更新这个字段而已*****/
}
