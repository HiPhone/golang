package operation

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"rest-demo/entity"
	"time"
)

func createOperation() {
	db, err := gorm.Open("mysql", "root:123456@/gorm?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	user := entity.User{Name: "HiPhone", Age: 18, Birthday: time.Now()}
	db.Create(&user)
	/** []struct可以重复使用, 但是struct只能使用一次, 用于接收数据库的信息 **/
	var resultOne entity.User
	var resultAll []entity.User
	//select * from users order  by id limit 1;
	db.First(&resultOne)

	//select * from users order by id DESC limit 1
	db.Last(&resultOne)

	//select * from users
	db.Find(&resultAll)

	//select * from users where id = 10
	db.First(&resultOne, 10)

	//select * from users where name = 'HiPhone' limit 1
	db.Where("name = ?", "HiPhone").First(&resultOne)

	//select * from users where name = 'HiPhone'
	db.Where("name = ?", "HiPhone").Find(&resultAll)

	//!=
	db.Where("name <> ?", "HiPhone").Find(&resultAll)

	//in
	db.Where("name in (?)", []string{"HiPhone", "HiPhone2"}).Find(&resultAll)

	//Like
	db.Where("name like ?", "%Hi%").Find(&resultAll)

	//and
	db.Where("name = ? and age >= ?", "HiPhone", "22").Find(&resultAll)

	//Time
	db.Where("update_at < ?", time.Now()).Find(&resultAll)

	//between
	db.Where("age between ? and ?", 1, 23)

	//struct
	//select * from users where name = "HiPhone" and age = 20 limit 1
	db.Where(&entity.User{Name: "HiPhone", Age: 20}).First(&resultOne)

	//map
	//select * from users where name = "HiPhone" and age = 20
	db.Where(map[string]interface{} {"name": "HiPhone", "age": 20}).Find(&resultAll)

	//主键slice
	//select * from users where id in (20, 21, 22)
	db.Where([]int64{20, 21, 22}).Find(&resultAll)

	//select * from users where name <> "HiPhone" limit 1
	db.Not("name", "HiPhone").First(&resultOne)

	//not in
	//select * from users where name not in ("HiPhone", "HiPhone2")
	db.Not("name", []string{"HiPhone", "HiPhone2"}).Find(&resultAll)

	//not in slice of primary keys
	//select * from users where id not in (1,2,3)
	db.Not([]int64{1,2,3}).First(&resultAll)

	//select * from users
	db.Not([]int64{}).First(&resultOne)

	// plain sql
	//select * from users where not(name = "HiPhone")
	db.Not("name = ?", "HiPhone").First(&resultOne)

	//struct
	//select * from users where name <> "HiPhone"
	db.Not(entity.User{Name: "HiPhone"}).First(&resultOne)

	//内联条件的查询
	// select * from users where name = "HiPhone"
	db.Find(&resultOne, "name = ?", "HiPhone")

	//select * from users where name <> "HiPhone" and age > 20
	db.Find(&resultAll, "name <> ? and age > ?", "HiPhone", 20)

	//struct
	//select * from users where age = 20
	db.Find(&resultAll, entity.User{Age: 20})

	//map
	//select * from users where age = 20
	db.Find(&resultAll, map[string]interface{}{"age": 20})

	//or条件查询
	//select * from users where role = 'admin' or role = 'super_admin'
	db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&resultAll)

	//struct
	//select * from users where name = "HiPhone" or name = "HiPhone2"
	db.Where("name = 'HiPhone'").Or(entity.User{Name: "HiPhone2"}).Find(&resultAll)

	//map
	db.Where("name = 'HiPhone").Or(map[string]interface{}{"name": "HiPhone2"}).Find(&resultAll)

	//link
	//select * from users where name <> 'HiPhone' and age >= 20 and role <> 'admin'
	db.Where("name <> ?", "HiPhone").Where("age >= ? and role <> ?", 20, "admin").Find(&resultAll)
	db.Where("role = ?", "admin").Or("role = ?", "super_admin").Not("name = ?", "HiPhone").Find(&resultAll)

	//firstOrInit: 获取第一个匹配的记录，或者使用给定的条件初始化一个新的记录（仅适用于struct，map条件）
	db.Where(entity.User{Name: "HiPhone"}).FirstOrInit(&resultOne)
	db.FirstOrInit(&resultOne, map[string]interface{}{"name": "HiPone"})

	//Attrs
	//如果未找到记录,则使用参数初始化结构
	db.Where(entity.User{Name: "HiPhone"}).Attrs(entity.User{Age: 20}).FirstOrInit(&resultOne)

	//Assign
	//将参数分配给结,不管是否被找到
	db.Where(entity.User{Name: "HiPhone"}).Assign(entity.User{Age: 20}).FirstOrInit(&resultOne)

	//firstOrCreate: 获取第一个匹配的记录，或者创建一个具有给定调价你的新纪录
	db.FirstOrCreate(&resultOne, entity.User{Name: "non_existing"})

	//select: 指定要从数据库检索的字段,默认情况下,将选择所有字段
	db.Select([]string{"name", "age"}).Find(&resultOne)
	db.Table("users").Select("COALESCE(age, ?)", 42).Rows()


	//select * from users order by age desc, name limit 10 offset 10
	db.Limit(10).Offset(10).Order("age desc").Order("name").Find(&resultAll)

	var count int64
	//count: select count(*) from delete_users
	db.Table("deleted_users").Count(&count)

	//group & having
	rows, err := db.Table("orders").Select("date(create_at) as date, sum(amount) as total").Group("date(create_at)").Having("sum(amout) > ?", 100).Rows()
	for rows.Next() {

	}

	//join
	db.Joins("JOIN emails ON emails.user_id = user.id AND emails.email = ?", "zhyzyhf@gmail.com").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "4111111").Find(&resultOne)

	//scan  将结果扫描到另一个结构中
	type Result struct {
		Name string
		Age int
	}

	var result Result
	db.Table("users").Select("name, age").Where("name = ?", 3).Scan(&result)
	//raw sql
	db.Raw("Select name, age from users where name = ?", 3).Scan(&result)




}

