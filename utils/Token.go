package utils

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var Key = "gyjgyjgyj" //密钥

var DB *sql.DB

func Gettoken(context *gin.Context) {
	Username := context.Query("username")
	Password := context.Query("password")
	DB, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/user")
	defer DB.Close()
	err := DB.Ping()
	if err != nil {
		panic("数据库链接失败")
	}
	// 执行查寻
	var username string
	var password string
	err = DB.QueryRow("SELECT username,password FROM tb_user WHERE username = ?", Username).Scan(&username, &password)
	if err != nil {
		log.Fatal(err)
	}
	if username == Username && Password == password {
		//设置token令牌的声明
		claims := jwt.MapClaims{
			"username": username,
			"password": password,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //签名
		tokenString, err := token.SignedString([]byte(Key))
		if err != nil {
			context.JSON(400, gin.H{"err": "token生成失败"})
			panic(err)
		} //设置请求头
		context.Request.Header.Set("Authorization", "Bearer "+tokenString)
		fmt.Println(context.Request.Header.Get("Authorization")) //调用Header.Get()方法来获取请求头中Authorization字段的值
		data := gin.H{
			"token": tokenString,
		}
		context.JSON(200, gin.H{
			"status": 10000,
			"info":   "success",
			"data":   data,
		})
	} else {
		context.JSON(400, gin.H{"error": "密码或用户名错误请重新输入"})
		panic("密码或用户名错误请重新输入")
	}

}
