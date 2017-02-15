package main

import "fmt"

//Movie aaa
type Movie struct {
	Actors      []string
	Rating      float32
	ReleaseYear int
	Title       string
}

// Counter aaaa
type Counter struct {
	Count int
}

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

	// maps
	starWarsYears := map[string]int{
		"A New Hope":              1977,
		"The Empire Strikes Back": 1980,
		"Return of the Jedi":      1983,
		"Attack of the Clones":    2002,
		"Revenge of the Sith":     2005,
	}
	starWarsYears["The Force Awakens"] = 2015
	fmt.Println(len(starWarsYears))

	// looping over maps
	for title, year := range starWarsYears {
		fmt.Println(title, "was released in", year)
	}

	// maps
	colours := map[string]string{
		"red":     "#ff0000",
		"green":   "#00ff00",
		"blue":    "#0000ff",
		"fuchsia": "#ff00ff",
	}
	redHex := colours["red"]
	fmt.Println(redHex)
	delete(colours, "fuchsia")
	code, exists := colours["wat?"]
	if exists {
		fmt.Println("exists ", code)
	} else {
		fmt.Println("not exists ", code)
	}

	// functions
	a, b := oneParamTwoReturns(3)
	fmt.Println(a, b)

	// pointers
	fruit := "banana"
	giveMePear(fruit)
	fmt.Println(fruit)

	giveMeFruitP(&fruit)
	fmt.Println(fruit)

	// structs
	episodeTV := Movie{
		Title:       "Star Wars: A New Hope",
		Rating:      5.0,
		ReleaseYear: 1977,
	}

	episodeTV.Actors = []string{
		"Mark Hamill",
		"Harrison Ford",
		"Carrie Fisher",
	}
	fmt.Println(episodeTV.Title, "has a rating of", episodeTV.Rating)

	// type methods
	episodeV := Movie{
		Title:       "Star Wars: The Empire Strikes Back",
		ReleaseYear: 1980,
	}
	fmt.Println(episodeV.DisplayTitle())

	counter := &Counter{}
	fmt.Println(counter.Count)
	counter.Increment()
	fmt.Println(counter.Count)
	counter.IncrementWithPointer()
	fmt.Println(counter.Count)
}

// functions
func noParamsNoReturn() {
	fmt.Println("No params return")
}

func twoParamsOneReturn(myInt int, myString string) string {
	return fmt.Sprintf("myInt: %d, myString: %s", myInt, myString)
}

func oneParamTwoReturns(myInt int) (string, int) {
	return fmt.Sprintf("Int: %d", myInt), myInt + 1
}

func twoSameTypedParams(myStr1, myStr2 string) {
	fmt.Println("String 1: ", myStr1)
	fmt.Println("String 2: ", myStr2)
}

func giveMePear(fruit string) {
	fruit = "pear"
}

func giveMeFruitP(fruit *string) {
	*fruit = "pear"
}

// DisplayTitle aaa
func (movie Movie) DisplayTitle() string {
	return fmt.Sprintf("%s (%d)", movie.Title, movie.ReleaseYear)
}

// Increment aaa
func (c Counter) Increment() {
	c.Count++
}

// IncrementWithPointer aaa
func (c *Counter) IncrementWithPointer() {
	c.Count++
}

// exported and unexported code
func myUnexportedFunc() {

}

// MyExportedFunc aaa
func MyExportedFunc() {

}

// MyExportedType aaa
type MyExportedType struct {
	ExportedField   string
	unexportedField string
}
