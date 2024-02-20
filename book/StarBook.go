package book

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"homework/utils"
	"net/http"
)

type Mybookclaim struct {
	ID int `json:"ID"`
	jwt.StandardClaims
}

func StarBook(context *gin.Context) {

	bookID := context.PostForm("book_id")

	authorization := context.Request.Header.Get("Authorization")
	if authorization == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header 为空"})
		context.Abort() //终止当前请求的处理
		return
	}
	if len(authorization) < 7 || authorization[:6] != "Bearer" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "令牌的格式不正确"})
		context.Abort()
		return
	}
	tokenString := authorization[7:]
	token, err := jwt.ParseWithClaims(tokenString, &Mybookclaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.Key), nil
	})
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "令牌解析出错"})
		return
	}
	//token.Valid用于验证令牌是否有效
	var id int
	if claims, ok := token.Claims.(*Mybookclaim); ok && token.Valid {
		// 从令牌中获取声明
		id = claims.ID
		fmt.Println("ID from token:", id)
	} else {
		fmt.Println("token无效")
	}
	//链接数据库
	utils.DB, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/book")
	defer utils.DB.Close()
	err = utils.DB.Ping() //可理解为时间太长就出错
	if err != nil {
		panic("数据库链接失败")
	} //更改is_star
	//DB.Exec 是 Go 中 database/sql 包中的一个方法，用于执行一个 SQL 语句，
	//但不返回任何行的结果。它通常用于执行诸如插入、更新、删除等不需要返回数据行的操作。
	_, err = utils.DB.Exec("UPDATE tb_book SET is_star = ? WHERE id = ?", 1, bookID)
	if err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"error": "数据更新失败"})
		return
	}
	context.JSON(200, gin.H{
		"info":   "success",
		"status": 10000,
	})
}
