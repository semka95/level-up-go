package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	//strings
	str := "it's a string"
	fmt.Println(str)

	anotherstr := str + " and anotha one"
	fmt.Println(anotherstr)

	//numbers
	myInt := 4
	myFloat := 1.5
	conv := float64(myInt) * myFloat
	fmt.Println(conv)

	//bool
	myBool := true
	if myBool {
		fmt.Println("it's true")
	} else {
		fmt.Println("it's false")
	}

	// arrays
	myArr := []int{1, 2, 3, 4, 5}
	myArr1 := myArr[0:3]
	myArr2 := myArr[1:4]
	fmt.Println(myArr1, myArr2, myArr2[2])
	fmt.Println(len(myArr))

	// loops
	animals := []string{"cat", "dog", "fish"}
	for i, animal := range animals {
		fmt.Println(animal, "is at index", i)
	}
	for _, animal := range animals {
		fmt.Println(animal)
	}
}
