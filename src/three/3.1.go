package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal("Request Form parse ", err)
	}
	fmt.Println("Form is ", r.Form)
	fmt.Println("Url Path is ", r.URL.Path)
	fmt.Println("Url Scheme is ", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("Key is ", k)
		fmt.Println("Value is ", strings.Join(v, ""))
	}

	_, err = fmt.Fprintf(w, "Hello World")
	if err != nil {
		log.Fatal("format print ", err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("THis func method ", r.Method)
	if r.Method == "GET" {
		timestamp := strconv.Itoa(time.Now().Nanosecond())
		hashWr := md5.New()
		hashWr.Write([]byte(timestamp))
		token := fmt.Sprintf("%x", hashWr.Sum(nil))
		t, err := template.ParseFiles("login.gtpl")
		if err != nil {
			log.Fatal("Got error ", err)
		}
		log.Println(t.Execute(w, token))
	} else {
		_ = r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {

		} else {
			fmt.Println("非法请求")
		}
		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])
		fmt.Println("token", r.Form["token"])

	}

}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Method is ", r.Method)
	if r.Method == "GET" {
		currentTime := time.Now().Unix()
		h := md5.New()
		_, _ = io.WriteString(h, strconv.FormatInt(currentTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		_ = t.Execute(w, token)
	} else {
		_ = r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()

		_, _ = fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/" + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()
		_, _ = io.Copy(f, file)
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)

	err := http.ListenAndServe(":9099", nil)
	if err != nil {
		log.Fatal("Listen And Serve ", err)
	}
}
