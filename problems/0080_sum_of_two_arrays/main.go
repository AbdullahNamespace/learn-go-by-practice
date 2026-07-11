// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 80: sum_of_two_arrays.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Arrays | المصفوفات
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that reads a positive number representing the array length, then
// fills two arrays with random numbers between 1 and 100. The program calculates
// the element-wise sum of the two arrays and stores the result in a third array.
// Finally, it prints all three arrays.
//
// العربية:
// اكتب برنامجًا يقرأ عددًا موجبًا يمثل طول المصفوفة، ثم يملأ مصفوفتين بأرقام
// عشوائية بين 1 و 100. يحسب البرنامج المجموع عنصر بعنصر (Element-wise sum)
// للمصفوفتين ويخزن النتيجة في مصفوفة ثالثة. أخيرًا، يطبع المصفوفات الثلاث.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:
//   How many elements? 3
// Output:
//   Array 1 elements:
//   10 20 30
//   Array 2 elements:
//   5 15 25
//   Sum of array1 and array2 elements:
//   15 35 55
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • All three arrays must have the same length (> 0). | جميع المصفوفات الثلاث يجب أن يكون لها نفس الطول (> 0).
// • Random numbers are in the range [1, 100]. | الأرقام العشوائية تكون في المدى [1، 100].
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPositiveNumber(prompt string) int
// func randomNumber(from, to int) int
// func fillArrayWithRandomNumbers(arr []int)
// func sumOf2Arrays(arr1 []int, arr2 []int) []int
// func printArray(arr []int)
//
// ────────────────────────────────────────────────────────────────────────────
package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

// ======================
//   UTILITY
// ======================

func printError(prompt string) {
	fmt.Printf("X Error : %s\n", prompt)
}

// ======================
//     INPUT FUNCTIONS
// ======================

func readString(prompt string) string {
	fmt.Printf("%s : ", prompt)

	var input string

	fmt.Scan(&input)

	return input
}

func readNumber(prompt string) int {
	for {
		input := readString(prompt)

		inputNumber, err := strconv.Atoi(input)

		if err != nil {
			printError("Invalid input please enter a valid number!")
			continue
		}

		return inputNumber
	}
}

func readPositiveNumber(prompt string) int {
	for {
		input := readNumber(prompt)

		if input <= 0 {
			printError("Invalid input please enter a positive number!")
			continue
		}

		return input
	}
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func randomNumber(from, to int) int {
	return rand.IntN(to-from+1) + from
}

func fillArrayWithRandomNumbers(arr []int) {
	for key := range arr {
		arr[key] = randomNumber(1, 100)
	}
}

func sumOf2Arrays(arr1 []int, arr2 []int) []int {
	arr := make([]int, len(arr1))

	for key := range arr1 {
		arr[key] = arr1[key] + arr2[key]
	}

	return arr
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printArray(arr []int) {
	fmt.Printf("Array Elements is %v\n", arr)
}

// ======================
//         MAIN
// ======================

func main() {
	lengthOfArray := readPositiveNumber("Please enter numbers of elements")

	arr1 := make([]int, lengthOfArray)
	arr2 := make([]int, lengthOfArray)

	fillArrayWithRandomNumbers(arr1)
	fillArrayWithRandomNumbers(arr2)

	fmt.Println("Array 1")
	printArray(arr1)

	fmt.Println("Array 2")
	printArray(arr2)

	sumArr := sumOf2Arrays(arr1, arr2)

	fmt.Println("Sum Of 2 Array")
	printArray(sumArr)

}
