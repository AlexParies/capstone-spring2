package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type spotVarApi struct {
	id     int  `json:"id"`
	status bool `json:"status"`
	x      int  `json:"x"`
	y      int  `json:"y"`
}

var i = 0
var sp []spotVarApi
var router *gin.Engine

func updateApi(spots []spotVar) {

	spotCount := len(spots)
	for i := 0; i < spotCount; i++ {

		var s spotVarApi

		sp = append(sp, s)

		sp[i].id = spots[i].id
		sp[i].status = spots[i].status
		sp[i].x = spots[i].x
		sp[i].y = spots[i].y
	}

}

func getSpots(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, sp)
}

func callme(wowzers []spotVar) {
	if i == 0 {
		fmt.Print("Api updated")
		fmt.Println(sp)
		router = gin.Default()
		i++
	}
	updateApi(wowzers)

	router.GET("/spots", getSpots)

	router.Run("localhost:8080")

}
