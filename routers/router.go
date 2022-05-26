package routers

import (
	"jasdo/controllers"
	"jasdo/controllers/admin"
	"jasdo/controllers/api"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/roles",
			beego.NSInclude(
				&api.RolesController{},
			),
		),

		beego.NSNamespace("/users",
			beego.NSInclude(
				&api.UsersController{},
			),
		),
	)

	ab := beego.NewNamespace("/admin",
		beego.NSRouter("/", &admin.IndexController{}),
		beego.NSRouter("/login/", &admin.LoginController{}),
		beego.NSRouter("/logout/", &admin.LoginController{}, "get:Logout"),
		beego.NSNamespace("/users",
			beego.NSInclude(
				&admin.UsersController{},
			),
		),
		beego.NSNamespace("/roles",
			beego.NSInclude(
				&admin.RolesController{},
			),
		),
	)

	beego.AddNamespace(ns, ab)
	beego.Router("/", &controllers.MainController{})

	beego.InsertFilter("/admin/*", beego.BeforeRouter, controllers.IamSecurity)
}
