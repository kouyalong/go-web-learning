package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"
)

func funcTime() {
	t := time.Now()
	for i :=0 ; i < 10000 ; i++ {
		for j := 0; j < 100; j++{

		}
	}
	t1 := time.Now()

	dt := t1.Nanosecond() - t.Nanosecond()
	fmt.Println("外大内小", dt)


	t2 := time.Now()
	for i :=0 ; i < 100; i++ {
		for j := 0; j < 10000; j++{

		}
	}
	t3 := time.Now()

	dt1 := t3.Nanosecond() - t2.Nanosecond()
	fmt.Println("外小内大", dt1)
}


type Friend struct {
	Fname	string
}

type Person struct {
	UserName	string
	Emails		[]string
	Friends		[]*Friend
}

func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1{
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}

	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}

	return substrs[0] + " at " + substrs[1]
}


func main() {
	funcTime()

	f1 := Friend{Fname: "minux,ma"}
	f2 := Friend{Fname: "xushiwei"}

	t := template.New("fieldName Example")
	t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	t, err := t.Parse(`
			hello {{.UserName}}!
				{{range .Emails}}
					an emails {{.|emailDeal}}
				{{end}}
				{{with .Friends}}
				{{range .}}
					my friend name is {{.Fname}}
				{{end}}
				{{end}}
			`)

	if err != nil {
		fmt.Println(err)
	}

	p := Person{UserName: "Astaxie",
				Emails: []string{"astaxie@beego.me", "astaxie@gmail.com"},
				Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}
