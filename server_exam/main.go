package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (b *Book) addBook(name string, price int) {
	b.name = name
	b.price = price
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("id : %v\n", vars["id"])

	book := new(Book)
	book.addBook("abc", 123)

	fmt.Println(book)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login/{id:[0-9]+}", LoginHandler).Methods("GET")
	srv := new(http.Server)
	srv.Handler = r
	srv.Addr = "127.0.0.1:8000"
	// srv := &http.Server{
	// 	Handler: r,
	// 	Addr:    "127.0.0.1:8000",
	// }

	srv.ListenAndServe()
}
