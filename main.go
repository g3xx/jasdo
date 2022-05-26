package main

import (
	"fmt"
	_ "jasdo/routers"
	"net/http"
	"text/template"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	c, err := beego.AppConfig.String("sqlconn")
	if err != nil {
		fmt.Println("wrong connection db")
	}
	orm.RegisterDataBase("default", "mysql", c)

}
func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/404.html")
	t.Execute(rw, "")
}
func main() {
	beego.ErrorHandler("404", page_not_found)
	beego.Run()
}
