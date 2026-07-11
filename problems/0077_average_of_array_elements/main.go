// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 77: average_of_array_elements.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Arrays & Math | المصفوفات والرياضيات
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that fills an array with random numbers between 1 and 100,
// prints the array elements, and then calculates and prints the average of all
// elements as a floating-point number.
//
// العربية:
// اكتب برنامجًا يملأ مصفوفة بأرقام عشوائية بين 1 و 100، يطبع عناصر المصفوفة،
// ثم يحسب ويطبع متوسط جميع العناصر كعدد عشري (Floating-point).
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:
//   Number of elements: 3
// Output:
//   Array Elements: 10 20 30
//   Average of all number is : 20
//
// Example 2:
// Input:
//   Number of elements: 2
// Output:
//   Array Elements: 15 25
//   Average of all number is : 20
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Array elements are positive integers (1 to 100). | عناصر المصفوفة أعداد صحيحة موجبة (من 1 إلى 100).
// • The result should be returned as a float to preserve decimal precision. | يجب إرجاع النتيجة كعدد عشري للحفاظ على الدقة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func randomNumber(from, to int) int
// func fillArrayWithRandomNumbers(arr []int)
// func sumArray(arr []int) int
// func arrayAverage(arr []int) float64
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

func sumArray(arr []int) int {
	sum := 0

	for _, value := range arr {
		sum += value
	}

	return sum
}

func arrayAverage(arr []int) float64 {
	return float64(sumArray(arr)) / float64(len(arr))
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printArray(arr []int) {
	fmt.Printf("Array Elements: %v \n", arr)
}

func printSumOfArray(arr []int) {
	fmt.Printf("Sum of all number is : %d\n", sumArray(arr))
}

func printAverageOfArray(arr []int) {
	fmt.Printf("Average of all number is : %.2f\n", arrayAverage(arr))
}

// ======================
//         MAIN
// ======================

func main() {
	lengthOfArray := readPositiveNumber("Please enter numbers of elements")

	arr := make([]int, lengthOfArray)

	fillArrayWithRandomNumbers(arr)

	printArray(arr)
	printSumOfArray(arr)
	printAverageOfArray(arr)
}
