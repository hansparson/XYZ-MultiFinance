package blueprints

import (
	dbcontroller "xyz-multifinance/db/db-controller"
	service "xyz-multifinance/services"

	"github.com/gin-gonic/gin"
)

type Router struct {
	control *dbcontroller.HandlersController
}

func ServiceRouther(control *dbcontroller.HandlersController) *Router {
	return &Router{control: control}
}

func (r *Router) Start(port string) {
	router := gin.New()

	router.POST("/create-user", func(ctx *gin.Context) {
		service.CreateUser(ctx, r.control)
	})

	router.POST("/transaction", func(ctx *gin.Context) {
		service.Transaction(ctx, r.control)
	})

	router.Run(port)
}
