package helper

import (
	sessionutil "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SaveSession(g *gin.Context, email string, userID uint) error {
	session := sessionutil.Default(g)
	session.Set("email", email)
	session.Set("userID", userID)
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}
