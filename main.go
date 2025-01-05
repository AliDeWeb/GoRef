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

package main

func main() {}
