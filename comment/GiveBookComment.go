package comment

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"homework/utils"
	"net/http"
)

type myclaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GiveBookComment(context *gin.Context) {
	BookId := context.Param("book_id")

	formcomment := context.PostForm("comment")
	formPostid := context.PostForm("postid")
	formPublishtime := context.PostForm("publishtime")
	formAvter := context.PostForm("avter")
	formnick := context.PostForm("nickname")
	fromPraisecount := context.PostForm("prisecount")
	fromIspraise := context.PostForm("ispraise")
	fromIsfocus := context.PostForm("isfocus")

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

	token, err := jwt.ParseWithClaims(tokenString, &myclaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.Key), nil
	})
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "令牌解析出错"})
		return
	}
	//token.Valid用于验证令牌是否有效
	var id string
	if claims, ok := token.Claims.(*myclaims); ok && token.Valid {
		// 从令牌中获取声明
		id = claims.Id
	} else {
		fmt.Println("token无效")
	}

	//链接数据库
	utils.DB, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/comment")
	defer utils.DB.Close()
	err = utils.DB.Ping() //可理解为时间太长就出错
	if err != nil {
		panic("数据库链接失败")
	}
	//赋值
	_, err = utils.DB.Exec("INSERT INTO tb_comment (book_id, content,user_id,post_id,publish_time,avatar,nickname,praise_count,is_praised,is_focus) VALUES (?,?,?,?,?,?,?,?,?,?)",
		BookId, formcomment, id, formPostid, formPublishtime, formAvter, formnick, fromPraisecount, fromIspraise, fromIsfocus)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"info":   "success",
		"status": 10000,
		"data":   formPostid,
	})

}
