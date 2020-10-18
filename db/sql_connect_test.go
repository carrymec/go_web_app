package db

import (
	"fmt"
	"testing"
)

func TestInitDB(t *testing.T) {
	err := InitDB()
	if err != nil {
		t.Fatal(err)
		return
	}
	_, err = Db.Exec("create table book(" +
		"id int auto_increment primary key," +
		"name varchar(128) ," +
		"detail varchar(255) );")
	if err != nil {
		t.Fatal(err)
	}
	println("连接成功")
}

func TestQuery(t *testing.T) {
	err := InitDB()
	if err != nil {
		t.Fatal(err)
		return
	}
	rows, err := Db.Query("select * from book")
	if err != nil {
		t.Fatal(err)
	}
scan:
	if rows.Next() {
		book := new(book)
		err := rows.Scan(&book.Id, &book.Name, &book.Detail)
		if err != nil {
			t.Fatal(err)
			return
		}
		fmt.Println(book.Id, book.Name, book.Detail)
		//代码块和for类似
		goto scan
	}

	println("连接成功")
}

type book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
}
