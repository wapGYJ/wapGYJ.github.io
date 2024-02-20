package book

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"homework/utils"
	"log"
	"net/http"
)

func Label(context *gin.Context) {
	label := context.Query("label")
	utils.DB, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/book")
	defer utils.DB.Close()
	err := utils.DB.Ping() //可理解为时间太长就出错
	if err != nil {
		panic("数据库链接失败")
	}
	var book Bookinfo
	err = utils.DB.QueryRow("select id,name,is_star,comment_num,score,cover,publish_time,link,author,lable from tb_book where lable=?", label).Scan(&book.Id,
		&book.Name, &book.Is_star, &book.Comment_num, &book.Score, &book.Cover,
		&book.Publish_time, &book.Link, &book.Author, &book.Label)
	if err != nil {
		log.Fatal(err)
	}
	context.JSON(http.StatusOK, gin.H{
		"status": 10000,
		"info":   "success",
		"data":   book,
	})
}
