package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Human2 struct {
	name string
	age int
	phone string
}

type Student2 struct {
	Human2
	school string
	loan float32
}

type Employee2 struct {
	Human2
	company string
	money float32
}

func (h Human2) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s \n", h.name, h.phone)
}

func (h Human2) Sing(lyrics string) {
	fmt.Println("La La La...", lyrics)
}

func (h Human2) String() string {
	return "<" + h.name + "-" + strconv.Itoa(h.age) + " years - 0" + h.phone + ">"
}


func (e Employee2) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s, Call me on %s \n", e.name, e.company, e.phone)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

type Element interface {}
type List []Element
type Person struct {
	name string
	age int
}

func (p Person) String() string {
	return "name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years"
}


func main() {
	mike := Student2{Human2{"Mike", 25, "222-222-222"}, "MIT", 0.1}
	paul := Student2{Human2{"Paul", 26, "333-333-333"}, "Harvard", 100}
	sam := Employee2{Human2{"Sam", 36, "444-444-444"}, "Google Inc.", 10000}
	tom := Employee2{Human2{"TOM", 37, "555-555-555"}, "Things Ltd.", 50000}

	var i Men
	i = mike
	fmt.Println("This is Mike, a student:")
	i.SayHi()
	i.Sing("November rain")

	i = tom
	fmt.Println("This is top, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, mike
	for _, value := range x{
		value.SayHi()
	}

	bob := Human2{"Bob", 39, "000-777-888"}
	fmt.Println("This Human is ", bob)

	list := make(List, 4)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"Dennis", 60}
	list[3] = 3.000

	for index, element := range list {
		if value, ok := element.(int); ok{
			fmt.Printf("List[%d] is an int and it's value is %d \n", index, value)
		} else if value, ok := element.(string); ok{
			fmt.Printf("List [%d] is a string and is's value is %s \n", index, value)
		} else if value, ok := element.(Person); ok{
			fmt.Printf("List [%d] is a person and its value is %s \n", index, value)
		} else {
			fmt.Printf("List [%d] is of a different type \n", index)
		}
	}

	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("List[%d] is an int and it's value is %d \n", index, value)
		case string:
			fmt.Printf("list [%d] is a string and it's value is %s \n", index, value)
		case Person:
			fmt.Printf("list [%d] is a Person and it's value is %s \n", index, value)
		default:
			fmt.Printf("list [%d] is of a different type \n", index)
		}
	}

	var cNumber float64 = 3.1415926
	p := reflect.ValueOf(&cNumber)
	//p.SetFloat(12)
	v:= p.Elem()
	v.SetFloat(7.1234)
	fmt.Println("C number is ", cNumber, "V number is ", v)
}
