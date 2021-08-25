package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bitly/go-simplejson"
)

type Server struct {
	ServerName	string
	ServerIP	string
}

type ServersLice struct {
	Servers	[]Server
}


type ServerNew struct {
	ID				int		`json:"-"`
	ServerName		string	`json:"serverName"`
	ServerName2		string	`json:"serverName2,string"`

	ServerIP		string	`json:"serverIP,omitempty"`
}


func main()  {
	var s ServersLice
	str := `{"servers": [{"serverName": "ShangHaiVpn", "serverIP": "127.0.0.1"},{"serverName": "BeijingVpn", "serverIP": "127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)

	js, _ := simplejson.NewJson([]byte(`{
		"test": {
			"array": [1, "2", 3],
			"int": 10,
			"float": 5.1001,
			"bigNum": 192812991279817289719279,
			"string": "simpleJson",
			"bool": true
		}
	}`))

	arr, _ := js.Get("test").Get("array").Array()
	fmt.Println(arr)
	for k, v := range arr {
		fmt.Println(k, v)
	}

	i, _ := js.Get("test").Get("int").Int()
	fmt.Println(i)

	ms := js.Get("test").Get("string").MustString()
	fmt.Println(ms)


	var newS ServersLice
	newS.Servers = append(newS.Servers, Server{ServerName: "JiangSuVPN", ServerIP: "192.168.20.11"})
	newS.Servers = append(newS.Servers, Server{ServerName: "ShenZhenVPN", ServerIP: "172.10.12.11"})
	b, err := json.Marshal(newS)
	if err != nil {
		fmt.Println("json error:", err)
	}
	fmt.Println(string(b))

	newSs := ServerNew {
		ID: 3,
		ServerName: `GO "1.0"`,
		ServerName2: `GO "1.0"`,
		ServerIP: ``,
	}
	bNew, _ := json.Marshal(newSs)
	os.Stdout.Write(bNew)

}
