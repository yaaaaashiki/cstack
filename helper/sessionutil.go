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

func GetSession(g *gin.Context) (email string, userID uint, error error) {
	session := sessions.Default(g)
	emailValue := session.Get("email")
	userIDValue := session.Get("userID")
	if emailValue == nil || userIDValue == nil {
		return "", 0, errors.New("It has not been authenticated")
	}
	return emailValue.(string), userIDValue.(uint), nil
}

func SaveSession(g *gin.Context, email string, userID uint) error {
	session := sessions.Default(g)
	session.Set("email", email)
	session.Set("userID", userID)
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}
