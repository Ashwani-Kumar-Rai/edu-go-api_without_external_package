package main

// import net/http package that will power the API .
import (
	"log"
	"net/http"
)

// we use http.HandleFunc function to define an end point .
// 1st input is the path : we are building an api to manage user profile .
// 2nd input is the function that will actually handle the request : called handler function .
func main() {
	http.HandleFunc("/user/profile", handleClientProfile)
	// logging some information .
	log.Println("Server is running one prot 8080 ...")
	// start the server on port 8080 with http.ListenAndServe function
	log.Fatal(http.ListenAndServe(":8080", nil))
}
