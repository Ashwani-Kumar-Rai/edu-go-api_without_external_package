// let's define handleClientProfile function .
package main

import (
	"encoding/json"
	"net/http"
)

// A handler function needs to take in a response writer and a pointer to the request
// w http.ResponseWriter : ResponseWriter handles the writing of the response we want to send back to the caller .
// r *http.Request : Request variable contains information about the incomming request .
// things like the method type , payload data , headers and so on .

// we are gonna make this function just be a router to another function depending on the request method .

// we'll create GET & PATCH requests to the /user/profile endpoint
// where the get request will just return the user profile
// and with the patch request we'll allow the profile to be updated .
// to do this we will use the switch statement : checking the method type in the request object and we will
// call the appropriate handler function . otherwise we'll return an HTTP error with method not allowed status .
func handleClientProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetClientProfile(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// we'll create the handlers
// for the GetClientProfile : we'll require the clientId to be passed as the query parameter .
// so the request parameter contains this information .
func GetClientProfile(w http.ResponseWriter, r *http.Request) {
	var clientId = r.URL.Query().Get("clientId") // will get the clientId
	clientProfile, ok := database[clientId]      // we'll use clientId to lookup client profile in the database .
	if !ok || clientId == "" {                   // if clientID does not exist or it was not passed at all , we'll return the forbidden message and the forbidden status error code .
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	// if everything is ok . we will return the client profile excluding the token .
	response := ClientProfile{
		Email: clientProfile.Email,
		Name:  clientProfile.Name,
		Id:    clientProfile.Id,
	}
	// write the data as a json to the response writer .
	json.NewEncoder(w).Encode(response)
}
