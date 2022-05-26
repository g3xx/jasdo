package admin

import (
	"encoding/json"
	"jasdo/controllers"
	"jasdo/models"
	"jasdo/module/msg"
	"strconv"
	"strings"
)

// RolesController operations for Roles
type RolesController struct {
	controllers.BaseController
}

// URLMapping ...
func (c *RolesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Roles
// @Param	body		body 	models.Roles	true		"body for Roles content"
// @Success 201 {int} models.Roles
// @Failure 403 body is empty
// @router / [post]
func (c *RolesController) Post() {
	var v models.Roles

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {

		if m := c.CeckPayload(v); m != "" {
			c.Data["json"] = msg.Pesan("error", "Pleas Ceck field "+m+" Not Valid")
			return
		}

		if _, err := models.AddRoles(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = msg.Pesan("success", "Rolles : "+v.Name+" success Createad  !")
		} else {
			c.Data["json"] = msg.Pesan("error", err.Error())
		}
	} else {
		c.Data["json"] = msg.Pesan("error", err.Error())
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Roles
// @Param	q		query	string	false	"mode index roles render"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Roles
// @Failure 403
// @router / [get]
func (c *RolesController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("q"); v != "all" {
		c.ActiveContent("admin/roles/index")
	} else {
		if v := c.GetString("fields"); v != "" {
			fields = strings.Split(v, ",")
		}
		// limit: 10 (default is 10)
		if v, err := c.GetInt64("limit"); err == nil {
			limit = v
		}
		// offset: 0 (default is 0)
		if v, err := c.GetInt64("offset"); err == nil {
			offset = v
		}
		// sortby: col1,col2
		if v := c.GetString("sortby"); v != "" {
			sortby = strings.Split(v, ",")
		}
		// order: desc,asc
		if v := c.GetString("order"); v != "" {
			order = strings.Split(v, ",")
		}
		// query: k:v,k:v
		if v := c.GetString("query"); v != "" {
			for _, cond := range strings.Split(v, ",") {
				kv := strings.SplitN(cond, ":", 2)
				if len(kv) != 2 {
					c.Data["json"] = msg.Pesan("error", "Error: invalid query key/value pair")
					c.ServeJSON()
					return
				}
				k, v := kv[0], kv[1]
				query[k] = v
			}
		}

		l, err := models.GetAllRoles(query, fields, sortby, order, offset, limit)
		if err != nil {
			c.Data["json"] = msg.Pesan("error", err.Error())
		} else {
			c.Data["json"] = l
		}
		c.ServeJSON()
	}
}

// Put ...
// @Title Put
// @Description update the Roles
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Roles	true		"body for Roles content"
// @Success 200 {object} models.Roles
// @Failure 403 :id is not int
// @router /:id [put]
func (c *RolesController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Roles{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateRolesById(&v); err == nil {
			c.Data["json"] = msg.Pesan("success", "Rolles : "+v.Name+" Updated !")
		} else {
			c.Data["json"] = msg.Pesan("error", "Failed Update Roles")
		}
	} else {
		c.Data["json"] = msg.Pesan("error", err.Error())
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Roles
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *RolesController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteRoles(id); err == nil {
		c.Data["json"] = msg.Pesan("success", "Rolles : Deleted Success !")
	} else {
		c.Data["json"] = msg.Pesan("error", err.Error())
	}
	c.ServeJSON()
}
