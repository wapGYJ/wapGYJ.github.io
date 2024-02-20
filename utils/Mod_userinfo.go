package utils

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Mod_userinfo(context *gin.Context) {
	new_nickname := context.Query("new_nickname")
	authorization := context.Request.Header.Get("Authorization")
	if authorization == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header 为空"})
		context.Abort()
		return
	}
	if len(authorization) < 7 || authorization[:6] != "Bearer" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "令牌的格式不正确"})
		context.Abort()
		return
	}
	tokenString := authorization[7:]

	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Key), nil
	})
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "令牌解析出错"})
		return
	}
	//token.Valid用于验证令牌是否有效
	var username string
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		// 从令牌中获取声明
		username = claims.Username
		fmt.Println("Username from token:", username)
	} else {
		fmt.Println("token无效")
	}
	//链接数据库
	DB, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/user")
	defer DB.Close()
	err = DB.Ping() //可理解为时间太长就出错
	if err != nil {
		panic("数据库链接失败")
	} //更改绰号
	//DB.Exec 是 Go 中 database/sql 包中的一个方法，用于执行一个 SQL 语句，
	//但不返回任何行的结果。它通常用于执行诸如插入、更新、删除等不需要返回数据行的操作。
	_, err = DB.Exec("UPDATE tb_user SET nickname = ? WHERE username = ?", new_nickname, username)
	if err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"error": "绰号更新失败"})
		return
	}
	context.JSON(200, gin.H{
		"info":   "success",
		"status": 10000,
	})
}
