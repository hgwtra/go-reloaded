// It is recommended to have test files for unit testing.

package main

import (
	"fmt"
	"io/ioutil"
	"os" //builtin package to open files
	"strconv"
	"strings"
)

func main() {
	input := os.Args[1]
	output := os.Args[2]

	//read the input file
	inputByte, err := ioutil.ReadFile(input)

	//check for err
	if err != nil {
		fmt.Println(err) //print error
		os.Exit(1)       //exit(1) means there was some err and that's why the program exited
	}

	//turn input bytes into string
	newstr := string(inputByte)
	//fmt.Println(newstr)

	//turn string into an arr
	arr := strings.Split(newstr, " ")

	//fixing cap up low hex bin
	for i, element := range arr {
		// cap
		if strings.Contains(arr[i], "(cap") {
			if strings.Contains(arr[i], "(cap)") {
				arr[i-1] = strings.Title(arr[i-1])
			} else {
				capArr := strings.Split(arr[i+1], "")
				// fmt.Println(capArr[0])
				num := sliceAtoi(capArr[0])
				for x := num; x > 0; x-- {
					arr[i-x] = strings.Title(arr[i-x])
				}
			}

			//remove cap and number
			if arr[i] == "(cap)" {
				arr[i] = ""
			} else if strings.Contains(arr[i], "(cap,") {
				arr[i] = ""
				arr[i+1] = arr[i+1][len("2)"):]
			} else {
				arr[i] = arr[i][len("(cap)"):]
			}
		}

		// //seperate . from i+1
		// if strings.Contains(arr[i], ",") {
		// 	splitword := strings.Split(arr[i], "")
		// 	splitword = splitword[1:]
		// 	// fmt.Println(splitword)
		// 	str := strings.Join(splitword, "")
		// 	// fmt.Println(str)

		// 	punc := puncTuations(arr[i])
		// 	// fmt.Println("this is punc", punc)

		// 	arr[i-2] = arr[i-2] + punc
		// 	arr[i] = ""
		// 	arr = append(arr, str)

		// }

		//up
		if strings.Contains(arr[i], "(up") {
			if strings.Contains(arr[i], "(up)") {
				arr[i-1] = strings.ToUpper(arr[i-1])
			} else {
				capArr := strings.Split(arr[i+1], "")
				num := sliceAtoi(capArr[0])
				for x := num; x > 0; x-- {
					arr[i-x] = strings.ToUpper(arr[i-x])
				}
			}

			//removeup
			if arr[i] == "(up)" {
				arr[i] = ""
			} else if strings.Contains(arr[i], "(up,") {
				arr[i] = ""
				arr[i+1] = arr[i+1][len("2)"):]
			} else {
				arr[i] = arr[i][len("(up)"):]
			}
		}

		//low
		if strings.Contains(arr[i], "(low") {
			if strings.Contains(arr[i], "(low)") {
				arr[i-1] = strings.ToLower(arr[i-1])
			} else {
				capArr := strings.Split(arr[i+1], "")
				num := sliceAtoi(capArr[0])
				for x := num; x > 0; x-- {
					arr[i-x] = strings.ToLower(arr[i-x])
				}
			}

			if arr[i] == "(low)" {
				arr[i] = ""
			} else if strings.Contains(arr[i], "(low,") {
				arr[i] = ""
				arr[i+1] = arr[i+1][len("2)"):]
			} else {
				arr[i] = arr[i][len("(low"):]
			}
		}

		//hex
		if strings.Contains(arr[i], "(hex)") {
			outputhex, err := strconv.ParseInt(arr[i-1], 16, 64) //to integer
			if err != nil {
				fmt.Println(err)
				return
			}
			arr[i-1] = strconv.Itoa(int(outputhex)) //integer to string

			if arr[i] == "(hex)" {
				arr[i] = ""
			} else {
				arr[i] = arr[i][len("(hex)"):]
			}
		}

		//bin
		if strings.Contains(arr[i], "(bin)") {
			outputbin, err := strconv.ParseInt(arr[i-1], 2, 64) //to integer
			if err != nil {
				fmt.Println(err)
				return
			}
			arr[i-1] = strconv.Itoa(int(outputbin))

			if arr[i] == "(bin)" {
				arr[i] = ""
			} else {
				arr[i] = arr[i][len("(bin)"):]
			}
		}

		if element == "A" && firstLetter(arr[i+1]) {
			arr[i] = "An"
			//fmt.Println(arr)
		}

		if element == "a" && firstLetter(arr[i+1]) {
			arr[i] = "an"
			//fmt.Println(arr)
		}

		if strings.Contains(arr[i], ",") {
			if element == "," {
				arr[i] = ""
				arr[i-1] = arr[i-1] + ","
			} else {
				arr[i-1] = arr[i-1] + ","
				arr[i] = arr[i][len("1"):]
			}
		}
	}

	//punctuations: . , ! ? : ; -

	//Except groups of punctuation

	//' awesome ' => 'awesome'

	str := strings.Join(arr, " ")                 //turns array into string
	res := strings.Join(strings.Fields(str), " ") //remove duplicate spaces

	// fmt.Println(str)
	fmt.Println(res)

	final := []byte(res) //string to byte

	//write to output file
	err = ioutil.WriteFile(output, final, 0666) //bytesRead, 0644) //what's this?

	if err != nil {
		fmt.Println(err)
		os.Exit(1) //exit(1) means there was some err and that's why the program exited
	}
}

func firstLetter(s string) bool {
	r := []rune(s)
	if r[0] == 'u' || r[0] == 'U' || r[0] == 'e' || r[0] == 'E' || r[0] == 'o' || r[0] == 'O' || r[0] == 'a' || r[0] == 'A' || r[0] == 'i' || r[0] == 'I' || r[0] == 'h' || r[0] == 'H' {
		return true
	}
	return false
}

func puncTuations(s string) string {
	r := []rune(s)
	return string(r[0])
}

func sliceAtoi(s string) int {
	intvar, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return intvar
}
