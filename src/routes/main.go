package routes

import (
	"golang-web/src/controllers"
	"golang-web/src/middleware"
	"net/http"

	"github.com/goddtriffin/helmet"
)

func Route() {
	helmet := helmet.Default()

	// http.HandleFunc("/products", helmet.Secure(http.HandlerFunc(controllers.ProductsController)))
	http.Handle("/products", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(controllers.ProductsController))))
	// http.Handle("/products", middleware.JwtMiddleware(http.HandlerFunc(controllers.ProductsController)))
	http.HandleFunc("/product/", controllers.ProductController)
	http.HandleFunc("/register", controllers.RegisterUser)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/upload", controllers.HandleUpload)

}
