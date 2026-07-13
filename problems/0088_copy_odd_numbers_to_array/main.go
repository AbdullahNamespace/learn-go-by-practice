// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 88: copy_odd_numbers_to_array.go
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
// Write a program that fills a source array with random numbers between 1 and
// 100, then scans it and copies only the odd numbers into a new destination
// slice (which grows dynamically as odd numbers are found). Finally, print
// both the source array and the resulting array of odd numbers.
//
// العربية:
// اكتب برنامجًا يملأ مصفوفة مصدر بأرقام عشوائية بين 1 و 100، ثم يفحصها وينسخ
// فقط الأرقام الفردية إلى شريحة وجهة جديدة (تنمو ديناميكيًا كلما وُجد رقم فردي).
// أخيرًا، اطبع المصفوفة المصدرية والمصفوفة الناتجة التي تحتوي على الأرقام الفردية.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Array 1 = [10 7 4 11 6 13]
// Output: Array 2 Odd numbers: 7 11 13
//
// Example 2:
// Input:  Array 1 = [2 4 6 8]
// Output: Array 2 Odd numbers: (empty)
//   Why:    None of the source elements are odd
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Array elements are positive integers (1 to 100) | عناصر المصفوفة أعداد صحيحة موجبة
// • A number is odd if number % 2 != 0 | العدد فردي إذا كان باقي قسمته على 2 لا يساوي صفر
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func fillArrayWithRandomNumbers(lengthOfArray int) []int
// func randomNumber(min, max int) int
// func isOdd(number int) bool
// func copyOddNumbers(arrSource []int) []int
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

func fillArrayWithRandomNumbers(lengthOfArray int) []int {
	arr := make([]int, lengthOfArray)

	for key := range lengthOfArray {
		arr[key] = randomNumber(1, 100)
	}

	return arr
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func randomNumber(min, max int) int {
	return rand.IntN(max-min+1) + min
}

func isOdd(number int) bool {
	return number%2 != 0
}

func copyOddNumbers(arrSource []int) []int {
	var arr []int

	for _, value := range arrSource {
		if isOdd(value) {
			arr = append(arr, value)
		}
	}

	return arr
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printArray(arr []int) {
	fmt.Printf("Array element : %v\n", arr)
}

// ======================
//         MAIN
// ======================

func main() {

	lengthOfArray := readPositiveNumber("Please enter number of elements")

	arr := fillArrayWithRandomNumbers(lengthOfArray)

	fmt.Println("All array element")
	printArray(arr)

	oddArr := copyOddNumbers(arr)

	fmt.Println("Odd element")
	printArray(oddArr)
}
