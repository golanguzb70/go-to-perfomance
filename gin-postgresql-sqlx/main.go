package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golanguzb70/go-to-perfomance/db"
)

func main() {
	r := gin.Default()

	d, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	r.GET("/users", func(c *gin.Context) {
		users, err := d.GetUsers()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, users)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		idInt, errconv := strconv.Atoi(id)
		if errconv != nil {
			c.JSON(400, gin.H{"error": errconv.Error()})
			return
		}

		user, err := d.GetUser(idInt)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, user)
	})

	r.POST("/users", func(c *gin.Context) {
		var user db.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := d.CreateUser(&user); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, user)
	})

	r.Run(":8888")
}
