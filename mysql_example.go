package main

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	//"github.com/go-redis/redis"
	"github.com/go-redis/redis/v8"
	_ "github.com/goinaction/code/chapter5/listing71/entities"
)

type Student struct{
	Id int
	FirstName string
	Age int
}

var ctx = context.Background()

func test_redis(){
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "get.item.1.2", "[]", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
func test_mysql(){
	dns := "root:taoquan1234@tcp(127.0.0.1:3306)/lam1?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Print("Hello World for db connector")
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN: dns,
	}), &gorm.Config{})
	if err != nil{
		fmt.Print("error happen when cannot connect to db: ", err)
		return
	}
	user := Student{FirstName: "Pham Ngoc La11111", Age:  25}

	result := db.Debug().Create(&user)
	fmt.Print("transaction error : ", result.Error, " and new thing: ", result.RowsAffected)
	var user1 Student
	db.Where("Id = 32").First(&user1)
	fmt.Print(user1)
}
func main(){
	//test_mysql()
	//test_redis()
	PrintKakakaka()
}