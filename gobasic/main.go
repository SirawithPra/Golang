package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Printf("SAVADDEE %d \n", 20000)
	fmt.Printf("%v\n", true)
	fmt.Printf("test %v\n", "two")
	println("Hello World")
	println("Test")
	fmt.Printf("test %#v\n", false)
	sp()

	//การประกาศตัวแปล

	var x1 int = 10 //ประกาศพร้อมค่า
	var x2 = 11     //ประกาศพร้อมค่า ละ type
	x3 := 12        //short decarlation แต่สร้างนอก main ไม่ได้
	var y int       //full
	y = 10
	var z int

	fmt.Printf("test x1 %v ; x1=10\n", y*x1)
	fmt.Printf("test x2 %v ; x2=11\n", y*x2)
	fmt.Printf("test x3 %v ; x3=12\n", y*x3)
	fmt.Printf("test z %v ; z=0 def value\n", y*z)
	sp()

	//array
	var a [3]int = [3]int{4, 5, 6}
	a1 := [3]int{1, 2, 3}
	a2 := [...]int{1, 2, 4, 5, 6, 7, 56}
	a3 := [10]int{12, 3, 3, 34}
	a4 := [2][3]int{{2, 32, 1}, {1, 12}}
	fmt.Println(a)
	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", a1)
	fmt.Printf("%#v\n", a2)
	fmt.Printf("%#v\n", a3)
	fmt.Printf("%#v\n", a4)
	sp()

	//slice or list in python
	x := []int{1, 2, 3} //ไม่จำกัดขนาด

	x = append(x, 4)  //ต้องเอาตัวแปลมารับ
	u := append(x, 5) //เอาตัวแปลใหม่มารับต้องสร้าง short decarlation ใหม่
	fmt.Printf("x = %v range = %v \n", x, len(x))
	fmt.Printf("u = %v range = %v \n", u, len(u))
	u = append(u, 10)
	fmt.Printf("u = %v range = %v \n", u, len(u))
	name_eng := "AIM"
	name_th := "เอม"
	fmt.Printf("len_name_eng = %v\n", len(name_eng))          //นับได้ปกติ เพราะ ตัวภาษาอังกฤษตัวละ byte
	fmt.Printf("len_name_th = %v\n", len(name_th))            //ภาษาอื่นๆ byte ไม่เท่ากัน
	println("len_name_th =", utf8.RuneCountInString(name_th)) //สิ่งที่ควรใช้นับ
	x = []int{10, 20, 30, 40, 50, 60, 70}
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", x[5:7])
	fmt.Printf("%v\n", x[5]*6/4)
	sp()

	// if clause
	point := 50
	if point > 40 {
		println("More than 40")
	}
	printFullName("Sirawith", "Pra") //สร้าง func print name
}

func printFullName(firstName string, lastName string) {
	fmt.Println("Hello" + firstName + " " + lastName)
}
func sp() {
	fmt.Println("       .....      ")
	fmt.Println()
}
