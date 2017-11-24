package routers

import (
	"github.com/astaxie/beego"
	"myapp/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/profile", &controllers.UserController{}, "get:HelloSitepoint")
	beego.Router("/play", &controllers.PlayController{}, "get:Play")
}
