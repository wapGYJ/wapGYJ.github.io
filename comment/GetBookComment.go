package comment

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"homework/utils"
	"log"
	"net/http"
)

type Comment struct {
	PostId      int    `json:"post_id"`
	PublishTime int64  `json:"publish_time"`
	Content     string `json:"content"`
	UserId      string `json:"user_id"`
	Avatar      string `json:"avatar"`
	Nickname    string `json:"nickname"`
	PraiseCount int    `json:"praise_count"`
	IsPraised   bool   `json:"is_praised"`
	IsFocus     bool   `json:"is_focus"`
	BookId      string `json:"book_id"`
}

func GetBookComment(context *gin.Context) {
	BookID := context.Param("book_id")
	utils.DB, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/comment")
	defer utils.DB.Close()
	err := utils.DB.Ping()
	if err != nil {
		panic("数据库链接失败")
	}

	rows, err := utils.DB.Query("SELECT post_id,publish_time,content,user_id,avatar,nickname,praise_count,is_praised,is_focus,book_id from tb_comment where book_id=?", BookID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var comments []Comment

	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.PostId, &comment.PublishTime, &comment.Content, &comment.UserId,
			&comment.Avatar, &comment.Nickname, &comment.PraiseCount, &comment.IsPraised,
			&comment.IsFocus, &comment.BookId); err != nil {
			log.Fatal(err)
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	if err := rows.Err(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回JSON响应
	context.JSON(http.StatusOK, comments)

}
