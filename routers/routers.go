package routers

import (
	"github.com/gin-gonic/gin"
	v1 "yhdm/api/v1"
)

func Routers()  *gin.Engine{
	r :=gin.Default()
	// 添加路由
	r.GET("/home",v1.Testget)

	
	return r
}
