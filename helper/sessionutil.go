package helper

import (
	"errors"

	sessionutil "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ClearSession(g *gin.Context) {
	session := sessionutil.Default(g)
	session.Clear()
	session.Save()
}

func GetSession(g *gin.Context) (email string, userId uint, error error) {
	session := sessionutil.Default(g)
	emailValue := session.Get("email")
	userIdValue := session.Get("userID")
	if emailValue == nil || userIdValue == nil {
		return "", 0, errors.New("It has not been authenticated")
	}
	return emailValue.(string), userIdValue.(uint), nil
}

func SaveSession(g *gin.Context, email string, userID uint) error {
	session := sessionutil.Default(g)
	session.Set("email", email)
	session.Set("userID", userID)
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}
