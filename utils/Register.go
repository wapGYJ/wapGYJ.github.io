package utils

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func Register(c *gin.Context) {

	Username := c.PostForm("username")
	Password := c.PostForm("password")

	//链接数据库
	DB, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/user")
	defer DB.Close()
	err := DB.Ping() //可理解为时间太长就出错
	if err != nil {
		panic("数据库链接失败")
	}
	// 检查用户名是否已存在
	var count int
	err = DB.QueryRow("SELECT COUNT(*) FROM tb_user WHERE username = ?", Username).Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户已存在，请重新输入"})
		return
	}

	_, err = DB.Exec("INSERT INTO tb_user (username,password)values(?,?)", Username, Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"info":   "success",
		"status": 10000,
	})

}
