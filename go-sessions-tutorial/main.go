package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var (
	//RedisHost ...
	RedisHost = os.Getenv("REDISHOST")
	//RedisPort ...
	RedisPort = os.Getenv("REDISPORT")
)

type UserController struct{}

func (ctr UserController) Signin(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("user_id", 1)
	session.Set("user_email", "demo@demo.com")
	session.Set("user_username", "demo")
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "User signed in", "user": "demo"})
}

func (ctrl UserController) Signout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "Signed out..."})
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("user_id")
		if sessionID == nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "not authed",
			})
			c.Abort()
		}
	}
}

func main() {

	r := gin.Default()

	store, _ := redis.NewStore(10, "tcp", RedisHost+":"+RedisPort, "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	userController := new(UserController)
	r.POST("/signin", userController.Signin)
	r.GET("/signout", userController.Signout)

	auth := r.Group("/auth")
	auth.Use(AuthRequired())
	{
		auth.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	r.Run()
}
