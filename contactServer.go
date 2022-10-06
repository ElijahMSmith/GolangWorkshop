package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Contact struct {
	id          int
	firstName   string
	lastName    string
	phone       string
	description string
	onShitList  bool
}

var nextId int = 0
var contacts []Contact

func newContact(c *gin.Context) {
	fmt.Println("New Contact")

	// curl -d "firstName=Eli&lastName=Smith" -X POST http://localhost:8080/contact
	shitListed, _ := strconv.ParseBool(c.PostForm("shitListed"))

	contacts = append(contacts, Contact{
		firstName:   c.PostForm("firstName"),
		lastName:    c.PostForm("lastName"),
		phone:       c.PostForm("phone"),
		description: c.PostForm("description"),
		onShitList:  shitListed,
		id:          nextId,
	})

	fmt.Println(contacts)

	c.JSON(http.StatusAccepted, nextId)
	nextId++
}

func getContact(c *gin.Context) {
	// TODO
	fmt.Println("Get Contact")

}

func deleteContact(c *gin.Context) {
	// TODO
	fmt.Println("Delete Contact")

}

func updateContact(c *gin.Context) {
	// TODO
	fmt.Println("Update Contact")

}
