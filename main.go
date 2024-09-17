package main

// import net/http package that will power the API .
import (
	"log"
	"net/http"
)

// Now to apply this middleware  in main func  wrap our Handler client profile with the middleware function .
// create a middleware type , which again is just a function which takes in a Handler Func and returns
// another Handler Func . and we can define a list of all of our middlewares .
// which for us is just the token off middleware function .
type Middleware func(http.HandlerFunc) http.HandlerFunc

var middlewares = []Middleware{
	TokenAuthMiddleware,
}

// we use http.HandleFunc function to define an end point .
// 1st input is the path : we are building an api to manage user profile .
// 2nd input is the function that will actually handle the request : called handler function .
func main() {
	// define our dummy handler variable which just points to our handle client profile function .
	// then continously wrap this this function with list of middleware .
	// lastly this Handler will now be passed into the handlefunc
	var handler http.HandlerFunc = handleClientProfile
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	http.HandleFunc("/user/profile", handler)
	// logging some information .
	log.Println("Server is running one prot 8080 ...")
	// start the server on port 8080 with http.ListenAndServe function
	log.Fatal(http.ListenAndServe(":8080", nil))
}
