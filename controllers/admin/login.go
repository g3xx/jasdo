package admin

import (
	"encoding/json"
	"jasdo/controllers"
	"jasdo/models"
	"jasdo/module/hash"
	"jasdo/module/msg"

	"github.com/beego/beego/v2/client/orm"
)

type LoginController struct {
	controllers.BaseController
}

func (this *LoginController) Get() {
	//this.ActiveContent("admin/index")
	this.TplName = "admin/login.tpl"

	sess := this.GetSession("jasdo_admin")
	if sess != nil {
		this.Redirect("/admin/users/", 302)
		return
	}

}

func (this *LoginController) Post() {

	defer this.ServeJSON()
	var u map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &u)

	email := u["Email"].(string)
	password := u["Pass"].(string)

	mail := validate.Var(email, "required,email,max=50")
	pass := validate.Var(password, "required,max=30")

	if mail != nil && pass != nil {
		this.Data["json"] = msg.Pesan("error", "Pleas Correct Email and Password")
		return
	}

	// pass validate , run check email from db
	user, err := models.GetUsersByEmail(email)
	if err == orm.ErrNoRows {
		this.Data["json"] = msg.Pesan("error", "No such user/email")
		return
	}

	match := hash.CheckPassword(password, user.Password)
	if !match {
		this.Data["json"] = msg.Pesan("error", "Wrong Password")
		return
	}

	this.SetLogin(user)
	//this.Data["json"] = user
	//this.Redirect("/admin/users", 302)
	this.Data["json"] = msg.Pesan("success", "login success")
	this.ServeJSON()

}

func (this *LoginController) Logout() {
	this.DelLogin()
}
