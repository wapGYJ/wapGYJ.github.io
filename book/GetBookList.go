package book

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"homework/utils"
	"log"
	"net/http"
)

type Bookinfo struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Is_star      bool    `json:"is_Star"`
	Comment_num  int     `json:"comment_Num"`
	Score        float32 `json:"score"`
	Cover        string  `json:"cover"`
	Publish_time string  `json:"publish_Time"`
	Link         string  `json:"link"`
	Author       string  `json:"author"`
	Label        string  `json:"label"`
}

func GetBookList(context *gin.Context) {
	utils.DB, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/book")
	defer utils.DB.Close()
	err := utils.DB.Ping()
	if err != nil {
		panic("数据库链接失败")
	}

	rows, err := utils.DB.Query("SELECT * FROM tb_book")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var books []Bookinfo

	for rows.Next() {
		var book Bookinfo
		if err := rows.Scan(&book.Id, &book.Name, &book.Is_star, &book.Comment_num,
			&book.Score, &book.Cover, &book.Publish_time, &book.Link, &book.Author, &book.Label); err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	if err := rows.Err(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回JSON响应
	context.JSON(http.StatusOK, books)
}
