package routes

import (
	"user-service/controllers"
	routes "user-service/routes/user"

	"github.com/gin-gonic/gin"
)

type Registry struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
}

type IRouterRegister interface {
	Serve()
}

func NewRouteRegistry(controller controllers.IControllerRegistry, group *gin.RouterGroup) IRouterRegister {
	return &Registry{controller: controller, group: group}
}

func (r *Registry) userRoute() routes.IUserRoute {
	return routes.NewUserRoute(r.controller, r.group)
}

func (r *Registry) Serve() {
	r.userRoute().Run()
}
