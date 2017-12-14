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
}
