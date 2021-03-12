package routers

import (
	"bubble/controller"

	"github.com/gin-gonic/gin"
)

func V2Router(Router *gin.RouterGroup) {
	v2Group := Router.Group("v2")
	{
		// 待办事项
		// 添加
		v2Group.POST("/todo1", controller.CreateTodo)
		// 查看所有的待办事项
		v2Group.GET("/todo1", controller.GetTodoList)
		// 修改某一个待办事项
		v2Group.PUT("/todo1/:id", controller.UpdateATodo)
		// 删除某一个待办事项
		v2Group.DELETE("/todo1/:id", controller.DeleteATodo)
	}
	return
}
