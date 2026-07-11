// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 79: copy_only_prime_numbers.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Arrays & Math | المصفوفات والرياضيات
// **Difficulty:** Medium | متوسط
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that fills an array with random numbers, then copies only the
// prime numbers from the source array into a destination array. Finally, it prints
// the source array and the destination array containing only the prime numbers.
//
// العربية:
// اكتب برنامجًا يملأ مصفوفة بأرقام عشوائية، ثم ينسخ فقط الأعداد الأولية (Prime)
// من المصفوفة المصدرية إلى مصفوفة الوجهة. أخيرًا، يطبع المصفوفة المصدرية
// والمصفوفة الوجهة التي تحتوي فقط على الأعداد الأولية.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:
//   Number of elements: 6
// Output:
//   Array 1 elements:
//   10 7 4 11 6 13
//   Prime Numbers in Array2:
//   7 11 13
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Array elements are positive integers (1 to 100). | عناصر المصفوفة أعداد صحيحة موجبة (من 1 إلى 100).
// • A prime number is a number greater than 1 that has no divisors other than 1 and itself. | العدد الأولي هو عدد أكبر من 1 لا يقبل القسمة إلا على 1 وعلى نفسه.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func checkPrime(number int) bool
// func randomNumber(from, to int) int
// func fillArrayWithRandomNumbers(arr []int)
// func copyPrimeNumbers(arr []int) []int
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
			printError("Invalid input, please enter a valid number!")
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

func checkPrime(number int) bool {
	if number <= 1 {
		return false
	}
	if number == 2 {
		return true
	}
	if number%2 == 0 {
		return false
	}
	for i := 3; i <= number; i += 2 {
		if number%i == 0 {
			return false
		}
	}
	// for i := 3; i*i <= number; i += 2 {
	// 	if number%i == 0 {
	// 		return false
	// 	}
	// }
	return true
}

func randomNumber(from, to int) int {
	return rand.IntN(to-from+1) + from
}

func fillArrayWithRandomNumbers(arr []int) {
	for key := range arr {
		arr[key] = randomNumber(1, 100)
	}
}

func copyPrimeNumbers(arr []int) []int {
	var result []int

	for _, value := range arr {
		if checkPrime(value) {
			result = append(result, value)
		}
	}
	return result
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printArray(arr []int) {
	fmt.Printf("Array Elements : %v\n", arr)
}

// ======================
//         MAIN
// ======================

func main() {
	lengthOfArray := readPositiveNumber("Please enter number of elements")

	arr := make([]int, lengthOfArray)

	fillArrayWithRandomNumbers(arr)
	printArray(arr)

	arrCopy := copyPrimeNumbers(arr)

	fmt.Println("Prime Numbers in Array2")
	printArray(arrCopy)
}
