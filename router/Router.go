package main

import (
	"github.com/gin-gonic/gin"
	"homework/book"
	"homework/comment"
	"homework/utils"
)

func main() {

	router := gin.Default()
	//用户注册
	router.POST("/register", utils.Register)

	UserRouter := router.Group("/user")
	{
		//用户登录,获取token
		UserRouter.GET("/token", utils.Gettoken)

		//用户修改密码
		UserRouter.PUT("/password", utils.RePassword)

		//获取用户的信息
		UserRouter.GET("/info/:id", utils.Get_info)

		//更改用户信息
		UserRouter.PUT("/info", utils.Mod_userinfo)
	}
	BookRouter := router.Group("/book")
	{
		//获取书籍列表
		BookRouter.GET("/list", book.GetBookList)

		//搜索书籍
		BookRouter.GET("/search", book.SearchBook)

		//收藏书籍
		BookRouter.PUT("/star", book.StarBook)

		//获取相应标签的书籍列表
		BookRouter.GET("/label", book.Label)

	}
	CommentRouter := router.Group("/comment")
	{
		//获取某本书下的所有书评
		CommentRouter.GET("/comment/:book_id", comment.GetBookComment)

		//给某本书写书评
		CommentRouter.POST("/comment/:book_id", comment.GiveBookComment)

		//删除书评
		CommentRouter.DELETE("/comment/:comment_id", comment.DeleteComment)

		//更新书评
		CommentRouter.PUT("/comment/:comment_id", comment.UpdateBookComment)

	}

	err := router.Run(":8084")
	if err != nil {
		return
	}
}
