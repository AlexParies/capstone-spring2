package main
import(
"net/http"
"fmt"
"github.com/gin-gonic/gin")

type spotVarApi struct {
	id     int `json:"id"`
	status bool `json:"status"`
	x      int `json:"x"`
	y      int `json:"y"`
}

var sp []spotVarApi

func updateApi(spots []spotVar) {
	sp = spots
	spotCount := len(spots)
	for  i := 0; i < spotCount; i++{
		sp(i).id = spots(i).id
		sp(i).status = spots(i).status
		sp(i).x = spots(i).x
		sp(i).y = spots(i).y
	}
	
	
	fmt.Println(sp)
}