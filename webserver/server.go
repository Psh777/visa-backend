package webserver

import (
	"fmt"
	"github.com/Psh777/visa-backend/modules/config"
	"github.com/Psh777/visa-backend/webserver/handlers"
	"github.com/gin-gonic/gin"
)

func Init(config config.MyConfig) {
	r := gin.Default()

	r.GET("/", handlers.IndexHandler)
	r.OPTIONS("/form/message", optionsHandler)
	r.POST("/form/message", handlers.FormHandler)

	_ = r.Run(fmt.Sprintf(":%v", config.Env.BackendPort))

}

func optionsHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "content-type,X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin, Authorization, Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	c.Writer.WriteHeader(200)
}
