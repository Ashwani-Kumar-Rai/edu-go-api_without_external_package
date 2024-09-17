// fake in-memory database

package main

type ClientProfile struct {
	Email string
	Id    string
	Name  string
	Token string
}

// this stores a map of client id to the profile data for two customers .
// the API we create will be able to authenticate client requests and retrieve and update customer profile data .
var database = map[string]ClientProfile{
	"user1": {
		Email: "ram@gmail.com",
		Id:    "user1",
		Name:  "Ram lal",
		Token: "123",
	},
	"user2": {
		Email: "shyam@gmail.com",
		Id:    "user2",
		Name:  "Shyam lal",
		Token: "456",
	},
}
