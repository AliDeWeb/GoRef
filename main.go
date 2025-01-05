/*
	<-- Go Useful Commands -->

	*** go mod init example.com/repo_name
	// this command is used for creating go module with repo_name as name and example.com as namespace

	*** go mod tidy
	// this command is used for managing app dependencies,
		when we run it, it would scan the whole code to see which dependencies are missed,
		and then it would install them

	*** go build .
	// this command is used for build the app and compile it to an executable(.exe) file

	*** ./build.exe
	// after the building process is over, we can use the command above to run our compiled app

	*** go run .
	// this command is used for run our app without compiling it, it is good for development

	*** go test
	// this command is used for run app tests

	*** go fmt
	// this command is used for format the source code
*/
/*
	<-- Go Data Types -->

	*** Numbers
		// int:
			Depends on platform:
			32 bits in 32 bit systems and
			64 bit in 64 bit systems
			+ range:
				-2147483648 to 2147483647 in 32 bit systems and
				-9223372036854775808 to 9223372036854775807 in 64 bit systems

		// int8:
			8 bits/1 byte
			+ range:
				-128 to 127

		// int16:
			16 bits/2 byte
			+ range:
				-32768 to 32767

		// int32:
			32 bits/4 byte
			+ range:
				-2147483648 to 2147483647

		// int64:
			64 bits/8 byte
			+ range:
				-9223372036854775808 to 9223372036854775807

		+*+ Unsigned Integers: these are numbers too, but they must be positive. the range is (int * 2) + 1

	*** Floats
		// float32:
			32 bits
			+ range:
				-3.4e+38 to 3.4e+38.

		// float64:
			64 bits
			+ range:
				-1.7e+308 to +1.7e+308.

	*** Boolean
		// bool:
			true | false

	*** String
		// string:
			text => must be in double quotes

	*** Rune
		// rune:
			it is a char like char in java => must be in single quote
*/
/*
	<-- Go Struct -->

	*+* we use struct to create objects. it is like a map and structure for our object.

	// Example:
		type struct_name struct {
		  member1 datatype;
		  member2 datatype;
		  member3 datatype;
		}
*/

// The start point package name must be main
package main

/*
	<-- Importing packages -->
	//+ if we import like, we can directly access of the internal functions of the package without mention its name
		import (
			. ***
		)


	import ***
	or
	import (
		***
		***
	)
	or
	import (
		***
		anotherNameForThisPackage ***
	)
*/
import (
	"fmt"
)

// we can define alias name for types
type UserId = int

// The start point function name must be main
func main() {
	//	declaring vars
	//* var -name -datatype = -data
	//* or
	//* -name := -data
	//+ if variable name is _, go will ignore it.
	//+ in go we cannot declare unused variables.
	//+ in second way of declaring variables, go will automatically guess the type of var.
	var _ int8 = 8 // we cannot define another var with this name
	age := 8       // we can define another var with this name
	//+ also we can do this
	name, age, address := "ali", 18, "iran"

	// printf will format out string with the args we pass to it
	//* %v => general | %s => string | %T => var type | %d => integer | %f => float | %q => escaped string | %c => rune
	//* \n creates a new line in console
	fmt.Printf("hello %d years old %s from %s \n", age, name, address)

	// print var in console
	//* to do it we should import fmt(formatter) package.
	fmt.Println(age)

	//	declaring a constance => Must be uppercase
	//* constance value must be defined at compile time, not in run time.
	//* const PI = 3.14 => true
	//* const A = x * 2 => err || it will define in run time.
	const PI = 3.1415926
	fmt.Println(PI)

	// default value of int datatypes is 0 and for strings is empty string "", for the others is nil(like null in js)
	var city string
	fmt.Println(city) // => ""

	// we can change int datatypes to each other
	var balance = 50
	var discountPercent float32 = 70.9
	var newDiscountPercent = int(discountPercent) // => 70, because it is not float anymore
	// var finalPrice int = balance * (100 - discountPercent) => this will get err, because datatype is int but result is float
	var finalPrice = int(float32(balance) * (float32(100) - discountPercent))
	fmt.Println(finalPrice, newDiscountPercent)

	// use func
	message, number := sayHello("ali", "iran", age)
	println(message, number)

	makeSlice()
}

// declaring functions
// + if func name starts with an uppercase letter, we can use it in other package. it is like we export it
// func sayHello(name string, address string, age int) string { return "" }
// or
func sayHello(name, address string, age int) (string, int) { // multi return
	// conditions
	if name == "john" && age == 20 {
		return "you are not allowed", 0
	} else if age < 18 || name == "mohammad" {
		return "you are not allowed", 0
	} else {
		fmt.Println("ha ha")
	}
	// += | *= | /= | -= | %=

	// switch case
	var test string = "test"
	switch test {
	case "john":
		fmt.Println("ha ha")
		fallthrough // this will run the next case without checking the case condition
	case "mohammad":
		fmt.Println("ha ha")
	default:
		fmt.Println("ha ha")
	}

	// we can declare var in condition scope like
	if a := 0; a > 1 {
		fmt.Println("ha ha")
	}

	// Loops
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	message := fmt.Sprintf("hello %d years old %s from %s", age, name, address)

	return message, 1
}

// Struct
// + the first letter of struct name must be uppercase
type Person struct {
	name string
	age  int
}

func makeStruct(name string, age int) (Person, Person) {

	var person1 Person
	person1.name = name
	person1.age = age

	var person2 = Person{name, age}

	return person1, person2
}

// Arrays
/*
	var array_name = [length]datatype{values} // here length is defined
	or
	var array_name = [...]datatype{values} // here length is inferred
	or
	array_name := [length]datatype{values} // here length is defined
	or
	array_name := [...]datatype{values} // here length is inferred
*/
func makeArray(name string, age int) [2]Person {
	person1, person2 := makeStruct(name, age)

	var arr1 = [2]Person{person1, person2}

	arr2 := [...]Person{person1, person2}

	fmt.Println(arr2)

	return arr1
}

// Slices
/*
	slices are a view of arrays, which are not fixed size.
*/
func makeSlice() []int {
	var arr = [4]int{} // => default values are 0

	// declaring a slice would be like
	var slice1 = []int{1, 2, 3, 4, 5}
	// or
	slice2 := []int{1, 2, 3, 4, 5}

	// we can slice the array!
	/*
		[start:end] => will include the start index of array until the end index (end will not be included)
		[start:] => will include the start index of array to the last index
		[:end] => will include the first index of array until the end index (end will not be included)
		[:] => will include whole array
	*/
	slice3 := arr[0:2]

	fmt.Println(arr, slice1, slice2, slice3)

	// 2D Slices
	slice2D := [][]string{
		{"Hello", "World"},
		{"Hola", "Mundo"},
		{"Bonjour", "Le", "Monde"},
	}

	// slice length
	var slice2Length int = len(slice2)
	fmt.Println(slice2Length)

	// range and for each loop
	// for index, range { code }
	// range keyword goes for the all the string and arrays indexes
	for i := 0; i < len(slice2D); i++ {
		fmt.Printf("the %dth item \n", i+1)

		for _, v := range slice2D[i] {
			for _, w := range v {
				fmt.Printf("        %q\n", w)
			}
		}
	}

	return slice3
}
