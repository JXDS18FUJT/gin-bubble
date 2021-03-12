package routers

import (
	"bubble/controller"

	"github.com/gin-gonic/gin"
)

func StudyRouter(Router *gin.RouterGroup) {
	// v1
	StudyGroup := Router.Group("study")
	{
		// 待办事项
		// 添加
		StudyGroup.GET("/todo", controller.CreateStudy)
		// // 查看所有的待办事项
		// StudyGroup.GET("/todo", controller.GetTodoList)
		// // 修改某一个待办事项
		// StudyGroup.PUT("/todo/:id", controller.UpdateATodo)
		// // 删除某一个待办事项
		// StudyGroup.DELETE("/todo/:id", controller.DeleteATodo)
	}

}
