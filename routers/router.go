package routers

import (
	"github.com/astaxie/beego"
	"myapp/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/profile", &controllers.UserController{})
	beego.Router("/play", &controllers.PlayController{}, "get:Play")
	beego.Router("/user/signup", &controllers.SignupController{})
	beego.Router("/user/login", &controllers.LoginController{})
}
