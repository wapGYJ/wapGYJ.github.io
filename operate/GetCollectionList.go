package operate

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"homework/utils"
	"log"
	"net/http"
)

type Collection struct {
	BookId      string `json:"book_Id"`
	Name        string `json:"name"`
	PublishTime string `json:"publishTime"`
	Link        string `json:"link"`
	Username    string `json:"username"`
}

func GetCollectList(context *gin.Context) {
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
	var username string
	if claims, ok := token.Claims.(*myclaims); ok && token.Valid {
		// 从令牌中获取声明
		username = claims.Username
		fmt.Printf("用户名为%s/n", username)
	} else {
		fmt.Println("token无效")
	} //链接数据库
	utils.DB, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/collections")
	defer utils.DB.Close()
	err = utils.DB.Ping() //可理解为时间太长就出错
	if err != nil {
		panic("数据库链接失败")
	}
	rows, err := utils.DB.Query("SELECT book_id,name,publish_time,link,username from tb_collections where username=?", username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var collections []Collection

	for rows.Next() {
		var collection Collection
		if err := rows.Scan(&collection.BookId, &collection.Name, &collection.PublishTime, &collection.Link,
			&collection.Username); err != nil {
			log.Fatal(err)
		}
		collections = append(collections, collection)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	if err := rows.Err(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 返回JSON响应

	context.JSON(http.StatusOK, gin.H{
		"status": 10000,
		"info":   "success",
		"data":   collections,
	})

}
