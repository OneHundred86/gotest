package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Test() {
	db, err := sqlx.Connect("mysql", "root:root123@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return
	}

	defer db.Close()

	/*
	   CREATE TABLE `site` (
	     `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
	     `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
	     `domain` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
	     `proto` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT 'http',
	     PRIMARY KEY (`id`)
	   ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
	*/

	// Get : select
	type maxid struct {
		ID sql.NullInt32 `db:"max_id"`
	}
	mxid := maxid{}
	err = db.Get(&mxid, "select max(id) max_id from site")
	if err != nil {
		fmt.Println("获取maxid失败", err)
		return
	}
	fmt.Println("max id:", mxid.ID)

	// Exec : insert/update/delete
	nextid := mxid.ID.Int32 + 1
	name := fmt.Sprintf("site%d", nextid)
	domain := fmt.Sprintf("test%d", nextid)
	result, err1 := db.Exec("insert into site (name, domain, proto) value(?,?,?)", name, domain, "http")
	if err1 != nil {
		fmt.Println("insert数据库失败", err1)
		return
	}
	id, _ := result.LastInsertId()
	fmt.Println("insert数据成功")

	//
	result, err = db.Exec("update site set proto = ? where id = ?", "https", id)
	if err != nil {
		fmt.Println("update失败", err)
	} else {
		if n, err2 := result.RowsAffected(); err2 == nil && n > 0 {
			fmt.Println("update 成功")
		}
	}

	type site struct {
		Id       int    `db:"id"`
		SiteName string `db:"name"`
		Domain   string
		Proto    string
	}

	// Select : select
	sts := []site{}
	err = db.Select(&sts, "select * from site where name like ?", "%site%")
	fmt.Println(sts)
}
