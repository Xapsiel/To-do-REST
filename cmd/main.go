package main

import (
	"net/http"

	"test_case/internal/app"
	"test_case/internal/db"

	"github.com/gorilla/mux"
)

func main() {
	run()
	// ys := speller.NewSpeller()
	// s, e := ys.CheckText("mosqow is sity", "en")
	// fmt.Println(s, e)
}
func init() {
	_ = db.Database{}
}
func run() {
	router := mux.NewRouter()
	router.HandleFunc("/add", app.Add)
	router.HandleFunc("/signIn", app.SignIn).Methods("POST")
	router.HandleFunc("/signUp", app.SignUp).Methods("POST")
	router.HandleFunc("/get/{userID}", app.Get).Methods("GET")
	http.ListenAndServe(":8080", router)

}
