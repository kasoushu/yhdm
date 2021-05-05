package global

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"yhdm/routers"
)

type server interface {
	ListenAndServe() error
}
func InitGlobal(add string)  {

	Router :=routers.Routers()
	s := initServer(add,Router)
	LOG.Error(s.ListenAndServe().Error())
}

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
