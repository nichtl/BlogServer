package initialize

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

type Server interface {
	ListenAndServe() error
}

func InitAndStartServeUseGin(port string, engine *gin.Engine) Server {
	s := endless.NewServer(port, engine)
	s.ReadHeaderTimeout = time.Second * 10
	s.WriteTimeout = time.Second * 10
	s.MaxHeaderBytes = 1 << 20
	return s
}
