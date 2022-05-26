package admin

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"jasdo/controllers"
	"jasdo/models"
	"jasdo/module/hash"
	"jasdo/module/msg"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

type UsersController struct {
	controllers.BaseController
}

// URLMapping ...
func (this *UsersController) URLMapping() {
	this.Mapping("Post", this.Post)
	this.Mapping("Get", this.GetAll)
	this.Mapping("Put", this.Put)
	this.Mapping("Delete", this.Delete)
}

//func (this *UsersController) NestPrepare() {
// set up for every route in users ctrl.
// will check if request no have session aka no login
//	this.IsSecured()
//}

// Post ...
// @Title Post
// @Description create Users
// @Param	body		body 	models.Users	true		"body for Users content"
// @Success 201 {int} models.Users
// @Failure 403 body is empty
// @router / [post]
func (this *UsersController) Post() {

	var v models.Users
	defer this.ServeJSON()

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &v); err == nil {
		v.Password = hash.Generate(v.Password)
	} else {
		this.Data["json"] = msg.Pesan("error", err.Error())
	}

	if m := this.CeckPayload(v); m != "" {
		this.Data["json"] = msg.Pesan("error", m)
		return
	}

	if u, err := models.GetUsersByEmail(v.Email); err == nil {
		this.Data["json"] = msg.Pesan("error", "Email : "+u.Email+" Already Registered, Pleas Forgot Password.")
		return
	}

	if _, err := models.AddUsers(&v); err == nil {
		this.Ctx.Output.SetStatus(201)
		this.Data["json"] = msg.Pesan("success", "Email : "+v.Email+" success Createad !")
	} else {
		this.Data["json"] = msg.Pesan("error", err.Error())
	}
	this.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Users
// @Param	email		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 email is empty
// @router /:id [delete]
func (this *UsersController) Delete() {
	defer this.ServeJSON()
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	errs := validate.Var(id, "number")
	if errs != nil {
		// filter if param not valid number set (404)
		this.Data["json"] = msg.Pesan("error", "Failed delete "+errs.Error())
	}
	// always filter coz to many hacker !!!!
	l, err := models.GetUsersById(id)
	if err != nil {
		this.Data["json"] = msg.Pesan("error", "User "+idStr+" Not Found")
	}

	if err := models.DeleteUsers(id); err == nil {
		this.Data["json"] = msg.Pesan("success", "Success Delete "+l.Email)
	} else {
		this.Data["json"] = msg.Pesan("error", "Failed delete "+err.Error())
	}

	this.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Users
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Users	true		"body for Users content"
// @Success 200 {object} models.Users
// @Failure 403 :sid is not int
// @router /:sid [put]
func (this *UsersController) Put() {

	defer this.ServeJSON()
	idStr := this.Ctx.Input.Param(":sid")
	id, _ := strconv.Atoi(idStr)
	errs := validate.Var(id, "number")
	fmt.Println(id)
	if errs != nil {
		// filter if param not valid Int set (404)
		this.Abort("404")
	}
	n := models.Users{Id: id}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &n); err != nil {
		this.Data["json"] = msg.Pesan("error", err.Error())
	}

	l, err := models.GetUsersById(id)
	if err != nil {
		this.Data["json"] = msg.Pesan("error", "User Id Not Found")
	}

	if n.Password != "" {
		n.Password = hash.Generate(n.Password)
	} else {
		n.Password = l.Password
	}

	// i want validate object , before to execute db
	if m := this.CeckPayload(n); m != "" {
		this.Data["json"] = msg.Pesan("error", "Pleas Ceck field "+m+" Not Valid")
		return
	}
	// filter if user replace other user email
	if l.Email != n.Email {
		u, err := models.GetUsersByEmail(n.Email)
		if err == nil {
			this.Data["json"] = msg.Pesan("error", "Invalid Update Email "+u.Email+" Has already Use.")
			return
		}
	}

	if err := models.UpdateUsersById(&n); err == nil {
		this.Data["json"] = msg.Pesan("Success", "update user "+n.Email)
	} else {
		this.Data["json"] = msg.Pesan("Error", "Failed Update User !"+err.Error())
	}
	this.ServeJSON()
}

// @Title Get All
// @Description get all users support filter with param
// @Param	q		query	string	false	"query for defined index"
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Users
// @Failure 403
// @router / [get]
func (this *UsersController) GetAll() {

	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	if v := this.GetString("q"); v != "all" {
		this.ActiveContent("admin/users/index")
	} else {
		// fields: col1,col2,entity.col3
		// [3]string{"Japan", "Australia", "Germany"}
		defer this.ServeJSON()
		fields := []string{"Id", "FullName", "Email", "Telp", "RolesId"}

		// limit: 10 (default is 10)
		if v, err := this.GetInt64("limit"); err == nil {
			limit = v
		}
		// offset: 0 (default is 0)
		if v, err := this.GetInt64("offset"); err == nil {
			offset = v
		}
		// sortby: col1,col2
		if v := this.GetString("sortby"); v != "" {
			sortby = strings.Split(v, ",")
		}
		// order: desc,asc
		if v := this.GetString("order"); v != "" {
			order = strings.Split(v, ",")
		}
		// query: k:v,k:v
		if v := this.GetString("query"); v != "" {
			for _, cond := range strings.Split(v, ",") {
				kv := strings.SplitN(cond, ":", 2)
				if len(kv) != 2 {
					this.Data["json"] = msg.Pesan("Error", "Error: invalid query key/value pair ")
					this.ServeJSON()
					return
				}
				k, v := kv[0], kv[1]
				query[k] = v
			}
		}

		l, err := models.GetAllUsers(query, fields, sortby, order, offset, limit)
		if err != nil {
			this.Data["json"] = msg.Pesan("Error", "Failed Get All Users !")
		} else {
			v := models.CountTable(new(models.Users))
			this.Data["json"] = msg.DbPesan(l, v)
		}
		this.ServeJSON()
	}
}
