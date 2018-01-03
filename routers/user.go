package routers

import (
	"curso/controllers"

	"github.com/labstack/echo"
)

// App instancia de Echo
var App *echo.Echo

func init() {
	App = echo.New()
	// Pagina inicial
	App.GET("/", controllers.Home)

	api := App.Group("/v1")
	api.POST("/user", controllers.Insert)
	api.GET("/user", controllers.Listar)
	api.DELETE("/user/:id", controllers.Deletar)
	api.PUT("/user/:id", controllers.Alterar)
}
