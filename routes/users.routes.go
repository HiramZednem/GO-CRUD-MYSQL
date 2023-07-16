package routes

import "net/http"

func GetUsersHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Get Users"))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Get one user"))
}

func PostUserHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("post user"))
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Delete User"))
}

// Para que funciona
// Websockets
// 