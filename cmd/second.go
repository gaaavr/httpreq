package main

import (
	"httpreq/pkg/init_DB"
	"httpreq/pkg/user"
	"net/http"
)

func main() {
	init_DB.InitDB()
	mux := http.NewServeMux()
	mux.HandleFunc("/create", user.CreateUser)
	mux.HandleFunc("/make_friends", user.MakeFriends)
	mux.HandleFunc("/user", user.DeleteUser)
	mux.HandleFunc("/friends/", user.GetFriends)
	mux.HandleFunc("/", user.AgeUpdate)
	http.ListenAndServe("localhost:8081", mux)
}
