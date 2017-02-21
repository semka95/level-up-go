package main

import (
	"fmt"
	"io"
	"os"

	"github.com/russross/blackfriday"
)

// custom types

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
	PB = 1024 * TB
)

// ByteSize aaa
type ByteSize float64

func (b ByteSize) String() string {
	switch {
	case b >= PB:
		return "Very Big"
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%fB", b)
}

// interfaces

// Fruit aaa
type Fruit interface {
	String() string
}

// Apple aaa
type Apple struct {
	Variety string
}

func (a Apple) String() string {
	return fmt.Sprintf("A %s apple.", a.Variety)
}

// Orange aaa
type Orange struct {
	Size string
}

func (o Orange) String() string {
	return fmt.Sprintf("A %s orange.", o.Size)
}

// PrintFruit aaa
func PrintFruit(fruit Fruit) {
	fmt.Println("I have this fruit:", fruit.String())
}

// error handling

// ErrInvalidStatusCode aaa
type ErrInvalidStatusCode int

func (code ErrInvalidStatusCode) Error() string {
	return fmt.Sprintf("Expected code 200, but got code %d", code)
}

// embedded types

// User aaa
type User struct {
	Name string
}

// IsAdmin aaa
func (u User) IsAdmin() bool { return false }

// DisplayName aaa
func (u User) DisplayName() string {
	return u.Name
}

// Admin aaa
type Admin struct {
	User
}

// IsAdmin aaa
func (a Admin) IsAdmin() bool { return true }

// DisplayName aaa
func (a Admin) DisplayName() string {
	return "[Admin] " + a.User.DisplayName()
}

// defer
func copyFile(dstName, srcName string) error {
	src, err := os.Open(srcName)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Open(dstName)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}

// Human aaa
type Human struct {
	Name string
	Age  int
}

// JoinTwoStrings trailing commas
func JoinTwoStrings(
	stringOne,
	stringTwo string,
) string {
	return stringOne + stringTwo
}

func main() {
	// custom types
	fmt.Println(ByteSize(2048))
	fmt.Println(ByteSize(3292528.64))

	// interfaces
	apple := Apple{"Golden Delicious"}
	orange := Orange{"large"}
	PrintFruit(apple)
	PrintFruit(orange)

	//error handling
	statusCode := 404
	if statusCode != 200 {
		fmt.Println(ErrInvalidStatusCode(statusCode).Error())
	}

	// embedded types
	u := User{"Normal User"}
	fmt.Println(u.Name)
	fmt.Println(u.DisplayName())
	fmt.Println(u.IsAdmin())

	a := Admin{User{"Admin User"}}
	fmt.Println(a.Name)
	fmt.Println(a.User.Name)
	fmt.Println(a.DisplayName())
	fmt.Println(a.IsAdmin())

	// libraries
	markdown := []byte(`
					# This is a header
					* and
					* this
					* is
					* a
					* list
 					`)

	html := blackfriday.MarkdownBasic(markdown)
	fmt.Println(string(html))

	// struct initialization
	me := Human{"Semyon", 21}
	fmt.Println(me)
	me = Human{Name: "Semyon", Age: 21}

	// empty variable initialization
	var myString string
	myString = "Hello"
	fmt.Println(myString)
	var myMap map[string]string
	myMap = map[string]string{}
	myMap["Test"] = "Hi"
	fmt.Println(myMap)

	// trailing commas
	mySlice := []string{"one", "two"}
	mySlice = []string{
		"one",
		"two",
	}
	fmt.Println(mySlice)

	myMap2 := map[string]int{"one": 1, "two": 2}
	myMap2 = map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println(myMap2)

	myString1 := JoinTwoStrings(
		"Hello",
		"World",
	)
	fmt.Println(myString1)

}
