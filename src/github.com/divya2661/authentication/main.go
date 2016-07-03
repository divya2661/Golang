package main

import(
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"net/http"
)

var cookieHanler = securecookie.New(
					securecookie.GenerateRandomKey(64),
					securecookie.GenerateRandomKey(32))

func indexPageHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, indexPage)
}

const indexPage =  '<form method="post" action="/login">
	    				<label for="name">User name</label>
	    				<input type="text" id="name" name="name">
	    				<label for="password">Password</label>
	    				<input type="password" id="password" name="password">
	    				<button type="submit">Login</button>
					</form>'



var router = mux.NewRouter()

func main() {
	
	fmt.Printf("hello, divya\n")
	
	router.HandleFunc("/", indexPageHandler)
	router.HandleFunc("/internal",internalPageHandler)

	router.HandleFunc("/login",loginHandler).Methods("POST")
	router.HandleFunc("/logout",logoutHandler).Methods("POST")

	http.Handle("/",router)
	http.ListenAndServe(":8000",nil)
}
