package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

func main() {
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 35}
	fmt.Println(myEngine.mpg, myEngine.gallons)

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

	// iteration/loops
	for name, age := range execMap {
		fmt.Printf("Name: %v, Age: %v", name, age)
	}

	for i, v := range arrayofInts {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := range 10 {
		fmt.Println(i)
	}

	//strings, runes and Bytes
	var myString = []rune("résumé")
	var indexed = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("\n The lenght of 'myString' is %v", len(myString))

	var strSlice = []string{"h", "e", "l", "l", "o"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)
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
