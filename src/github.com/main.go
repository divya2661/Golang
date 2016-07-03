package main

import(
	"fmt"
	"gorilla/mux"
	"gorilla/securecookie"
	"net/http"
)

var cookieHanler = securecookie.new(
					securecookie.GenerateRandomKey(64),
					securecookie.GenerateRandomKey(32)
						)

var router = mux.NewRouter()

func func main() {
	
	router.HandleFunc("/", indexPageHandler);
	router.HandleFunc("/internal",internalPageHandler)

	router.HandleFunc("/login",loginHandler).Methods("POST")
	router.HandleFunc("/logout",logoutHandler).Methods("POST")

	http.Handle("/",router)
	http.ListenAndServe(":8000",nil)
}
