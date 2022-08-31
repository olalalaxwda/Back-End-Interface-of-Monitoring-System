package controller

import (
	"github.com/gin-gonic/gin"
	"monitoring/orm"
	"monitoring/orm/tables"
	"net/http"
)

func Post(r *gin.Engine) {

	r.POST("/monitoringapi/jspe", func(context *gin.Context) {
		jspeCpstate := &struct {
			tables.Jspe
			StateAr []tables.Cpstate `json:"stateAr"`
		}{}
		context.ShouldBindJSON(jspeCpstate)

		jspe := &tables.Jspe{
			Filename: jspeCpstate.Filename,
			Lineno:   jspeCpstate.Lineno,
			Message:  jspeCpstate.Message,
		}
		orm.Db.Create(jspe)
		cpstate := &jspeCpstate.StateAr
		for i := 0; i < len(*cpstate); i++ {
			(*cpstate)[i].ErrorId = jspe.ID
			(*cpstate)[i].Type = "jspe"
		}
		orm.Db.Create(&cpstate)
		context.JSON(http.StatusOK, gin.H{
			"id": "ok",
		})
	})

	r.POST("/monitoringapi/getjspe", func(context *gin.Context) {
		jspes := []tables.Jspe{}
		orm.Db.Find(&jspes)
		context.JSON(http.StatusOK, gin.H{
			"jspe": jspes,
		})
	})

	r.POST("/monitoringapi/rse", func(context *gin.Context) {
		rseCpstate := &struct {
			tables.Rse
			StateAr []tables.Cpstate `json:"stateAr"`
		}{}
		context.ShouldBindJSON(rseCpstate)

		rse := &tables.Rse{
			TargetClassName: rseCpstate.TargetClassName,
			Url:             rseCpstate.Url,
		}
		orm.Db.Create(rse)
		cpstate := &rseCpstate.StateAr
		for i := 0; i < len(*cpstate); i++ {
			(*cpstate)[i].ErrorId = rse.ID
			(*cpstate)[i].Type = "rse"
		}
		orm.Db.Create(&cpstate)
		context.JSON(http.StatusOK, gin.H{
			"id": "ok",
		})
	})

	r.POST("/monitoringapi/getrse", func(context *gin.Context) {
		rse := []tables.Rse{}
		orm.Db.Find(&rse)
		context.JSON(http.StatusOK, gin.H{
			"rse": rse,
		})
	})

	r.POST("/monitoringapi/reject", func(context *gin.Context) {
		rejectCpstate := &struct {
			tables.Reject
			StateAr []tables.Cpstate `json:"stateAr"`
		}{}
		context.ShouldBindJSON(rejectCpstate)

		reject := &tables.Reject{
			Reason: rejectCpstate.Reason,
		}
		orm.Db.Create(reject)
		cpstate := &rejectCpstate.StateAr
		for i := 0; i < len(*cpstate); i++ {
			(*cpstate)[i].ErrorId = reject.ID
			(*cpstate)[i].Type = "reject"
		}
		orm.Db.Create(&cpstate)
		context.JSON(http.StatusOK, gin.H{
			"id": "ok",
		})
	})

	r.POST("/monitoringapi/getreject", func(context *gin.Context) {
		reject := []tables.Reject{}
		orm.Db.Find(&reject)
		context.JSON(http.StatusOK, gin.H{
			"reject": reject,
		})
	})

	r.POST("/monitoringapi/apisucceed", func(context *gin.Context) {
		apiSucceedCpstate := &struct {
			tables.ApiSucceed
			StateAr []tables.Cpstate `json:"stateAr"`
		}{}
		context.ShouldBindJSON(apiSucceedCpstate)

		apiSucceed := &tables.ApiSucceed{
			Status: apiSucceedCpstate.Status,
			Url:    apiSucceedCpstate.Url,
		}
		orm.Db.Create(apiSucceed)
		cpstate := &apiSucceedCpstate.StateAr
		for i := 0; i < len(*cpstate); i++ {
			(*cpstate)[i].ErrorId = apiSucceed.ID
			(*cpstate)[i].Type = "apiSucceed"
		}
		orm.Db.Create(&cpstate)
		context.JSON(http.StatusOK, gin.H{
			"id": "ok",
		})
	})

	r.POST("/monitoringapi/getapisucceed", func(context *gin.Context) {
		apiSucceed := []tables.ApiSucceed{}
		orm.Db.Find(&apiSucceed)
		context.JSON(http.StatusOK, gin.H{
			"apisucceed": apiSucceed,
		})
	})

	r.POST("/monitoringapi/apifailed", func(context *gin.Context) {
		apiFailedCpstate := &struct {
			tables.ApiFailed
			StateAr []tables.Cpstate `json:"stateAr"`
		}{}
		context.ShouldBindJSON(apiFailedCpstate)

		apiFailed := &tables.ApiFailed{
			Code:    apiFailedCpstate.Code,
			Message: apiFailedCpstate.Message,
			Url:     apiFailedCpstate.Url,
		}
		orm.Db.Create(apiFailed)
		cpstate := &apiFailedCpstate.StateAr
		for i := 0; i < len(*cpstate); i++ {
			(*cpstate)[i].ErrorId = apiFailed.ID
			(*cpstate)[i].Type = "apiFailed"
		}
		orm.Db.Create(&cpstate)
		context.JSON(http.StatusOK, gin.H{
			"id": "ok",
		})
	})

	r.POST("/monitoringapi/getapifailed", func(context *gin.Context) {
		apiFailed := []tables.ApiFailed{}
		orm.Db.Find(&apiFailed)
		context.JSON(http.StatusOK, gin.H{
			"apifailed": apiFailed,
		})
	})

	r.POST("/monitoringapi/ws", func(context *gin.Context) {
		ws := &tables.Ws{}
		context.ShouldBindJSON(ws)
		orm.Db.Create(ws)
		context.JSON(http.StatusOK, gin.H{
			"id": "ok",
		})
	})

	r.POST("/monitoringapi/getws", func(context *gin.Context) {
		ws := []tables.Ws{}
		orm.Db.Find(&ws)
		context.JSON(http.StatusOK, gin.H{
			"ws": ws,
		})
	})

	r.POST("/monitoringapi/keyperformancedata", func(context *gin.Context) {

		keyPerformanceData := &tables.KeyPerformanceData{}
		context.ShouldBindJSON(keyPerformanceData)
		orm.Db.Create(keyPerformanceData)
		context.JSON(http.StatusOK, gin.H{
			"id": "ok",
		})
	})

	r.POST("/monitoringapi/getkeyperformancedata", func(context *gin.Context) {
		keyPerformanceData := []tables.KeyPerformanceData{}
		orm.Db.Find(&keyPerformanceData)
		context.JSON(http.StatusOK, gin.H{
			"keyperformancedata": keyPerformanceData,
		})
	})

	r.POST("/monitoringapi/ps", func(context *gin.Context) {
		ps := &tables.Ps{}
		context.ShouldBindJSON(ps)
		orm.Db.Create(ps)
		context.JSON(http.StatusOK, gin.H{
			"id": "ok",
		})
	})

	r.POST("/monitoringapi/getps", func(context *gin.Context) {
		ps := []tables.Ps{}
		orm.Db.Find(&ps)
		context.JSON(http.StatusOK, gin.H{
			"ps": ps,
		})
	})

	r.POST("/monitoringapi/pvInfo", func(context *gin.Context) {
		pv := &tables.Pv{}
		context.ShouldBindJSON(pv)
		orm.Db.Create(pv)
		context.JSON(http.StatusOK, gin.H{
			"id": "ok",
		})
	})

	r.POST("/monitoringapi/getpvInfo", func(context *gin.Context) {
		pv := []tables.Pv{}
		orm.Db.Find(&pv)
		context.JSON(http.StatusOK, gin.H{
			"pvInfo": pv,
		})
	})

	r.POST("/monitoringapi/uvInfo", func(context *gin.Context) {
		uv := &tables.Uv{}
		context.ShouldBindJSON(uv)
		orm.Db.Create(uv)
		context.JSON(http.StatusOK, gin.H{
			"id": "ok",
		})
	})

	r.POST("/monitoringapi/getuvInfo", func(context *gin.Context) {
		uv := []tables.Uv{}
		orm.Db.Find(&uv)
		context.JSON(http.StatusOK, gin.H{
			"uvInfo": uv,
		})
	})

	r.POST("/monitoringapi/getcpstate", func(context *gin.Context) {
		cpstate := []tables.Cpstate{}
		orm.Db.Find(&cpstate)
		context.JSON(http.StatusOK, gin.H{
			"cpstate": cpstate,
		})
	})

}
