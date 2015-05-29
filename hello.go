package hello

import (
	"html/template"
	"net/http"

	"appengine"
	"appengine/channel"
)

const channelKey = "tno_chat"

func init() {
	http.HandleFunc("/send", chatHandler)
	http.HandleFunc("/", mainHandler)
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)

	username := r.FormValue("username")
	message := r.FormValue("message")

	if err := channel.SendJSON(context, channelKey, []string{username, message}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)

	tmpl, err := template.ParseFiles("static/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token, err := channel.Create(context, channelKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, map[string]string{"token": token})
}
