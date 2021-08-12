package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	s "session"

	_ "memory"
)

var globalSession *s.Manager

func init() {
	globalSession, _ := s.NewManager("memory", "go-session-id", 3600)
	go globalSession.GC()

}


func Login(w http.ResponseWriter, r *http.Request) {
	session := globalSession.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, session.Get("username"))
	} else {
		session.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
}

func Count(w http.ResponseWriter, r *http.Request) {
	sess := globalSession.SessionStart(w, r)
	createTime := sess.Get("create-time")
	if createTime == nil {
		sess.Set("create-time", time.Now().Unix())
	} else if (createTime.(int64) + 3600) < (time.Now().Unix()) {
		globalSession.SessionDestroy(w, r)
		sess = globalSession.SessionStart(w, r)
	}

	ct := sess.Get("count-num")
	if ct == nil {
		sess.Set("count-num", 1)
	} else {
		sess.Set("count-num", ct.(int) + 1)
	}

	t, _ := template.ParseFiles("count.gtpl")
	w.Header().Set("Content-type", "text/html")
	t.Execute(w,sess.Get("count-num"))
}


func main() {
	http.HandleFunc("/login", Login)
	http.HandleFunc("/count", Count)

	err := http.ListenAndServe(":9099", nil)
	if err != nil {
		log.Fatal("Listen and server ", err)
	}
}
