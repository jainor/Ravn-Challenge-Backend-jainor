package main

import (
	"encoding/json"
	"log"
	"strconv"

	"net/http"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"

	dbent "internal/db/entities"
	mq "internal/messagequeue"
)

const page = "/authors"
const hellopage = "/hello"
const paramName = "count"

func main() {

	router := gin.Default()
	router.GET(page, getAuthors)
	router.GET(hellopage,getHello)
	//router.Run("localhost:8080")

	router.Run()
}
// getHello return a Hello message
func getHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "{hello with kubernetes}")
	return
}

// getAuthors responds with the list of all authors as JSON.
func getAuthors(c *gin.Context) {
	cnt := c.Query(paramName)

	icnt, err := strconv.Atoi(cnt)
	if err != nil {
		log.Println("Invalid argument")
		return
	}

	res, err := mq.SendMessage(icnt)
	if err != nil {
		log.Println("Error on message")
		return
	}
	var result []dbent.Author
	json.Unmarshal(res, &result)
	c.IndentedJSON(http.StatusOK, result)
	return
}
