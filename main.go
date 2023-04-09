package main

import (
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/gin-gonic/gin"
	"github.com/taskfactory/http_access/services/sources"
)

func main() {
	mux := &tars.TarsHttpMux{}
	r := mux.GetGinEngine()
	registerRESTRouter(r)

	// Get server config
	cfg := tars.GetServerConfig()
	tars.AddHttpServant(mux, cfg.App+"."+cfg.Server+".http")
	tars.Run()
}

func registerRESTRouter(r *gin.Engine) {
	r.GET("/source/list", sources.GetSources)
	r.POST("/source/upsert", sources.Upsert)
}
