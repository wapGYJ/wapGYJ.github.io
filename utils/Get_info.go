package utils

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Userinfo struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Telephone string `json:"telephone"`
	Gender    string `json:"gender"`
	Qq        string `json:"qq"`
}

func Get_info(context *gin.Context) {
	id := context.Param("id")
	DB, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/user")
	defer DB.Close()
	err := DB.Ping() //可理解为时间太长就出错
	if err != nil {
		fmt.Println("数据库连接失败")
	}
	// 执行查寻
	var userinfo Userinfo
	err = DB.QueryRow("select id,username,nickname,telephone,gender,qq from tb_user where id=? ",
		id).Scan(&userinfo.Id, &userinfo.Username, &userinfo.Nickname, &userinfo.Telephone, &userinfo.Gender, &userinfo.Qq)
	if err != nil {
		log.Fatal(err)
	}
	context.JSON(200, gin.H{
		"status": 10000,
		"info":   "success",
		"data":   userinfo,
	})
}
