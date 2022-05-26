package controllers

import (
	"html/template"
	"jasdo/models"
	"strings"
	"time"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/go-playground/validator/v10"
)

type BaseController struct {
	beego.Controller

	Userinfo *AdminSession
	IsLogin  bool
}

type AdminSession struct {
	Email     string
	Username  string
	Roles     string
	timestamp time.Time
}

var (
	validate = validator.New()
)

// adminCtrl implements global settings for all other routers.
// Prepare implements Prepare method for baseRouter.

func (c *BaseController) Prepare() {
	//c.SetParams()

	c.IsLogin = c.GetSession("jasdo_admin") != nil
	if c.IsLogin {
		c.Userinfo = c.GetLogin()
	}

	c.Data["IsLogin"] = c.IsLogin
	c.Data["Userinfo"] = c.Userinfo
	c.Data["_csrf"] = c.XSRFToken()
	c.Data["_csrfhtml"] = template.HTML(`<meta name="_csrf" content="` + c.XSRFToken() + `"/>`)
}

func (c *BaseController) GetLogin() *AdminSession {
	return c.GetSession("jasdo_admin").(*AdminSession)
}

func (c *BaseController) DelLogin() {
	c.DelSession("jasdo_admin")
	c.Redirect("/admin/login", 302)
}

func (c *BaseController) SetLogin(user *models.Users) {

	m := &AdminSession{}
	m.Email = user.Email
	m.Username = user.FullName
	m.Roles = user.RolesId.Name
	m.timestamp = time.Now()
	c.SetSession("jasdo_admin", m)

}

func (c *BaseController) CeckPayload(s interface{}) string {

	var text string
	if err := validate.Struct(s); err != nil {
		//validate struck before proccess to database
		validationErrors := err.(validator.ValidationErrors)
		for _, validationError := range validationErrors {
			text = validationError.StructField()
		}
	}
	return text
}

func (this *BaseController) ActiveContent(view string) {
	this.Layout = "base.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Sidebar"] = "sidebar.tpl"
	this.LayoutSections["Footer"] = "footer.tpl"
	this.TplName = view + ".tpl"

}

func IamSecurity(ctx *context.Context) {
	// Don't redirect itself to prevent the circle
	if strings.HasPrefix(ctx.Input.URL(), "/admin/login") {
		return
	}
	_, ok := ctx.Input.Session("jasdo_admin").(*AdminSession)
	if !ok {
		ctx.Redirect(302, "/admin/login")
	}
}
