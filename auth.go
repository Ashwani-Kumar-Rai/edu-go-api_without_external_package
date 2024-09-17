// Now allowing anyone who calls the endpoint to get an update any client information doesn't make much sense
// So adding authentication to our endpoints . the best way to do is by middleware .
// Middleware is similar to wrappers in other languages ,
// it's a function that which takes in a function and returns another function .
// Middleware(fn function ) -> function 


package main

import (
	"net/http"
	"strings"
)

// In the case of middleware for our API it will take in and return a particular type of function 
// which is the HandlerFunc type , we saw this type before with our get client profile and our update client profile 
// so our token auth middleware will take in a Handler Func and return a Handler Func .

// First step is validate the clientId as we did before in our other handlers . 
// the auth method we will use here is a bearer token , this is a token which is passed in through the header under 
// the authorization key , we'll compare this token to what we have in our database for the client so 
// let's make a helper function , which checks if the token starts with a bearer prefix and then checks 
// that it matches what we have in our database if it doesn't we again throw a forbidden status 
// if all is good we'll call the next handler in line , in our case we don't have any other middleware 
// so this will just be the handleClientProfile handler now to apply this middleware  in main func
// wrap our Handler client profile with the middleware function .  
func TokenAuthMiddleware(next http.HandlerFunc ) http.HandlerFunc {
	return func(w http.ResponseWriter , r *http.Request) {
	var clientId = r.URL.Query().Get("clientId") // will get the clientId
	clientProfile, ok := database[clientId]      // we'll use clientId to lookup client profile in the database .
	if !ok || clientId == "" {                   // if clientID does not exist or it was not passed at all , we'll return the forbidden message and the forbidden status error code .
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	token := r.Header.Get("Authorization")
	if !isValidToken(clientProfile, token ){
		 http.Error(w, "Forbidden", http.StatusForbidden)
		 return
	}
	next.ServeHTTP(w,r)
  }
}

func isValidToken(clientProfile ClientProfile, token string) bool {
	if strings.HasPrefix(token, "Bearer"){
		return strings.TrimPrefix(token,"Bearer ") == clientProfile.Token 
	}
	return false
}