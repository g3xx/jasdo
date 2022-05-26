package models

import "github.com/beego/beego/v2/client/orm"

func CountTable(table interface{}) int64 {

	o := orm.NewOrm()
	q, _ := o.QueryTable(table).RelatedSel().Count()
	return q

}
