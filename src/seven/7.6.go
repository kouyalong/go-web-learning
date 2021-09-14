package main

import (
	"fmt"
	"strconv"
	"strings"
)


func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}


func sOp() {
	fmt.Println(strings.Contains("seafood", "sea"))
	fmt.Println(strings.Contains("", ""))

	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))

	fmt.Println(strings.Index("chicken","ken"))
	fmt.Println(strings.Index("chicken", "ddd"))

	fmt.Println("ba" + strings.Repeat("na", 2))

	fmt.Println(strings.Replace("Oink Oink", "k", "ky", 2))
	fmt.Println(strings.Replace("ok ok ok", "ok", "Yes", -1))

	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split(" xyz  ", ""))

	fmt.Printf("[%q]\n", strings.Trim("!!!Apache!!!", "!"))
	fmt.Printf("Fields are: %q\n", strings.Fields(" foo bar  baz"))

	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 123, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abc")
	str = strconv.AppendQuoteRune(str, '作')
	fmt.Println(string(str))

	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.333, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)

	pa, err := strconv.ParseBool("false")
	checkError(err)

	pb, err := strconv.ParseFloat("123.345", 64)
	checkError(err)

	pc, err := strconv.ParseInt("32312", 10, 64)
	checkError(err)

	pd, err := strconv.ParseUint("98765", 10, 64)
	checkError(err)

	pe, err := strconv.Atoi("1023")
	checkError(err)

	fmt.Println(pa, pb, pc, pd, pe)
}

func main() {
	sOp()
}
