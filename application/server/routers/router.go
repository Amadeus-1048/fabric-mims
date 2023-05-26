package routers

import (
	v2 "application/api/v2"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由信息
func InitRouter() *gin.Engine {
	r := gin.Default()

	apiV2 := r.Group("/api/v2")
	{
		apiV2.GET("/hello", v2.Hello)
		apiV2.POST("/createAccountV2", v2.CreateAccountV2)
		apiV2.POST("/queryAccountV2List", v2.QueryAccountV2List)
		apiV2.POST("/createPrescription", v2.CreatePrescription)
		apiV2.POST("/queryPrescription", v2.QueryPrescriptionList)
		apiV2.POST("/createInsuranceCover", v2.CreateInsuranceCover)
		apiV2.POST("/queryInsuranceCoverList", v2.QueryInsuranceCoverList)
		apiV2.POST("/updateInsuranceCover", v2.UpdateInsuranceCover)
		apiV2.POST("/deleteInsuranceCover", v2.DeleteInsuranceCover)
		apiV2.POST("/createDrugOrder", v2.CreateDrugOrder)
		apiV2.POST("/queryDrugOrderList", v2.QueryDrugOrderList)
	}
	return r
}
