package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

//func Match(pattern string, b []byte) (matched bool, error error)
//func MatchReader(pattern string, r io.RuneReader)(matched bool, error error)
//func MatchString(pattern string, s string)(matched bool, error error)


func IsIP(ip string)(m bool){
	m, _ = regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip)
	return
}

func IsNum(s string)(m bool){
	if m, _ := regexp.MatchString("^[0-9]+$", s); m {
		fmt.Println("数字")
	} else {
		fmt.Println("不是数字")
	}
	return
}


func Download(url string){
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http get error")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
		return
	}

	src := string(body)

	re, _ := regexp.Compile(`<[\S\s]+?>`)
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	re, _ = regexp.Compile(`<style[\S\s]+?</style>`)
	src = re.ReplaceAllString(src, "")

	re, _ = regexp.Compile(`&lt;style[\S\s]+?&lt;/style&gt;`)
	src = re.ReplaceAllString(src, "")

	re, _ = regexp.Compile(`<script[\S\s]+?</script>`)
	src = re.ReplaceAllString(src, "")

	re, _ = regexp.Compile(`&lt;script[\S\s]+?&lt;/script&gt;`)
	src = re.ReplaceAllString(src, "")

	re, _ = regexp.Compile(`<[\S\s]+?>`)
	src = re.ReplaceAllString(src, "\n")

	re, _ = regexp.Compile(`\s{2,}`)
	src = re.ReplaceAllString(src, "\n")


	fmt.Println(strings.TrimSpace(src))
}

func main(){
	f := IsIP("192.168.0.1")
	fmt.Println(f)

	IsNum("100w")

	url := "http://www.baidu.com"
	Download(url)

	src := []byte(`
			call hello alice
			hello bob
			call hello eve
	`)

	pat := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s(?P<arg>.+)\s*$`)
	res := []byte{}
	for _, s := range pat.FindAllSubmatchIndex(src, -1) {
		res = pat.Expand(res, []byte("$cmd('$arg')\n"), src, s)
	}
	fmt.Println(string(res))

}