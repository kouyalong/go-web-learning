package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}


type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width*r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type Color byte

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Box struct {
	width, height, depth float64
	color Color
}

type BoxList []Box

func (b Box) Volume() float64 {
	return b.width*b.height*b.depth
}

func (b *Box) SetColor(c Color) {
	b.color = c
}


func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack() {
	for i := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

type HumanNew struct {
	name string
	age int
	weight int
	phone string
}

type StudentNew struct {
	HumanNew
	school string
}

type EmployeeNew struct {
	HumanNew
	company string
}

func (h *HumanNew) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s \n", h.name, h.phone)
}

func (e *EmployeeNew) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s, Call me on %s \n", e.name, e.company, e.phone)
}

func main() {
	r1 := Rectangle{12.5, 8.9}
	r2 := Rectangle{8, 9}
	c1 := Circle{10}
	c2 := Circle{9.7}

	fmt.Println("r1 area's", r1.area())
	fmt.Println("r2 area's", r2.area())
	fmt.Println("c1 area's", c1.area())
	fmt.Println("c2 area's", c2.area())

	boxs := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("we have %d boxs in out set\n", len(boxs))
	fmt.Println("The Volume of the first one is", boxs[0].Volume(), "cm3")
	fmt.Println("The color of the last one is", boxs[len(boxs)-1].color.String())
	fmt.Println("The biggest one is", boxs.BiggestColor().String())
	fmt.Println("Let's  paint them all black")
	boxs.PaintItBlack()
	fmt.Println("The color if the second is", boxs[1].color.String())
	fmt.Println("Obviously now, the biggest one is ", boxs.BiggestColor().String())

	student := StudentNew{HumanNew{"Selina", 18, 88, "010-1234"}, "Beijing University"}
	employee := EmployeeNew{HumanNew{"Peter", 44, 145, "010-4321"}, "MeiDI"}
	student.SayHi()
	employee.SayHi()

}