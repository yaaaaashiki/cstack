package helper

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ClearSession(g *gin.Context) {
	session := sessions.Default(g)
	session.Clear()
	session.Save()
}

func GetSession(g *gin.Context) (email string, userId uint, error error) {
	session := sessions.Default(g)
	emailValue := session.Get("email")
	userIdValue := session.Get("userId")
	if emailValue == nil || userIdValue == nil {
		return "", 0, errors.New("It has not been authenticated")
	}
	return emailValue.(string), userIdValue.(uint), nil
}

func SaveSession(g *gin.Context, email string, userId uint) {
	session := sessions.Default(g)
	session.Set("email", email)
	session.Set("userId", uint(userId))
	session.Save()
}
