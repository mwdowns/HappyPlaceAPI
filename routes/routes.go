package routes

import (
	"fmt"
	"happyplace/api/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// favicon stuff fix later
	router.GET("/favicon.ico", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{}) })
	// user signup and login
	router.POST("/signup", createUser)
	router.POST("/login", login)
}

func createObject(object interfaces.Saver, context *gin.Context, name string) {
	id, err := object.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save " + name, "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": name + " created", name: object})
	fmt.Printf("this is the object id: %v\n", id)
}
