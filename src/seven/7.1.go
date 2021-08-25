package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type RecurlyServers struct {
	XMLName		xml.Name	`xml:"servers"`
	Version		string		`xml:"version,attr"`
	Svs			[]server	`xml:"server"`
	Description	string		`xml:",innerxml"`
}

type server struct {
	XMLName		xml.Name	`xml:"server"`
	ServerName	string		`xml:"serverName"`
	ServerIP	string		`xml:"serverIp"`
}


type Servers struct {
	XMLName		xml.Name		`xml:"servers"`
	Version		string			`xml:"version,attr"`
	Svs			[]NewServer		`xml:"server"`
}

type NewServer struct {
	ServerName	string		`xml:"serverName"`
	ServerIP	string		`xml:"serverIP"`
}


func readXML() {
	file, err := os.Open("servers.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	v := RecurlyServers{}

	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(v)
}


func writeXML() {
	s := &Servers{Version: "1"}
	s.Svs = append(s.Svs, NewServer{"TianJinVPN", "192.168.0.1"})
	s.Svs = append(s.Svs, NewServer{"ShenZhenVPN", "192.168.0.2"})
	outPut, err := xml.MarshalIndent(s, " ", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(outPut)
}

func main() {

	readXML()
	writeXML()
}


