package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andremelinski/web-dev-todd/servers/14-sessions/db"
	"github.com/julienschmidt/httprouter"
)

func HandleRequests() {
	mux := httprouter.New()
	mux.POST("/", db.CreateUser)
	mux.GET("/", db.SignUp)
	mux.GET("/login", login)
	mux.GET("/logout", db.Logout)
	log.Fatal(http.ListenAndServe(":8080", mux))
}


func login(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	_, err := req.Cookie("session-id")
	if err == http.ErrNoCookie {
		//  http.Redirect(res, req, "/logout", 300)
		json.NewEncoder(res).Encode("logged out")
		return
	}
	json.NewEncoder(res).Encode("Hello you are logged in")
}
