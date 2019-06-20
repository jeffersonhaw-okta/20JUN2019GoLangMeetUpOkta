package main

import (
	"fmt"
	"net/http"
    "github.com/gorilla/sessions"
)

var sessionStore = sessions.NewCookieStore([]byte("session-store"))
var state = "ApplicationState"
var nonce = "NonceNotSetYet"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	//HELLO
	http.HandleFunc("/Hello", func(w http.ResponseWriter, r *http.Request) {
	
	session, err := sessionStore.Get(r, "session-store")

	if err != nil || session.Values["id_token"] == nil || session.Values["id_token"] == "" {
		fmt.Fprintf(w, "Sorry!")
	}
	if session.Values["id_token"] == "Logged In" {
		fmt.Fprintf(w, "Hello World!")
	}
	})

	//LOGIN
	http.HandleFunc("/Login", func(w http.ResponseWriter, r *http.Request) {
	nonce = "Oktago1234"
	state = "abc1234"
	var redirectPath string

	q := r.URL.Query()
	q.Add("client_id", "<client_id>")
	q.Add("response_type", "token")
	q.Add("response_mode", "fragment")
	q.Add("scope", "openid profile email")
	q.Add("redirect_uri", "http://localhost:80/implicit/callback")
	q.Add("state", state)
	q.Add("nonce", nonce)

	redirectPath = "https://<domain>.oktapreview.com/oauth2/<auth_server_id>/v1/authorize?client_id=<client_id>&nonce=Oktago1234&redirect_uri=http%3A%2F%2Flocalhost%3A80%2Fimplicit%2Fcallback&response_mode=fragment&response_type=token&scope=openid+profile+email&state=abc1234"
	
	fmt.Println("redirectPath: "+ redirectPath)

	http.Redirect(w, r, redirectPath, http.StatusMovedPermanently)
	})

	//CALLBACK
	http.HandleFunc("/implicit/callback", func(w http.ResponseWriter, r *http.Request) {
		session, err := sessionStore.Get(r, "session-store")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
		session.Values["id_token"] = "Logged In"
		session.Save(r, w)
		http.Redirect(w, r, "/Hello", http.StatusMovedPermanently)
	})

	http.ListenAndServe(":80", nil)
}