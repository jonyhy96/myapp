package controllers

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"net"
	"os"
)

type MainController struct {
	beego.Controller
}
type UserController struct {
	beego.Controller
}
type PlayController struct {
	beego.Controller
}
type LoginController struct {
	beego.Controller
}
type SignupController struct {
	beego.Controller
}

var OEM = os.Getenv("OEM")
var VER = os.Getenv("VER")

func (c *PlayController) Play() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "play.tpl"
	c.Data["Id"] = c.Ctx.Input.Param(":id")
}

func (main *LoginController) Post() {
	username := main.GetString("username")
	password := main.GetString("password")
	main.TplName = "login.tpl"
	h := md5.New()
	h.Write([]byte(password))
	cipherStr := h.Sum(nil)
	password = hex.EncodeToString(cipherStr)
	db, err := linkTomysql()
	defer db.Close()
	if err != nil {
		glog.Info(err)
	} else {
		if username == "" {
			main.Redirect("/user/signup", 302)
			return
		} else {
			rows, err := db.Query("select username,password from customer where username=? and password=?", username, password)
			if err != nil {
				glog.Infoln("查询失败")
				glog.Info(err)
				main.Redirect("/user/signup", 302)
				return
			} else {
				if rows.Next() != false {
					main.SetSession("username", main.GetString("username"))
					main.SetSession("password", main.GetString("password"))
					rows, err = db.Query("select userinfo from customer where username=? ", username)
					if err != nil {
						glog.Info(err)
					}
					var s string
					for rows.Next() {
						if err := rows.Scan(&s); err != nil {
							glog.Info(err)
						}
					}
					main.SetSession("userinfo", s)
					glog.Infoln("登录成功！")
					main.Redirect("/user/profile", 302)
					return
				} else {
					glog.Infoln("登录失败！")
					main.Redirect("/user/signup", 302)
					return
				}
			}
		}
	}
}
func (main *LoginController) Get() {
	main.Data["OEM"] = OEM
	main.Data["VER"] = VER
	main.TplName = "login.tpl"
}
func (main *SignupController) Post() {
	db, err := linkTomysql()
	if err != nil {
		glog.Info(err)
	}
	defer db.Close()
	username := main.GetString("username")
	password := main.GetString("password")
	userinfo := main.GetString("userinfo")
	h := md5.New()
	h.Write([]byte(password))
	cipherStr := h.Sum(nil)
	db.Exec("insert into customer(username, password,userinfo) values(?,?,?)", username, hex.EncodeToString(cipherStr), userinfo)
	main.Redirect("/user/login", 302)
}

func (main *SignupController) Get() {
	main.Data["OEM"] = OEM
	main.Data["VER"] = VER
	main.TplName = "default/register.tpl"
}

func (main *UserController) Get() {
	main.Data["OEM"] = OEM
	main.Data["VER"] = VER
	main.Data["Username"] = main.GetSession("username")
	main.Data["Password"] = main.GetSession("password")
	main.Data["Userinfo"] = main.GetSession("userinfo")
	main.TplName = "default/hello-sitepoint.tpl"
	req := main.Ctx.Request
	dial, _ := net.Dial("tcp", "baidu.com:80")
	ip := dial.LocalAddr().String()
	Host, _ := os.Hostname()
	UserAgent := req.Header.Get("User-Agent")
	main.Data["UserAgent"] = UserAgent
	main.Data["user_ip"] = ip
	main.Data["Host"] = Host
}

func (c *MainController) Get() {
	c.Data["OEM"] = OEM
	c.Data["VER"] = VER
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
func linkTomysql() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.34.38)/myapp?charset=utf8")
	if err != nil {
		glog.Infoln("连接服务器失败！")
	}
	return db, nil
}
