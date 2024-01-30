package main

import (
	"fmt"
	"golang-web/src/config"
	"golang-web/src/helper"
	"golang-web/src/routes"
	"net/http"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	config.InitDB()
	helper.Migration()
	defer config.DB.Close()
	routes.Route()
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
