package routes

import(
	"github.com/gorilla/mux"
	"github.com/dzakideveloper/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(routes *mux.Router){
	routes.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	routes.HandleFunc("/book/", controllers.GetAllBooks).Methods("GET")
	routes.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	routes.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	routes.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}