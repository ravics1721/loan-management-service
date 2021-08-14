package server

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/ravics1721/loan-management-service/config"
	"github.com/ravics1721/loan-management-service/routes"
)

var (
	once    sync.Once
	cfg     *config.Configuration
	ginMode string
)

func Init() {
	const production string = "production"
	once.Do(func() {
		cfg = config.Config()
		if cfg.Mode == production {
			ginMode = gin.ReleaseMode
		}
	})
}

func Start() {
	r := routes.Router(ginMode)
	err := r.Run(":" + strconv.Itoa(cfg.Port))
	if err != nil {
		fmt.Printf("Error in Starting the HTTP Server, Err: %s", err.Error())
		return
	}
	fmt.Print("Initializing routes... 🚀 ")
	fmt.Print("Mode:: ", ginMode)
}
