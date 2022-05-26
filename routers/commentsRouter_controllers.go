package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["jasdo/controllers/admin:IndexController"] = append(beego.GlobalControllerRouter["jasdo/controllers/admin:IndexController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["jasdo/controllers/admin:RolesController"] = append(beego.GlobalControllerRouter["jasdo/controllers/admin:RolesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["jasdo/controllers/admin:RolesController"] = append(beego.GlobalControllerRouter["jasdo/controllers/admin:RolesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["jasdo/controllers/admin:RolesController"] = append(beego.GlobalControllerRouter["jasdo/controllers/admin:RolesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["jasdo/controllers/admin:RolesController"] = append(beego.GlobalControllerRouter["jasdo/controllers/admin:RolesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["jasdo/controllers/admin:UsersController"] = append(beego.GlobalControllerRouter["jasdo/controllers/admin:UsersController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["jasdo/controllers/admin:UsersController"] = append(beego.GlobalControllerRouter["jasdo/controllers/admin:UsersController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["jasdo/controllers/admin:UsersController"] = append(beego.GlobalControllerRouter["jasdo/controllers/admin:UsersController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["jasdo/controllers/admin:UsersController"] = append(beego.GlobalControllerRouter["jasdo/controllers/admin:UsersController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:sid",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["jasdo/controllers/api:RolesController"] = append(beego.GlobalControllerRouter["jasdo/controllers/api:RolesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["jasdo/controllers/api:RolesController"] = append(beego.GlobalControllerRouter["jasdo/controllers/api:RolesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["jasdo/controllers/api:RolesController"] = append(beego.GlobalControllerRouter["jasdo/controllers/api:RolesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["jasdo/controllers/api:RolesController"] = append(beego.GlobalControllerRouter["jasdo/controllers/api:RolesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["jasdo/controllers/api:RolesController"] = append(beego.GlobalControllerRouter["jasdo/controllers/api:RolesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["jasdo/controllers/api:UsersController"] = append(beego.GlobalControllerRouter["jasdo/controllers/api:UsersController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["jasdo/controllers/api:UsersController"] = append(beego.GlobalControllerRouter["jasdo/controllers/api:UsersController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["jasdo/controllers/api:UsersController"] = append(beego.GlobalControllerRouter["jasdo/controllers/api:UsersController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["jasdo/controllers/api:UsersController"] = append(beego.GlobalControllerRouter["jasdo/controllers/api:UsersController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["jasdo/controllers/api:UsersController"] = append(beego.GlobalControllerRouter["jasdo/controllers/api:UsersController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
