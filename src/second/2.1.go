package main

import "fmt"

const (
	num1	= iota
	num2
	num3
	num4, num5, num6	= iota, iota, iota
	num7	= iota
	num8	= iota
)


func computeValue() int {
	return 3
}

func SumAndProduct(a, b int) (int, int){
	return a+b, a*b
}

func updateMap(key string, value int, pMap map[string]int) {
	pMap[key] = value
}


type testInt func(int) bool

func isOdd(num int) bool {
	if num % 2 == 0{
		return false
	}
	return true
}

func isEven(num int) bool {
	if num % 2 == 0{
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice{
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}


type person struct {
	name string
	age int
}

func olderOne(p1, p2 person) (person, int) {
	if p1.age>p2.age{
		return p1, p1.age-p2.age
	}
	return p2, p2.age-p1.age
}

type Human struct {
	name string
	age int
	weight int
	phone string
}

type Student struct {
	Human
	speciality string
}

type Employee struct {
	Human
	int
	phone string
}

func main() {
	const Pi float32 = 3.1415926
	var vName1 string
	var varInt32 int32
	var isActive bool
	isActive = true
	_, b := 34, 35
	var frenchHello string
	var emptyString string = ""

	frenchHello = "Bonjour"

	hello := "hello"
	sliceHello := []byte(hello)
	sliceHello[0] = 'C'
	cello := string(sliceHello)

	moreLineString := `这是一个多行的字符串
							你看出来了么`

	var fslice []int
	newArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var bSlice []int
	bSlice = newArr[2:8]
	fslice = newArr[3:]

	fmt.Printf("hello, world or 你好 \n")
	fmt.Printf("%s and %d \n", vName1, varInt32)
	fmt.Printf("%d \n", b)
	fmt.Printf("圆周率%f \n", Pi)
	fmt.Printf("布尔值%t \n", isActive)
	fmt.Printf("French %s, empty %s \n", frenchHello, emptyString)
	fmt.Printf("%s, %s, %s, %s \n", hello, sliceHello, cello, hello + cello)
	fmt.Printf("%s \n", moreLineString)
	fmt.Printf("%d, %d, %d, %d, %d, %d, %d, %d \n", num1, num2, num3, num4, num5, num6, num7, num8)

	cSlice := make([]int, 10)
	count := copy(cSlice, newArr)
	newArr[2] = 12
	newArr[3] = 13
	fmt.Printf("copy 个数 %d \n", count)
	fmt.Printf("%v, %v, %v, %v, \n", bSlice, fslice, cSlice, newArr)
	fmt.Printf("%d, %d, \n", cap(cSlice), len(cSlice))
	fmt.Printf("%d, %d, \n", cap(newArr), len(newArr))

	var vaSlice = []int{1, 2, 3}
	vaSlice = append([]int{0}, vaSlice...)
	vaSlice = append([]int{-3, -2, -1}, vaSlice...)
	vaSlice = append(vaSlice, 4)
	vaSlice = append(vaSlice[:8], append([]int{5}, vaSlice[8:]...)...)
	vaSlice = append(vaSlice, []int{6, 10, 11}...)
	vaSlice = append(vaSlice[:10], append([]int{7, 8, 9}, vaSlice[10:]...)...)
	fmt.Printf("%v \n", vaSlice)

	var numbers1 map[string]int
	numbers2 := make(map[string]int)

	numbers1 = make(map[string]int)
	numbers1["one"] = 1
	numbers1["two"] = 2
	numbers1["three"] = 3

	numbers2["one"] = 1
	numbers2["two"] = 2
	numbers2["three"] = 3
	fmt.Printf("%v, %v \n", numbers1, numbers2)

	four, ok := numbers1["four"]
	if ok {
		fmt.Printf("numbers1 has %d \n", four)
	}else{
		fmt.Printf("numbers1 not have %s \n", "four")
	}

	if x := computeValue(); x == 3{
		fmt.Printf("compute value is 3 \n")
	}else{
		fmt.Printf("computer value is not 3 \n")
	}

	xValue := computeValue()
	if xValue == 3{
		fmt.Printf("xValue is 3 \n")
	}else if xValue < 3{
		fmt.Printf("xValue less than 3\n")
	}else{
		fmt.Printf("xValue is greater than 3\n")
	}

	for k, v := range numbers1{
		fmt.Printf("key is %s, value is %d \n", k, v)
	}

	nValue, _ := SumAndProduct(1, 5)
	switch nValue {
	case 1:
		fmt.Printf("nValue is 1\n")
	case 2:
		fmt.Printf("nValue is 2\n")
	case 3:
		fmt.Printf("nValue is 3\n")
	default:
		fmt.Printf("nValue is not known, Last Print nValue is %d \n", nValue)
	}

	updateMap("one", 10, numbers1)
	fmt.Printf("After update map numbers1 %v \n", numbers1)

	for i := 0; i < 5; i ++{
		defer fmt.Printf("%d \n", i)
	}

	slice := []int{0, 1, 2, 3, 4, 5}
	odd := filter(slice, isOdd)
	fmt.Printf("odd is %v \n", odd)
	even := filter(slice, isEven)
	fmt.Printf("even is %v \n", even)

	var personOne person
	personOne.name = "Bob"
	personOne.age = 32
	fmt.Printf("This is %v \n", personOne)

	personTwo := person{"Jone", 25}
	fmt.Printf("This is %v \n", personTwo)

	mark := Student{Human{"Mark", 25, 110, "010-333111"}, "Computer Science"}
	fmt.Printf("This is %s's info  %v \n", mark.name, mark)

	employee := Employee{Human{"Jeff", 54, 120, "xxx111"}, 10, "010-1212"}
	fmt.Println("This employee's phone is", employee.phone)
	fmt.Println("This employee's work phone is", employee.Human.phone)

}
