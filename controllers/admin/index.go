package admin

import "jasdo/controllers"

type IndexController struct {
	controllers.BaseController
}

// @router / [get]
func (b *IndexController) Get() {
	b.Data["json"] = "pong"
	b.ServeJSON()
}
