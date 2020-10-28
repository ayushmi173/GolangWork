//go mod init
//go get -u github.com/gin-gonic/gin

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Human struct {
	Id   string `json : "id" `
	Name string `json : "name"`
}

//my collection
var Collection []Human

func main() {
	Collection = append(Collection, Human{Id: "1", Name: "Ayush Mishra"})
	Collection = append(Collection, Human{Id: "2", Name: "Kashyap Ayush"})

	router := gin.Default()

	router.GET("/get", getting)
	router.POST("/post", creating)
	router.DELETE("/delete", deleting)
	router.PUT("/put", putting)

	router.Run(":5000") // hardcoded 5000 ports
}

func getting(c *gin.Context) {

	for i, item := range Collection {
		c.String(http.StatusOK, "%d element is: %#v \n", i+1, item)
	}
}

func creating(c *gin.Context) {

	id := c.Query("id")
	name := c.Query("name")

	Collection = append(Collection, Human{Id: id, Name: name})
	c.String(http.StatusOK, "%s is id %s is name", id, name)

}

func deleting(c *gin.Context) {

	id := c.Query("id")
	Intid, _ := strconv.Atoi(id)
	Collection = append(Collection[:Intid-1], Collection[Intid:]...)

}

func putting(c *gin.Context) {

	id := c.Query("id")
	name := c.Query("name")

	for i, item := range Collection {

		if item.Id == id {
			item.Name = name
			Collection[i] = item
			c.String(http.StatusOK, "%#v", Collection[i])
			return
		}
	}
}
