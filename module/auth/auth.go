package auth

import (
	//"context"

	"jasdo/models"

	beego "github.com/beego/beego"
	"github.com/beego/beego/context"
	"github.com/beego/beego/session"
)

// login user
func LoginUser(user *models.Users, ctx *context.Context, remember bool) {
	// werid way of beego session regenerate id...
	beego.GlobalSessions.SessionRegenerateID(ctx.ResponseWriter, ctx.Request)
	ctx.Input.CruSession.SessionRelease(ctx.ResponseWriter)
	ctx.Input.CruSession.Set("auth_user_id", user.Id)
}

// logout user
func LogoutUser(ctx *context.Context) {
	ctx.Input.CruSession.Delete("auth_user_id")
	ctx.Input.CruSession.Flush()
	beego.GlobalSessions.SessionDestroy(ctx.ResponseWriter, ctx.Request)
}

func GetUserIdFromSession(sess session.Store) int {
	if id, ok := sess.Get("auth_user_id").(int); ok && id > 0 {
		return id
	}
	return 0
}

// get user if key exist in session
func GetUserFromSession(user *models.Users, sess session.Store) bool {
	id := GetUserIdFromSession(sess)
	if id > 0 {
		u, err := models.GetUsersById(id)
		if err == nil {
			user = u
			return true
		}
	}
	return false
}

// verify username/email and password
func VerifyUser(user *models.Users, username, password string) (success bool) {
	// search user by username or email

	return
}
