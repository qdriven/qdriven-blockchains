package router

import (
	v1 "chains-gotest-backend/api/v1"
	"chains-gotest-backend/middleware"
	"github.com/gin-gonic/gin"
)

func InitSysDictionaryRouter(Router *gin.RouterGroup) {
	SysDictionaryRouter := Router.Group("sysDictionary").Use(middleware.OperationRecord())
	{
		SysDictionaryRouter.POST("createSysDictionary", v1.CreateSysDictionary)   // 新建SysDictionary
		SysDictionaryRouter.DELETE("deleteSysDictionary", v1.DeleteSysDictionary) // 删除SysDictionary
		SysDictionaryRouter.PUT("updateSysDictionary", v1.UpdateSysDictionary)    // 更新SysDictionary
		SysDictionaryRouter.GET("findSysDictionary", v1.FindSysDictionary)        // 根据ID获取SysDictionary
		SysDictionaryRouter.GET("getSysDictionaryList", v1.GetSysDictionaryList)  // 获取SysDictionary列表
	}
}
