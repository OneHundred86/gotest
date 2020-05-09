package mysql

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type GoUser struct {
	gorm.Model
	Name  string
	Age   int
	Email string `gorm:"type:varchar(100);unique_index;not null"`
	Phone string `gorm:"size:32"` // 设置字段大小为32
}

// `Profile` 属于 `User`， 外键是`UserID`
type GoProfile struct {
	ID       int `gorm:"PRIMARY_KEY"`
	GoUserID int `gorm:"INDEX:user_id;"` // 外键
	GoUser
	Key string
	Val string
}

func InitGorm() {
	var err error
	db, err = gorm.Open("mysql", "root:root123@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败:" + err.Error())
	}
}

func Migrate() {
	db.AutoMigrate(&GoUser{}, &GoProfile{})
}

func TestGorm() {
	InitGorm()
	defer db.Close()
	Migrate()

	var user GoUser

	// 根据主键查询最后一条记录
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	db.Last(&user)
	fmt.Printf("last: %v \n", user)

	// 随机种子
	rand.Seed(time.Now().UnixNano())

	// insert
	lastid := user.ID
	user = GoUser{
		Name:  fmt.Sprintf("name%d", lastid+1),
		Age:   rand.Intn(100),
		Email: fmt.Sprintf("test%d@t.cn", lastid+1),
		Phone: "10086",
	}
	db.Create(&user)
	fmt.Printf("create: %+v \n", user)

	// 获取第一个匹配的记录
	// SELECT * FROM users ORDER BY id LIMIT 1;
	var user1 GoUser
	db.Where("id = ?", "1").First(&user1)
	fmt.Printf("first: %v \n", user1)

	// SELECT * FROM users;
	var users []GoUser
	db.Find(&users)
	fmt.Println(users)

	// 事物
	err := db.Transaction(func(tx *gorm.DB) (err error) {
		user := GoUser{}
		user.ID = 1
		if err = tx.Create(&user).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Println("执行事物失败:", err)
	}

	TestRelation()
}

func TestRelation() {
	TestBelongsTo()
}

// 反向一对一
func TestBelongsTo() {

	var profile GoProfile
	db.Debug().Where("id = ?", 1).First(&profile)

	var user GoUser
	db.Debug().Model(&profile).Related(&user)
	fmt.Printf("belongs to: %+v, %+v", profile, user)
}
