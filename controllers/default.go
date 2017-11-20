package controllers

import (
	"github.com/astaxie/beego"
	"net"
	"os"
)

type MainController struct {
	beego.Controller
}
type UserController struct {
	beego.Controller
}

func (main *UserController) HelloSitepoint() {
	main.Data["Website"] = "My Website"
	main.Data["Email"] = "hy352144278@gmail.com"
	main.Data["EmailName"] = "JONY"
	main.TplName = "default/hello-sitepoint.tpl"
	main.Data["Id"] = main.Ctx.Input.Param(":id")
	req := main.Ctx.Request
	//ip := getLocalIP()
	dial, _ := net.Dial("tcp", "baidu.com:80")
	//network := dial.RemoteAddr().Network()
	ip := dial.LocalAddr().String()
	//ip := req.Header.Get("Host")
	Host, _ := os.Hostname()
	UserAgent := req.Header.Get("User-Agent")
	main.Data["UserAgent"] = UserAgent
	//main.Data["network"] = network
	main.Data["user_ip"] = ip
	main.Data["Host"] = Host
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "default/hello-sitepoint.tpl"
	c.Data["Id"] = c.Ctx.Input.Param(":id")
}

func getLocalIP() string {
	addrSlice, err := net.InterfaceAddrs()
	if err != nil {
		return "localhost"
	}
	for _, addr := range addrSlice {
		ipnet, ok := addr.(*net.IPNet)
		if ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "localhost"
}
