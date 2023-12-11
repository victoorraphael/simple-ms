package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/victoorraphael/simple-ms/users/service"
	"log"
	"net/http"
	"strconv"
)

type Provider struct {
	SRV service.UserService
}

// FetchAll list all users
func (p Provider) FetchAll(c *gin.Context) {
	list, err := p.SRV.List(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, list)
}

// FetchOne returns a user by id
func (p Provider) FetchOne(c *gin.Context) {
	idSTR := c.Param("id")
	id, err := strconv.Atoi(idSTR)
	if err != nil {
		log.Printf("failed to convert ID: err: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "failed to retrieve user by id",
		})
		return
	}

	user, err := p.SRV.Get(c, int64(id))
	if err != nil {
		log.Printf("failed to get user: err: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "failed to retrieve user by id",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
