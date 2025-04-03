package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello World!")
	// variables and how to declare/use them.
	var intNum int = 12
	fmt.Println(intNum)

	var floatNum float64 = 123.345598759
	fmt.Println(floatNum)

	// operations with numbers
	var result float64 = floatNum + float64(intNum)
	fmt.Print(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myText string = "Hi"
	var myLongText string = `This is a 
	multiline text lol`

	// string methods
	fmt.Println(myText)
	fmt.Println(myLongText)
	fmt.Println(utf8.RuneCountInString(myText))

	var myRune rune = 'a'
	fmt.Println(myRune)

	myNewText := "hi"
	fmt.Println(myNewText)

	const myConst string = "this is a const"
	printConst(myConst)

	var resultSum = sumOfTwoNumbers(4, 8)
	fmt.Println(resultSum)

	var originalValue, resultMult int = multiplyTwo(3, 4)
	fmt.Printf("original value is %v and the result is %v", originalValue, resultMult)

	// division of integers with conditionals.
	var resultDiv, remainder, err = intDivision(3, 4)

	switch {
	case err != nil:
		fmt.Print(err.Error())
	case remainder == 0:
		fmt.Printf("the result of the division is %v", resultDiv)
	default:
		fmt.Printf("The result of the division is %v with remainder %v", resultDiv, remainder)
	}

	/* Arrays in Go */
	var arrayofInts [3]int32
	arrayofInts[1] = 4
	fmt.Println(arrayofInts[1])
	//fmt.Println(&arrayofInts[1])

	/* Slices in Go */
	var intSlice []int32 = []int32{1, 5, 4}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 9)
	fmt.Println(intSlice)
	fmt.Printf("The lenght of the slice is %v and the capacity is %v", len(intSlice), cap(intSlice))

	/* maps in Go */
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var execMap = map[string]uint8{"James": 15, "Rick": 59}
	fmt.Println(execMap["James"])

	var age, ok = execMap["James"]
	//delete(execMap, "James")
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid name")
	}

	// iteration
	for name := range execMap {
		fmt.Printf("Name: %v", name)
	}
}

// -- functions in Go
func printConst(printValue string) {
	fmt.Println(printValue)
}

func sumOfTwoNumbers(numberOne int, numberTwo int) int {
	var result int = numberOne + numberTwo
	return result
}

// function with two return values
func multiplyTwo(originalVal int, timesMult int) (int, int) {
	var result int = originalVal * timesMult
	return originalVal, result
}

// function for division of integers.
func intDivision(numerator int, denominator int) (int, int, error) {
	var err error

	if denominator == 0 {
		err = errors.New("cannot divide by 0")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
